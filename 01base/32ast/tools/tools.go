package tools

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/ast"
	"go/format"
	"go/token"
	"log"
	"os"
	"strings"
)

func OutStdout(f *ast.File) {
	if err := format.Node(os.Stdout, token.NewFileSet(), f); err != nil {
		log.Fatalln("Error:", err)
	}
}
func OutFile(f *ast.File, fileName,outFile string) {
	CreateDir("generate")

	fileName = strings.TrimSuffix(fileName[strings.LastIndex(fileName, "/")+1:],".go")
	if outFile==""{
		outFile = "./generate/"+fileName+"_generate.go"
	}
	open, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	if err := format.Node(open, token.NewFileSet(), f); err != nil {
		log.Fatalln("Error:", err)
	}


}
func OutFile1(f *dst.File, fileName,outFile string) {
	CreateDir("generate")

	fileName = strings.TrimSuffix(fileName[strings.LastIndex(fileName, "/")+1:],".go")
	if outFile==""{
		outFile = "./generate/"+fileName+"_generate.go"
	}
	open, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	err = decorator.Fprint(open, f)
	if err!=nil{
		panic(err)
	}
}

// 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}
// 创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}