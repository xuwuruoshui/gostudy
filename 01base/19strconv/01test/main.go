package main

import (
	"fmt"
	"strconv"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-02 15:31:51
* @content: strconv
 */

func main() {
	// 1. string 2 int
	a := "123"
	b, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", b, b)

	// 2. int 2 string
	c := 123
	d := strconv.Itoa(c)
	fmt.Printf("%T, %v\n", d, d)

	// 3. string 2 bool
	e := "true"
	f, err := strconv.ParseBool(e)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", f, f)

	// 4. string 2 int
	g := "-12"
	h, err := strconv.ParseInt(g, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", h, h)

	// 5. string 2 uint
	i := "33"
	j, err := strconv.ParseUint(i, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", j, j)

	// 6. string 2 float
	k := "3.14"
	l, err := strconv.ParseFloat(k, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T, %v\n", l, l)

	// 7. bool 2 string
	m := true
	n := strconv.FormatBool(m)
	fmt.Printf("%T, %v\n", n, n)

	// 8. int 2 string
	o := 123
	p := strconv.FormatInt(int64(o), 10)
	fmt.Printf("%T, %v\n", p, p)

	// 9. uint 2 string
	q := uint(33)
	r := strconv.FormatUint(uint64(q), 10)
	fmt.Printf("%T, %v\n", r, r)

	// 10. float 2 string
	s := 3.14
	// 传入值, 进制, 精度,位数
	t := strconv.FormatFloat(s, 'f', 3, 64)
	fmt.Printf("%T, %v\n", t, t)

	// 11. isPrint
	fmt.Println(strconv.IsPrint(32))	
	
	// 12. 是否为单行的、没有空格和tab之外控制字符的反引号字符串
	fmt.Println(strconv.CanBackquote("hahah fdsfaf\t\n"))
}
