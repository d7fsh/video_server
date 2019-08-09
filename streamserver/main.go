package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := RegisterHandler()
	// 设置一个中间件
	// 限制播放视频的最大连接数
	mh := NewMiddleWareHandler(router, 2)
	http.ListenAndServe(":8081", mh)
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	// 播放视频请求
	router.GET("/videos/:vid-id", streamHandler)
	// 访问上传视频页面
	router.GET("/testpage", testPageHandler)
	// 上传视频的处理
	router.POST("/upload/:vid-id", uploadHandler)
	return router
}

func NewMiddleWareHandler(r *httprouter.Router, maxSize int) http.Handler {
	m := middleWareHandler{}
	m.router = r
	m.l = NewConnLimiter(maxSize)
	return m
}

type middleWareHandler struct {
	router *httprouter.Router
	l      *ConnLimiter
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 此处需要对同时访问服务器, 播放视频的连接数进行限制数量
	// 1. 如果获取连接失败(返回失败原因)
	if !m.l.GetConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	// 2. 如果获取连接成功, 正常访问服务器
	m.router.ServeHTTP(w, r)

	// 3. 在访问完了服务器后, 释放连接
	defer m.l.DropConn()
}
