package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	log "github.com/shenjing023/llog"
)

const (
	XUEQIU_URL = "https://xueqiu.com"
)

var (
	defaultOptions = options{
		timeout:  5,
		maxRetry: 3,
		maxConns: 10,
	}
	defaultXueqiu = NewXueqiu()
)

type Xueqiu struct {
	c       *http.Client
	cookies []*http.Cookie
	m       sync.RWMutex
	opt     *options
}

type options struct {
	timeout  int
	maxRetry int
	maxConns int
}

type StockData struct {
	Symbol string `json:"code"`
	Name   string `json:"query"`
	State  int    `json:"state"`
	Type   int    `json:"type"`
}

type CommentData struct {
	ErrorResp `json:",omitempty"`
	List      []struct {
		description string `json:"-"`
		Id          int    `json:"id"`
		CreatedAt   int64  `json:"created_at"`
		ViewCount   int    `json:"view_count"`
	} `json:"list,omitempty"`
}

type ErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SearchResp struct {
	ErrorResp
	Data []StockData `json:"data,omitempty"`
}

// Option set option function
type Option func(*options)

func NewXueqiu(opts ...Option) *Xueqiu {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	xq := &Xueqiu{
		opt: &o,
		c: &http.Client{
			Timeout: time.Duration(o.timeout) * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: o.maxConns,
				MaxConnsPerHost:     o.maxConns,
			},
		},
	}
	go func() {
		for {
			xq.updateCookie()
			time.Sleep(time.Hour)
		}
	}()
	return xq
}

func (x *Xueqiu) updateCookie() {
	res, err := http.Get(XUEQIU_URL)
	if err != nil {
		log.Infof("update cookie failed: %v", err)
		return
	}
	x.m.Lock()
	defer x.m.Unlock()
	x.cookies = res.Cookies()
}

func WithTimeout(timeout int) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

func WithMaxRetry(maxRetry int) Option {
	return func(o *options) {
		o.maxRetry = maxRetry
	}
}

func WithMaxConns(maxConns int) Option {
	return func(o *options) {
		o.maxConns = maxConns
	}
}

func (x *Xueqiu) request(url, method string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Host", "xueqiu.com")
	req.Header.Add("Referer", "https://xueqiu.com/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	x.m.RLock()
	for _, cookie := range x.cookies {
		req.AddCookie(cookie)
	}
	x.m.RUnlock()
	return x.c.Do(req)
}

func (x *Xueqiu) SearchStock(symbolorname string) ([]StockData, error) {
	url := fmt.Sprintf("https://xueqiu.com/query/v1/suggest_stock.json?q=%s&count=5", url.QueryEscape(symbolorname))
	for i := 0; i < x.opt.maxRetry; i++ {
		res, err := x.request(url, http.MethodGet, nil)
		if err != nil {
			log.Infof("url[%s] request failed: %v", url, err)
			time.Sleep(time.Second)
			continue
		}
		defer res.Body.Close()
		t, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var searchResp SearchResp
		err = json.Unmarshal(t, &searchResp)
		if err != nil {
			return nil, fmt.Errorf("json unmarshal failed: %v", err)
		}
		if searchResp.Code != 200 && i != x.opt.maxRetry-1 {
			log.Infof("url[%s] request failed msg[%s] and retry: %d", url, searchResp.Message, i+1)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		return searchResp.Data, nil
	}
	return nil, errors.New("request failed")
}

func SearchStock(symbolorname string) ([]StockData, error) {
	return defaultXueqiu.SearchStock(symbolorname)
}

func GetXqCommentData(lastPullTime int64, symbol string) (map[string]int, error) {
	return defaultXueqiu.GetCommentData(lastPullTime, symbol)
}

// 拉取股票讨论数据
func (x *Xueqiu) GetCommentData(lastPullTime int64, symbol string) (map[string]int, error) {
	result := make(map[string]int)
	for page := 1; page <= 50; page++ {
		url := "https://xueqiu.com/query/v1/symbol/search/status.json?count=20&comment=0&symbol=%s&hl=0&source=all&sort=time&page=%d&q=&type=11"
		url = fmt.Sprintf(url, symbol, page)
		data, err := x.getComment(url)
		if err != nil {
			return nil, err
		}
		flag := x.parseComment(data, lastPullTime, result)
		if flag {
			break
		}
		time.Sleep(time.Second)
	}
	return result, nil
}

func (x *Xueqiu) getComment(url string) (*CommentData, error) {
	var data CommentData
	for i := 0; i < x.opt.maxRetry; i++ {
		res, err := x.request(url, http.MethodGet, nil)
		if err != nil {
			log.Infof("url[%s] request failed: %v", url, err)
			time.Sleep(time.Second)
			continue
		}
		defer res.Body.Close()
		if res.StatusCode == http.StatusOK {
			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				return &data, fmt.Errorf("json unmarshal failed: %v", err)
			}
		}

		if data.Code != 0 && i != x.opt.maxRetry-1 {
			log.Infof("url[%s] request failed msg[%s] and retry: %d", url, data.Message, i+1)
			time.Sleep(time.Second * 1)
			continue
		}
		break
	}
	return &data, nil
}

func (x *Xueqiu) parseComment(data *CommentData, lastPullTime int64, result map[string]int) (flag bool) {
	for _, v := range data.List {
		t := v.CreatedAt / 1000
		if t >= lastPullTime {
			day := time.Unix(t, 0).Local().Format("2006-01-02")
			if _, ok := result[day]; !ok {
				result[day] = 1
			} else {
				result[day]++
			}
		} else {
			flag = true
			break
		}
	}
	return
}
