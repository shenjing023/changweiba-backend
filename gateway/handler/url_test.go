package handler

import (
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	target := "svc-account:///account"
	u, err := url.Parse(target)
	if err != nil {
		panic(err)
	}
	t.Log(u.Scheme)

}
