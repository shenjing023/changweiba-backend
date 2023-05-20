package handler

import (
	"fmt"
	"testing"
	"time"
)

func TestSearchStock(t *testing.T) {
	xq := NewXueqiu()
	time.Sleep(time.Second * 30)
	data, err := xq.SearchStock("SH601020")
	if err != nil {
		fmt.Println(err)
	}
	t.Log(data)
}

func TestParseComment(t *testing.T) {
	xq := NewXueqiu()
	data := CommentData{
		List: []struct {
			description string `json:"-"`
			Id          int    `json:"id"`
			CreatedAt   int64  `json:"created_at"`
			ViewCount   int    `json:"view_count"`
		}{
			{CreatedAt: 1638276800000, Id: 4, ViewCount: 0},
			{CreatedAt: 1628276800000, Id: 3, ViewCount: 0},
			{CreatedAt: 1538276800000, Id: 2, ViewCount: 0},
			{CreatedAt: 1528276800000, Id: 1, ViewCount: 0},
		},
	}

	result := make(map[string]int)
	flag := xq.parseComment(&data, 1628276800, result)
	fmt.Println(flag)
	t.Log(result)
}

func TestGetStockCommentData(t *testing.T) {
	xq := NewXueqiu()
	time.Sleep(time.Second * 3)
	data, err := xq.GetCommentData(1640074878, "SH601020")
	if err != nil {
		fmt.Println(err)
	}
	t.Log(data)
}
