package defs

// 1. 定义一个error的结构体, 在error结构体中包含具体错误原因, 错误代号(仅仅匹配错误原因)
// 比如: 代号001, 代表客户端给服务器发送的请求不正确
type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

// 2. 定义一个响应错误的结构体
// 响应码
// 响应内容: (包含Err的具体对象)
type ErrResponse struct {
	HttpSC int // 响应码, 符合http请求响应标准
	Error  Err // 具体错误(错误代号, 错误原因)
}

var (
	// 服务器解析用户数据失败
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "request body is not correct.", ErrorCode: "001"}}
	// 用户认证失败
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed", ErrorCode: "002"}}
	// 数据库访问失败
	ErrorDBError = ErrResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	// 服务器内部错误(session访问失败)
	ErrorInternalFaults = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal session error", ErrorCode: "004"}}
)
