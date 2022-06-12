package main

import (
	"fmt"
)

var (
	greenBg      = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	whiteBg      = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellowBg     = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	redBg        = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blueBg       = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magentaBg    = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyanBg       = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	green        = string([]byte{27, 91, 51, 50, 109})
	white        = string([]byte{27, 91, 51, 55, 109})
	yellow       = string([]byte{27, 91, 51, 51, 109})
	red          = string([]byte{27, 91, 51, 49, 109})
	blue         = string([]byte{27, 91, 51, 52, 109})
	magenta      = string([]byte{27, 91, 51, 53, 109})
	cyan         = string([]byte{27, 91, 51, 54, 109})
	reset        = string([]byte{27, 91, 48, 109})
	disableColor = false
)

func main() {
	str := "hello world"
	fmt.Println(greenBg, str, reset)
	fmt.Println(whiteBg, str, reset)
	fmt.Println(yellowBg, str, reset)
	fmt.Println(redBg, str, reset)
	fmt.Println(blueBg, str, reset)
	fmt.Println(magentaBg, str, reset)
	fmt.Println(cyanBg, str, reset)
	word := "I love you"
	fmt.Println(green, word, reset)
	fmt.Println(white, word, reset)
	fmt.Println(yellow, word, reset)
	fmt.Println(red, word, reset)
	fmt.Println(blue, word, reset)
	fmt.Println(magenta, word, reset)
	fmt.Println(cyan, word, reset)

	fmt.Print(0x1B, '[', ';', 'm', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', "\n")
	fmt.Printf("%#X\t%c\t%c\t%c\t", 27, 91, 59, 109)
	fmt.Printf("%c\t%c\t%c\t%c\t%c\t%c\t%c\t%c\t%c\t%c\t", 48, 49, 50, 51, 52, 53, 54, 55, 56, 57)
	fmt.Println()
	// 27代表0x1B

	// 91代表[

	// 59代表;

	// 109代表m

	// 57代表9，表示设置字体颜色

	// 52代表4，表示设置背景色

	// 51代表3，表示设置前景色，也就是文字的颜色

	// 90到97与30到37的效果一样，一个是设置字体颜色，一个是设置前景色，所以57和51可以互换，效果完全一样，

	// reset表示0x1B[0m，表示清除颜色设置。

	fmt.Printf("\x1b[%dmhello world 30: 黑 \x1b[0m\n", 30)
	fmt.Printf("\x1b[%dmhello world 31: 红 \x1b[0m\n", 31)
	fmt.Printf("\x1b[%dmhello world 32: 绿 \x1b[0m\n", 32)
	fmt.Printf("\x1b[%dmhello world 33: 黄 \x1b[0m\n", 33)
	fmt.Printf("\x1b[%dmhello world 34: 蓝 \x1b[0m\n", 34)
	fmt.Printf("\x1b[%dmhello world 35: 紫 \x1b[0m\n", 35)
	fmt.Printf("\x1b[%dmhello world 36: 深绿 \x1b[0m\n", 36)
	fmt.Printf("\x1b[%dmhello world 37: 白色 \x1b[0m\n", 37)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 47: 白色 30: 黑 \n", 47, 30)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 46: 深绿 31: 红 \n", 46, 31)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 45: 紫   32: 绿 \n", 45, 32)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 44: 蓝   33: 黄 \n", 44, 33)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 43: 黄   34: 蓝 \n", 43, 34)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 42: 绿   35: 紫 \n", 42, 35)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 41: 红   36: 深绿 \n", 41, 36)
	fmt.Printf("\x1b[%d;%dmhello world \x1b[0m 40: 黑   37: 白色 \n", 40, 37)

}
