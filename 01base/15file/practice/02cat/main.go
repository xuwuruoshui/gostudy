package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-29 14:50:22
* @content: 实现linux下的cat命令
 */

func main() {

	// 1.解析命令行参数
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("no param")
		//cat(bufio.NewReader(os.Stdin))
	}

	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil { 
			fmt.Fprintf(os.Stdout, "reading from %s faild,err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err != nil {
			// 
			fmt.Fprintf(os.Stdout, "%s", buf)
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
