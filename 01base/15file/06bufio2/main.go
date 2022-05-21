package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
* @creator: xuwuruoshui
* @date: 2021-10-28 15:16:24
* @content: bufio
 */

func main() {
	
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0444))
	file.Chmod(0777)
	if err != nil {
		fmt.Println("open file failed,err", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("zhangsan" + strconv.Itoa(i) + "\n")
	}

	writer.Flush()

}
