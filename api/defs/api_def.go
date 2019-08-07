package defs

type SimpleSession struct {
	Username string // session指向的用户名
	TTL      int64  // session有效时间
}

type UserCredential struct {
	UserName string `json:"user_name"` // 存储用户名称字段
	Pwd      string `json:"pwd"`       // 存储用户密码字段
}

// 存储某用户登录与否的结构体
type SignedUp struct {
	Succss    bool   `json:"succss"`     // 登录 true, 非登录 flase
	SessionId string `json:"session_id"` // sessionId字段
}
