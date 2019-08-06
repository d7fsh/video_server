package demo

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// 匹配请求体中的多个字段
type VideoInfo struct {
	Id       string `json:"id"`
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

func main() {
	// 1. 封装对路由的监听
	r := RegisterHandler()
	// 2. 指定监听的端口, 响应请求路由
	log.Fatal(http.ListenAndServe(":8080", r))
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user/:user_name", CreateUser)
	// 让工程中的竞态文件能够被访问到
	router.ServeFiles("/static/*filepath", http.Dir("./templates"))
	return router
}

func CreateUser(resp http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	// 匹配上了请求链接后, 会触发的代码逻辑
	userName := ps.ByName("user_name")
	fmt.Println("username = ", userName)
	// 获取请求的方式
	fmt.Println("req.Method = ", req.Method)
	// 获取请求完整连接地址
	fmt.Println("req.URL = ", req.URL)

	// 获取ajax请求过程中, 带过来的json数据
	// 1. 从Request中获取请求体
	reqBody, _ := ioutil.ReadAll(req.Body)

	// 2. 将json数据匹配到结构体中的每一个字段中
	info := &VideoInfo{}
	if err := json.Unmarshal(reqBody, info); err != nil {
		log.Println("err = ", err)
		return
	}

	// 3. 将结构体中包含的信息打印一下
	fmt.Println("info = ", info)
	// 通过response写出响应吗
	resp.WriteHeader(200)
	// 通过resp写出给浏览器的内容
	io.WriteString(resp, "success")

}
