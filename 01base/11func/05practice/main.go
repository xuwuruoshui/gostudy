package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 分金币
 */

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println("每个人获取的金币：", distribution)
}

func dispatchCoin() int{
	for _, ws := range users {
		for _, w := range ws {
			switch {
			case w=='E' || w=='e':
				distribution[ws] ++
				coins--
			case w=='I' || w=='i':
				distribution[ws]+=2
				coins-=2
			case w=='O' || w=='o':
				distribution[ws]+=3
				coins-=3
			case w=='U' || w=='u':
				distribution[ws]+=4
				coins-=4
			}
		}
	}
	return coins
}