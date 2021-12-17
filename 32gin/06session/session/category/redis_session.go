package category

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId string

	//TODO
	pool *redis.Pool

	sessionMap map[string]interface{}

	rwlock sync.RWMutex

	// 记录内存中map是否变化
	flag int
}

const(
	// 内存数据无变化
	SessionFlagNoe = iota

	// 内存数据有变化
	SessionFlagModify
)

func NewRedisSession(id string,pool *redis.Pool)*RedisSession {
	return &RedisSession{
		sessionId:  id,
		sessionMap: make(map[string]interface{},16),
		pool:       pool,
		flag:       SessionFlagNoe,
	}
}

func (r *RedisSession) Set(key string,value interface{}) (err error){
	r.rwlock.Lock()
	defer r.rwlock.Lock()
	r.sessionMap[key] = value
	r.flag = SessionFlagModify
	return
}

func (r *RedisSession) Get(key string)(value interface{},err error){
	r.rwlock.RLock()
	defer r.rwlock.RUnlock()

	value,ok := r.sessionMap[key]
	if !ok{
		err = errors.New("key not exist")
	}

	return
}

func (r *RedisSession)loadFromRedis()(err error){
	conn := r.pool.Get()
	reply,err := conn.Do("GET",r.sessionId)
	if err!=nil{
		return
	}

	// 转字符串
	data,err := redis.String(reply,err)
	if err!=nil{
		return
	}

	// 取到的东西，反序列化到内存的map
	err = json.Unmarshal([]byte(data),&r.sessionMap)
	if err!=nil{
		return
	}
	return
}

func (r *RedisSession) Del(key string) (err error){
	r.rwlock.Lock()
	defer r.rwlock.Lock()
	delete(r.sessionMap,key)

	return
}

func (r *RedisSession) Save() (err error){
	r.rwlock.Lock()
	defer r.rwlock.Lock()

	// 判断redis中是否有变化
	if r.flag== SessionFlagNoe {
		return
	}

	// 内存中的sessionMap进行序列化
	data,err := json.Marshal(r.sessionMap)
	if err!=nil{
		return
	}
	// 获取redis连接
	conn := r.pool.Get()

	// 保存kv
	_,err = conn.Do("SET",r.sessionId,string(data))
	if err!=nil{
		return
	}
	r.flag = SessionFlagNoe
	return
}