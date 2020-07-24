package test

import (
	"fmt"
	"github.com/zngue/go_tool/src/michttp"
	"testing"
)

func TestHttp(t *testing.T)  {

	c:=michttp.RequesAll(michttp.HttpRequestMap{
		"syName":michttp.MicHttpRequest{
			ServiceId:"sy:composition",
			EndPoint: "/es/artList",
			Method: "GET",
		},
		"userName":michttp.MicHttpRequest{
			ServiceId: "sy:go:integral_mall",
			EndPoint: "/api/shop/getList",
			Method: "GET",
		},
	})
	fmt.Println(c)
	fmt.Println("123132123")

}
