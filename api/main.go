package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := RegisterHandler()
	// 使用中间件区分要登录的请求和无需登录的请求 ---------> 校验用的session是否有效
	// 装饰器模式, 就是原有功能的增强
	// 1. 原有的功能需要保留
	// 2. 原有的功能无法满足需求
	mh := NewMiddleWareHandler(router)
	log.Fatal(http.ListenAndServe(":8080", mh))
}

// 1. 定义一个方法, 用于返回http.Handler这个结构体实现类对象
func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.router = r
	return m
}

//2. 提供一个实现了http.Handler接口的结构体
type middleWareHandler struct {
	// 2.1 在中间件中保留原有的功能, 原有功能是存储在router
	router *httprouter.Router
}

// 3. 让middleWareHandler实现ServeHttp方法(重写ServerHTTP方法, 添加新功能)
func (m middleWareHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// 3.1 增强功能, 检测session是否存在
	validateUserSession(req)
	// 3.2 原有功能, 原有处理请求功能
	m.router.ServeHTTP(resp, req)

}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	//router.POST("/", CreateUser)
	router.GET("/test", HandlerOriginalTest)
	router.GET("/test/:user_name", HandlerRestFullTest)
	return router
}
