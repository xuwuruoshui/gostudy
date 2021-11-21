package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

/**
* @creator: xuwuruoshui
* @date: 2021-11-21 16:58:14
* @content:
 */

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int `ini:"database"`
}

type Config struct {
	*MysqlConfig `ini:"mysql"`
	*RedisConfig `ini:"redis"`
}

func loadIni(conf string, data interface{}) (err error) {

	// 1.判断是否为指针类型，是否为结构体类型
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() != reflect.Ptr {
		err = errors.New("不是指针类型")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("不是结构体类型")
		return
	}

	// 2.读文件(按开头读判断)
	file, err := os.Open(conf)
	if err != nil {
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	// 3.按换行拆分，遍历
	lineContent := strings.Split(string(content), "\r\n")
	// 遍历到当前字段
	var currFiled string = ""
	for _, line := range lineContent {

		// 3.1、去空格
		line = strings.TrimSpace(line)

		
		// 3.2、注释
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") || line == "" {
			continue
		} else if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// 3.3、头
			currFiled = string(line[1 : len(line)-1])
			continue
		} else {
			// 3.4、具体配置
			if currFiled != "" {
				//3.4.1 判断配置是否符合书写规范
				filedContent := strings.Split(line, "=")
				if strings.HasPrefix(line, "=") || len(filedContent)!=2 {
					err = errors.New("ini文件内容不正确")
					return
				}
				
				// redis或者mysql的struct
				var childStuct reflect.StructField
				var childStuctValue reflect.Value
				for i := 0; i < t.Elem().NumField(); i++ {
					if t.Elem().Field(i).Tag.Get("ini") == currFiled{
						childStuct = t.Elem().Field(i)
						childStuctValue = v.Elem().Field(i)
					}
				}

				//遍历mysql或者redis的属性，判断是否与ini中配置相同
				for i := 0; i < childStuct.Type.Elem().NumField(); i++ {
					if childStuct.Type.Elem().Field(i).Tag.Get("ini")==filedContent[0]{
						if childStuct.Type.Elem().Field(i).Type.Kind()==reflect.String{
							childStuctValue.Elem().Field(i).SetString(filedContent[1])
						}
						if childStuct.Type.Elem().Field(i).Type.Kind()==reflect.Int{
							value,_ := strconv.ParseInt(filedContent[1],10,64)
							childStuctValue.Elem().Field(i).SetInt(value)
						}
					}
				}

				
			}
		}
	}
	return
}

func main() {
	var mysql = MysqlConfig{}
	var redis = RedisConfig{}
	var cfg Config = Config{&mysql, &redis}
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg.MysqlConfig)
	fmt.Println(cfg.RedisConfig)
}
