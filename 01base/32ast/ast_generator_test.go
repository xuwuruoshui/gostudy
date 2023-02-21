package test

import (
	"32ast/tools"
	"testing"
)

func TestGenerateByCode(t *testing.T) {
	type args struct {
		fileName string
		funcName string
		outFile  string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "写代码生成", args: args{fileName: "./example/code.go", funcName: "Haha", outFile: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tools.TimeCalculateByCode(tt.args.fileName, tt.args.funcName, tt.args.outFile)
		})
	}
}

func TestGenerateByTemplate(t *testing.T) {
	type args struct {
		fileName     string
		templateName string
		funcName     string
		outFile      string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "写代码生成", args: args{fileName: "./example/code.go", templateName: "./example/test.txt", funcName: "Haha", outFile: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tools.TimeCalculateTemplate(tt.args.fileName, tt.args.templateName, tt.args.funcName, tt.args.outFile)
		})
	}
}

func TestConstGenerateMap(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "生成map", args: args{fileName: "./example/enum.go"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tools.ConstGenerateMap(tt.args.fileName)
		})
	}
}

func TestConstGenerateMap2(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test", args: args{fileName: "./example/enum.go"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tools.ConstGenerateMap2(tt.args.fileName)
		})
	}
}
