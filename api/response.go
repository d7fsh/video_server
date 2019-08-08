package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(resp http.ResponseWriter, errResp defs.ErrResponse) {
	//给浏览器一个响应码
	resp.WriteHeader(errResp.HttpSC)
	//将errResp中Error字段的结果,转换成字符串,写出给浏览器
	errorByte, _ := json.Marshal(&errResp.Error)
	io.WriteString(resp, string(errorByte))
}

//参数一:用于给浏览器写出数据
//参数二:响应给浏览器的字符串
//参数三:响应码
func sendNormalResponse(resp http.ResponseWriter, responseStr string, sc int) {
	resp.WriteHeader(sc)
	io.WriteString(resp, responseStr)
}
