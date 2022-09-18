package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
}

func GetEnv(s string)int{

	// windows $Env:ENV=1
	// linux export DEV=1
	viper.AutomaticEnv()
	return viper.GetInt(s)
}


func main()  {
	v := viper.New()

	env := GetEnv("ENV")
	v.SetConfigFile("./dev.yaml")
	if env!=0{
		v.SetConfigFile("./prod.yaml")
	}

	v.ReadInConfig()
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err!=nil{
		panic(err)
	}

	fmt.Println(cfg.Host)
	fmt.Println(cfg.Name)

}
