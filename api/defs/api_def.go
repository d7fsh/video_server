package defs

type SimpleSession struct {
	Username string // session指向的用户名
	TTL      int64 // session有效时间
}
