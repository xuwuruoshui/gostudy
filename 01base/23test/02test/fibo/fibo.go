package fibo

func Fibo(n int) int{
	if n<2{
		return 1
	}
	return Fibo(n-1)+Fibo(n-2)
}