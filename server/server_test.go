package main

import (
	"changweiba-backend/common"
	"changweiba-backend/conf"
	"changweiba-backend/graphql"
	"changweiba-backend/graphql/dataloader"
	"changweiba-backend/pkg/middleware"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestServer(t *testing.T){
	pwd, _ := os.Getwd()
	execDir := flag.String("d", pwd, "execute directory")
	flag.Parse()
	fmt.Println("Current execute directory is:", *execDir)
	conf.InitConfig(*execDir)
	graphql.InitRPCConnection()
	registerSignalHandler()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Setting up Gin
	r := gin.Default()
	r.Use(common.GinContextToContextMiddleware())
	r.Use(middleware.JWTMiddleware(conf.Cfg.SignKey,conf.Cfg.QueryDeep))
	r.Use(dataloader.LoaderMiddleware())

	r.POST("/graphql", graphqlHandler())

	//构建返回值
	w := httptest.NewRecorder()
	//构建请求
	a:=`{"query":"query a{\n  posts(page:0,pageSize:1){\n    nodes{\n      id,\n      reply_num,\n      \n    },\n    total_count\n  }\n}\n"}`
	req:= httptest.NewRequest("POST", "/graphql",strings.NewReader(a) )
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRhY2htZW50IjoxMDAwMywiZXhwIjoxNTcxNjU1MTE3LCJpYXQiOjE1NzE2NTE1MTcsIm5iZiI6MTU3MTY1MTUxN30.B7Rf88WlHkw6yfEqicrE9-iNsvuezlwbxmz1aIwE2wg")
	//调用请求接口
	r.ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}

func init(){
	
}

func BenchmarkServer(t *testing.B){
	//conf.InitConfig(`C:\GoProjects\src\changweiba-backend\server`)
	//graphql.InitRPCConnection()
	//registerSignalHandler()
	//port := os.Getenv("PORT")
	//if port == "" {
	//	port = defaultPort
	//}
	//
	//// Setting up Gin
	//r := gin.Default()
	//r.Use(common.GinContextToContextMiddleware())
	//r.Use(middleware.JWTMiddleware(conf.Cfg.SignKey,conf.Cfg.QueryDeep))
	//r.Use(dataloader.LoaderMiddleware())
	//
	//r.POST("/graphql", graphqlHandler())
	//
	////构建返回值
	//w := httptest.NewRecorder()
	//构建请求
	t.StopTimer()
	a:=`{"query":"query a{\n  posts(page:0,pageSize:1){\n    nodes{\n      id,\n      reply_num,\n      \n    },\n    total_count\n  }\n}\n"}`
	req,_:= http.NewRequest("POST", "http://127.0.0.1:8088/graphql",strings.NewReader(a) )
	req.Header.Set("Content-Type","application/json")
	req.Header.Set("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRhY2htZW50IjoxMDAwMywiZXhwIjoxNTcxNjYwNjE0LCJpYXQiOjE1NzE2NTcwMTQsIm5iZiI6MTU3MTY1NzAxNH0.PaN6lvqeG3l76-83sjpOXAAwsOl7kNkbPqNxROqIKjI")
	//调用请求接口
	//r.ServeHTTP(w, req)

	c:= http.DefaultClient
	
	t.StartTimer()
	for i := 0; i < t.N; i++ {
		resp,_:=c.Do(req)
		_, _ = ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
	}
	
	//t.RunParallel(func(pb *testing.PB) {
	//	for pb.Next(){
	//		req,err:= http.NewRequest("POST", "http://127.0.0.1:8088/graphql",strings.NewReader(a) )
	//		if err!=nil{
	//			fmt.Println(err)
	//			t.Fail()
	//		}
	//		req.Header.Set("Content-Type","application/json")
	//		req.Header.Set("token","eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRhY2htZW50IjoxMDAwMywiZXhwIjoxNTcxNjYwNjE0LCJpYXQiOjE1NzE2NTcwMTQsIm5iZiI6MTU3MTY1NzAxNH0.PaN6lvqeG3l76-83sjpOXAAwsOl7kNkbPqNxROqIKjI")
	//		resp,err:=c.Do(req)
	//		if err!=nil{
	//			fmt.Println(err.Error())
	//			t.Fail()
	//		}
	//		_, err = ioutil.ReadAll(resp.Body)
	//		if err!=nil{
	//			fmt.Println(err)
	//			t.Fail()
	//		}
	//		defer resp.Body.Close()
	//	}
	//})
}
