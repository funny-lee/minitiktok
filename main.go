// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"minitiktok/biz/dal"
)

func main() {
	Init()
	h := server.Default()

	register(h)
	h.Spin()
}

func Init() {
	dal.InitDB()
}
