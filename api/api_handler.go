package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {

}

func HandlerOriginalTest(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 1. 获取请求完整连接, 获取查询参数
	query := req.URL.Query()
	// 通过关键字key获取对应的value
	name := query.Get("user_name")
	pwd := query.Get("pwd")
	fmt.Printf("name = %s, pwd = %s\n", name, pwd)
}

func HandlerRestFullTest(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	uname := ps.ByName("user_name")
	fmt.Printf("userName = %s\n", uname)
	io.WriteString(resp, uname)
}
