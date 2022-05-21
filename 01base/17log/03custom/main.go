package main

import (
	"log"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 16:38:29
* @content: 客制化日志
 */

func main() {
	// 标准输出，前缀，配置
	logger := log.New(os.Stdout, "<Golang>", log.Lshortfile|log.Lshortfile|log.Ltime)
	logger.Println("客制化日志")
}
