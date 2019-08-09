package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 1. 获取浏览器请求中的vid-id
	vid := ps.ByName("vid-id")
	// 2. 根据vid-id在服务器工程的videos文件夹下进行文件查找
	videoPath := VIDEO_DIR + vid

	video, err := os.Open(videoPath)
	// 2.1 未找到, 提示用户此视频不存在
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal error")
		return
	}
	// 2.2 找到, 写出给浏览器, 进行播放
	// 告知响应的数据类型
	w.Header().Set("Content-Type", "video/mp4")
	// 播放视频
	http.ServeContent(w, r, "", time.Now(), video)
}

func testPageHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 接收到请求后, 需要让其跳转到指定的界面
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)

}
func uploadHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 1. 限制浏览器发送视频的大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	// 2. 需要将上传的视频文件进行解析, 并且写入到video文件夹中
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	// 2.1 读取文件
	// 2.2 读取成功后,需要将文件存储到磁盘中
	data, err := ioutil.ReadAll(file)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	fileName := ps.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fileName, data, 0644)
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// 3. 告知用户上传数据成功或失败
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Upload Successfully")

}
