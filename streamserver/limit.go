package main

import "fmt"

type ConnLimiter struct {
	maxConn int      // 记录最大连接数
	bucket  chan int //记录当前连接数(用管道容量进行控制),控制代码执行(逻辑)
}

// 结构体创建
func NewConnLimiter(maxSize int) *ConnLimiter {
	return &ConnLimiter{
		maxConn: maxSize,
		// 管道缓冲区大小, 和maxSize相同
		bucket: make(chan int, maxSize),
	}
}

// 获取链接
func (cl *ConnLimiter) GetConn() bool {
	// 1. 获取管道中, 已用的缓冲区大小, 和最大连接数进行比较
	if len(cl.bucket) >= cl.maxConn {
		fmt.Println("已达最大连接数")
		return false
	}
	// 1.1 已用缓冲区的大小 >= 最大连接数(不能再有更多的访问链接)
	// 1.2 已用缓冲区的大小 < 最大连接数(可以建立新的链接), 把信连接记录效率, 想管道中再写入一份数据
	cl.bucket <- 1
	return true
}

// 释放连接
func (cl *ConnLimiter) DropConn() {
	// 从bucket中读取数据, 每读取一份数据, 等同于给chan腾出一个缓冲区空间, 可以接入新的链接
	c := <-cl.bucket
	fmt.Println("读取bucket中的数据: ", c, "可接入新连接")
}
