package main

import "gostudy/15file/practice/03log/logger"

/**
* @creator: xuwuruoshui
* @date: 2021-11-13 11:56:15
* @content:
 */

func main(){
	log := logger.NewLogger("debug")

	log.Debug("Debug")
	log.Info("Info")
	log.Warning("Warning")
	log.Error("Error")
}