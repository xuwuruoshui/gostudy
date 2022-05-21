package fibo

import (
	"testing"
	"time"
)

func benchmarkFibo(b *testing.B,n int) {
	for i := 0; i < b.N; i++ {
			Fibo(n)
	}
}

func BenchmarkFibo1(b *testing.B){
	benchmarkFibo(b,1)
}

func BenchmarkFibo2(b *testing.B){
	benchmarkFibo(b,2)
}

func BenchmarkFibo3(b *testing.B){
	benchmarkFibo(b,3)
}

func BenchmarkFibo4(b *testing.B){
	benchmarkFibo(b,10)
}

func BenchmarkFibo5(b *testing.B){
	benchmarkFibo(b,20)
}

// 应该多执行几次结果才有效
// go test -bench=Fibo6 -benchtime=20s
func BenchmarkFibo6(b *testing.B){
	benchmarkFibo(b,40)
}

// 存在耗时操作
func BenchmarkFibo7(b *testing.B){
	time.Sleep(time.Second*1)
	// 重置计时器
	b.ResetTimer()
	benchmarkFibo(b,5)
}

