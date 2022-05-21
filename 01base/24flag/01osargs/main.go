package main

import (
	"fmt"
	"os"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-08 17:32:12
* @content:
 */

func main(){
	// os.Args 是一个 []string
	if len(os.Args)>0{
		for index,arg := range os.Args {
			fmt.Printf("args[%d]=%v\n",index,arg)
		}
	}
}