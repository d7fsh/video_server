package main

import (
	"net/http"
	"video_server/api/session"
)

// 请求头(cookies和session对应关系)
// 提供一个常量用于记录服务器给客户端生成的sssionId
var HEADER_FILED_SESSION = "x-Session-Id"
var HEADER_FILED_USERNAME = "x-User-Name"

func validateUserSession(req *http.Request) bool {
	// 根据常量, 从客户端发送过来的请求中, 获取sessionId
	sid := req.Header.Get(HEADER_FILED_SESSION)
	// 校验过程, 判断sid是否存在
	if len(sid) == 0 {
		return false
	}
	// 根据sid获取session值
	username, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	// 将username存储在请求中, 让后续处理请求的方法使用
	req.Header.Add(HEADER_FILED_USERNAME, username)
	return true
}
