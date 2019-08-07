package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/defs"
)

func CreateUser(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 1. 从req中获取浏览器发送过来的请求内容(账号, 密码(加密))
	ubody := &defs.UserCredential{}
	res, _ := ioutil.ReadAll(req.Body)
	fmt.Printf("ctx.Request.body = %s", res)
	err := json.Unmarshal(res, ubody)
	if err != nil && ubody.UserName != "" && ubody.Pwd != "" {
		// 如果解析并封装到结构体的过程中出现问题, 说明浏览器发送的字段没有按照服务器的要求
		sendErrorResponse(resp, defs.ErrorRequestBodyParseFailed)
		return
	}
	// 2. 将获取数据插入数据库
	// 3. 针对当前的登录用户生成session
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
