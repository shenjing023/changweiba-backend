package handler

import (
	"fmt"
	"testing"
	"time"
)

func TestSearchStock(t *testing.T) {
	xq := NewXueqiu()
	time.Sleep(time.Second * 3)
	data, err := xq.SearchStock("sz123456")
	if err != nil {
		fmt.Println(err)
	}
	t.Log(data)
}
