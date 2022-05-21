package main

import (
	"context"
	"fmt"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-30 21:50:08
* @content: deadline
 */

func main() {

	d := time.Now().Add(5 * time.Second)
	ctx, cancle := context.WithDeadline(context.Background(), d)
	defer cancle()

	for{
		select {
		case <-time.After(time.Second):
			fmt.Println("haha")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}

}
