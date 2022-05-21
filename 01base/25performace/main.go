package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-27 18:30:30
* @content: 性能分析
* 相关命令:go tool pprof cpu.pprof
 */

func testCode(){
	var c chan int
	for{
		select{
		case x:=<-c:
			fmt.Println(x)
		default:
		}
	}	
}

func main(){
	var isCPUProf bool
	var isMemProf bool

	flag.BoolVar(&isCPUProf,"cpu",false,"CPU性能测试")
	flag.BoolVar(&isMemProf,"mem",false,"内存性能测试")
	flag.Parse()

	// CPU性能测试
	if isCPUProf{
		file,err := os.Create("./cpu.pprof")
		if err!=nil{
			panic(err)
		}
		pprof.StartCPUProfile(file)
		//defer file.Close()
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 6; i++ {
		go testCode()
	}

	time.Sleep(time.Second*5)

	// 内存性能测试
	if isMemProf{
		file,err := os.Create("./mem.pprof")
		if err!=nil{
			panic(err)
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
	
}