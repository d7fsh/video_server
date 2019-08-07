package main

import (
	"encoding/json"
	"io"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(resp http.ResponseWriter, errResp defs.ErrResponse) {
	// 1. 给浏览器一个响应码
	resp.WriteHeader(errResp.HttpSC)
	// 2. 将errResp中Error字段的结果, 转换成字符串, 写出给浏览器
	errorBytes, _ := json.Marshal(&errResp.Error)
	io.WriteString(resp, string(errorBytes))
}
