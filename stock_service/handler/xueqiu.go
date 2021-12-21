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
	c        *http.Client
	cookies  []*http.Cookie
	m        sync.RWMutex
	maxRetry int
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

type SearchResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    []StockData `json:"data,omitempty"`
}

// Option set option function
type Option func(*options)

func NewXueqiu(opts ...Option) *Xueqiu {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	xq := &Xueqiu{
		c: &http.Client{
			Timeout: time.Duration(o.timeout) * time.Second,
			Transport: &http.Transport{
				MaxIdleConnsPerHost: o.maxConns,
				MaxConnsPerHost:     o.maxConns,
			},
		},
		maxRetry: o.maxRetry,
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
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.87 Safari/537.36")
	x.m.RLock()
	for _, cookie := range x.cookies {
		req.AddCookie(cookie)
	}
	x.m.RUnlock()
	return x.c.Do(req)
}

func (x *Xueqiu) SearchStock(symbolorname string) ([]StockData, error) {
	url := fmt.Sprintf("https://xueqiu.com/query/v1/suggest_stock.json?q=%s&count=5", url.QueryEscape(symbolorname))
	for i := 0; i < x.maxRetry; i++ {
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
		if searchResp.Code != 200 && i != x.maxRetry-1 {
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
