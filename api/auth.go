package main

import (
	"net/http"
	"video_server/api/session"
)

//请求头(cookies)和session对应关系
//提供一个常量用于记录服务器给客户端生成的sessionId
var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_USERNAME = "X-User-Name"

func validateUserSession(request *http.Request) bool {
	//根据常量,从客户端(浏览器)发送过来的请求中,获取sessionId
	sid := request.Header.Get(HEADER_FILED_SESSION)
	//校验过程,判断sid是否存在
	if len(sid) == 0 {
		return false
	}
	//根据sid获取session值
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	//将uname存储在请求头中,让后续处理请求的方法使用
	request.Header.Add(HEADER_FILED_USERNAME, uname)
	return true
}
