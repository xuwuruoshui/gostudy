package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

/**
* @creator: xuwuruoshui
* @date: 2022-02-15 18:38:01
* @content: viper使用
 */

type Config struct {
	Port    int    `mapstructure:"port"`
	Address string `mapstructure:"address"`
	Version string `mapstructure:"version"`
}

func main() {

	viper.SetConfigFile("./03viper/config/conf.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic("Read file error:" + err.Error())
	}
	var conf = &Config{}
	if err := viper.Unmarshal(conf); err != nil {
		panic("File exchange error:" + err.Error())
	}
	fmt.Println(conf.Address, conf.Port)
	// 监控配置文件的变化
	viper.WatchConfig()
	// viper文件发送变化事件
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(conf); err != nil {
			panic("File exchange error:" + err.Error())
		}
	})

	r := gin.Default()
	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, conf.Version)
	})
	//err := r.Run(fmt.Sprintf("%s:%d", viper.GetString("address"), viper.GetInt("port")))
	fmt.Println(conf.Address, conf.Port)
	err := r.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	if err != nil {
		panic(err)
	}
}
