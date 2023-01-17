package test

import (
	"fmt"
	"testing"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-04 16:39:10
* @content: service
 */

func TestXX(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			if i==4{
				close(ch)
				close(ch)
			}
			ch<-i
		}

	}()

	for data := range ch {
		fmt.Println(data)
	}
}
