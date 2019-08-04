package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
)

func main() {
	// 1. 封装对路由的监听
	r := RegisterHandler()
	// 2. 指定监听的端口, 响应请求路由
	log.Fatal(http.ListenAndServe(":8080", r))
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/:user_name", CreateUser)
	return router
}

func CreateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 匹配上了请求链接后, 会触发的代码逻辑
	userName := ps.ByName("user_name")
	fmt.Println("username = ", userName)
	// 获取请求的方式
	fmt.Println("r.Method = ", r.Method)
	// 获取请求完整连接地址
	fmt.Println("r.URL = ", r.URL)

	// 通过response写出响应吗
	w.WriteHeader(200)
	// 通过resp写出给浏览器的内容
	io.WriteString(w, "success")

}
