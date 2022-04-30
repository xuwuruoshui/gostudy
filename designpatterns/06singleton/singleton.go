package singleton

import (
	"math/rand"
	"sync"
)

// 单例模式

type Instance struct {
	id int
}

var instance *Instance

var once sync.Once

func GetInstance() *Instance{
	
	once.Do(func() {
		instance = &Instance{id:rand.Int()}
	})
	return instance
}