package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/defs"
	"video_server/api/session"
	"video_server/api/user"
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
	if err := user.AddUser(ubody.UserName, ubody.Pwd); err != nil {
		// 2.1 告知用户操作数据库的时候出现错误
		sendErrorResponse(resp, defs.ErrorDBError)
		return
	}
	// 3. 针对当前的登录用户生成session
	sessionId := session.GenerateNewSessionId(ubody.UserName)
	// 3.1 记录此sessionId指向的用户, 处于登录状态
	su := &defs.SignedUp{true, sessionId}

	if res, err := json.Marshal(su); err != nil {
		sendErrorResponse(resp, defs.ErrorInternalFaults)
		return
	} else {
		// 4. 完整的完成登录流程, 告知用户注册成功

	}

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
