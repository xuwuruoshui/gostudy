package main

import (
	"bufio"
	"net"
	"os"
	"strings"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-25 18:20:36
* @content: 客户端
 */

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(conn)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		
		content, err := reader.ReadString('\n')

		if err != nil {
			panic(err)
		}
		if strings.HasSuffix(content, "Q") {
			conn.Write([]byte(content))
			break
		}
		_, err = conn.Write([]byte(content))
		if err != nil {
			panic(err)
		}
	}

}
