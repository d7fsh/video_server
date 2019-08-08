package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/session"
	"video_server/api/user"
)

func CreateUser(resp http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	//1.从req中获取浏览器发送过来的请求内容(账号,密码(加密))
	ubody := &defs.UserCredential{}
	res, _ := ioutil.ReadAll(req.Body)
	fmt.Println("ctx.Request.body = ", string(res))

	if err := json.Unmarshal(res, ubody); err != nil {
		//如果解析并封装到结构体的过程中,出现问题,说明浏览器发送的字段,没有按照服务器要求传递
		sendErrorResponse(resp, defs.ErrorRequestBodyParseFailed)
		return
	}
	//2.将获取数据插入数据库
	if err := user.AddUser(ubody.UserName, ubody.Pwd); err != nil {
		//2.1 告知用户操作数据的时候出现错误
		sendErrorResponse(resp, defs.ErrorDBError)
		return
	}
	//3.针对当前的登录用户生成session
	sessionId := session.GenerateNewSessionId(ubody.UserName)
	//3.1记录此sessionId指向的用户,处于登录状态
	su := &defs.SignedUp{true, sessionId}

	//3.2 结构体转json
	if suJson, err := json.Marshal(su); err != nil {
		sendErrorResponse(resp, defs.ErrorInternalFaults)
		return
	} else {
		//4,完整的完成注册流程,告知用户注册成功
		sendNormalResponse(resp, string(suJson), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 1. 获取请求体中传递过过来的json
	res, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", res)
	var ubody = &defs.UserCredential{}

	// 2. 解析json, 通过json内容, 查询数据库中是否存在此用户
	if err := json.Unmarshal(res, ubody); err != nil {
		//如果解析并封装到结构体的过程中,出现问题,说明浏览器发送的字段,没有按照服务器要求传递
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}
	pwd, err := dbops.GetUserCredential(ubody.UserName)
	if err != nil || len(pwd) == 0 || pwd != ubody.Pwd {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
	}
	// 3. 给用户生成sessionId, 记录此用户登录标识
	sessionId := session.GenerateNewSessionId(ubody.UserName)
	//3.1记录此sessionId指向的用户,处于登录状态
	su := &defs.SignedUp{true, sessionId}
	//3.2 结构体转json
	if suJson, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	} else {
		//4,完整的完成注册流程,告知用户注册成功
		sendNormalResponse(w, string(suJson), 200)
	}

}

func HandlerOriginalTest(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
	//获取请求完整链接,获取查询参数
	query := req.URL.Query()
	//通过关键字key获取相应value
	name := query.Get("user_name")
	pwd := query.Get("pwd")

	fmt.Println("name = ", name)
	fmt.Println("pwd = ", pwd)
}

//func HandlerTest(resp http.ResponseWriter, req *http.Request, params httprouter.Params) {
//	uname := params.ByName("user_name")
//	fmt.Println("uname = ", uname)
//	io.WriteString(resp, uname)
//}
