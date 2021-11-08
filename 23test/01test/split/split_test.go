package split

import (
	"reflect"
	// "os"
	"testing"
)

// // 1.普通单元测试(以Test开头)
// func TestSplit(t *testing.T) {
// 	got := Split("a:b:c", ":")
// 	want := []string{"a", "b", "c"}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("expect:%v,got:%v", want, got)
// 	}
// }

// func TestSplit1(t *testing.T) {
// 	got := Split("abcdef", "bc")
// 	want := []string{"a", "def"}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("expect:%v,got:%v", want, got)
// 	}
// }

// // 2.测试组
// func TestSplitGroup(t *testing.T) {
// 	// 1.定义测试类
// 	type test struct {
// 		input string
// 		sep   string
// 		want  []string
// 	}

// 	tests := []test{
// 		{input: "a:bc:d:e", sep: ":", want: []string{"a", "bc", "d", "e"}},
// 		{input: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
// 		{input: "呵哈呵哈哈呵哈", sep: "呵", want: []string{"哈", "哈哈", "哈"}},
// 	}

// 	for _, test := range tests {
// 		got := Split(test.input, test.sep)
// 		if !reflect.DeepEqual(got, test.want) {
// 			t.Errorf("expect:%#v,got:%#v", test.want, got)
// 		}
// 	}
// }

// 3.子测试
// // 单条测试命令:go test -v -run=./split
// func TestSplitChild(t *testing.T) {
// 	// 1.定义测试类
// 	type test struct {
// 		input string
// 		sep   string
// 		want  []string
// 	}

// 	tests := map[string]test{
// 		"easy":   {input: "a:bc:d:e", sep: ":", want: []string{"a", "bc", "d", "e"}},
// 		"normal": {input: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
// 		"hard":   {input: "呵哈呵哈哈呵哈", sep: "呵", want: []string{"", "哈", "哈哈", "哈"}},
// 	}

// 	for name, test := range tests {
// 		t.Run(name, func(t *testing.T) {
// 			got := Split(test.input, test.sep)
// 			if !reflect.DeepEqual(got, test.want) {
// 				t.Errorf("expect:%#v,got:%#v", test.want, got)
// 			}
// 		})

// 	}
// }

// // 4.覆盖率相关命令
// // go test -cover													测试覆盖率
// // go test -cover -coverprofile file		  覆盖率测试输出到一个文件
// // go tool cover -html=file								html中查看文件中覆盖率的情况

// // 5.基准测试(以Benchmark开头)
// func BenchmarkSplit(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Split("呵哈呵哈哈呵哈", "呵")
// 	}
// }

// 命令 go test -bench=Split -benchmem
// -benchmem可以看内存分配
// 基准测试结果
// Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkSplit$ gostudy\23test\01test\split

// goos: windows
// goarch: amd64
// pkg: gostudy/23test/01test/split
// cpu: Intel(R) Core(TM) i5-6300HQ CPU @ 2.30GHz
// BenchmarkSplit-4   	 3909456	       299.8 ns/op	     112 B/op	       3 allocs/op
// PASS
// ok  	gostudy/23test/01test/split	1.598s

// 其中BenchmarkSplit-4是指cup核心数即GOMAXPROCS
// 3909456次调用,平均结果为:每次调用耗时299.8 ns
// 每次操作分配了112B的内存,每次操作进行了3次内存分配
// 可以优化代码，比如提前分配足够的空间，每次操作就能避免扩容影响执行效率

// 6.并行测试
// func BenchmarkSplitParallel(b *testing.B) {
// 	b.RunParallel(func(p *testing.PB) {
// 		for p.Next() {
// 			Split("呵哈呵哈哈呵哈", "呵")
// 		}
// 	})
// }

// 7.测试前Setup和测试后TearDown
// func TestMain(m *testing.M) {
// 	fmt.Println("测试前相关操作")
// 	retCode := m.Run()
// 	fmt.Println("测试后相关操作")
// 	os.Exit(retCode)
// }

// 8.子测试，测试前Setup和测试后TearDown
// 所有测试之前
func setupTestCase(t *testing.T) func(t *testing.T){
	t.Log("所有测试之前")
	return func(t *testing.T){
		t.Log("所有测试之后")
	}
}
func setupSubTest(t *testing.T) func(t *testing.T){
	t.Log("单测试之前")
	return func(t *testing.T){
		t.Log("单条测试之后")
	}
}
func TestSplitChild(t *testing.T) {
	// 1.定义测试类
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := map[string]test{
		"easy":   {input: "a:bc:d:e", sep: ":", want: []string{"a", "bc", "d", "e"}},
		"normal": {input: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
		"hard":   {input: "呵哈呵哈哈呵哈", sep: "呵", want: []string{"", "哈", "哈哈", "哈"}},
	}
	tearDownTest := setupTestCase(t)
	defer tearDownTest(t)

	for name, test := range tests {	
		t.Run(name, func(t *testing.T) {
			tearDownTest := setupSubTest(t)
		defer tearDownTest(t)
			got := Split(test.input, test.sep)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("expect:%#v,got:%#v", test.want, got)
			}
		})

	}
}
