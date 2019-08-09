package main

import (
	"io"
	"net/http"
)

// 给浏览器异常响应
func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string) {
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}

