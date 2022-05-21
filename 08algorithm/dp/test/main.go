package main

import "fmt"

/**
* @creator: xuwuruoshui
* @date: 2021-10-06 22:32:27
* @content: 最长回文字符串
 */

func main() {
	fmt.Println(longestPalindrome("ccc"))
}

func longestPalindrome(s string) string {

	max := 0
	str := ""


	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			str1 := s[i : j+1]
			if max > (len(str1)) {
				break
			}
			ok := isPalindrome(str1)
			if ok && len(str1) > max {
				str = str1
				max = len(str1)
				break
			}
		}
	}
	return str
}

// 判断是否为回文数
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
