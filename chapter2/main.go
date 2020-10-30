package main

import (
	_ "chapter2/matchers"
	"chapter2/search"
	"log"
	"os"
)

func init() {
	//将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	//使用特定的项做搜索
	search.Run("president")
}