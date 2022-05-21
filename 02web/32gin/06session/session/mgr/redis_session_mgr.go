package mgr

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
	session2 "gostudy/32gin/06session/session/category"
	"gostudy/32gin/06session/session/inf"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	addr string
	passwd string
	pool *redis.Pool
	rwlock sync.RWMutex
	sessionMap map[string]inf.Session
}

func NewRedisSessionMgr() inf.SessionMgr {
	return &RedisSessionMgr{
		sessionMap: make(map[string]inf.Session,32),
	}
}


func (m *RedisSessionMgr)Init(addr string,options ...string) (err error){
	if len(options)>0{
		m.passwd = options[0]
	}
	// 创建连接池
	m.pool = createPool(addr,m.passwd)
	m.addr = addr
	return
}

func createPool(addr,password string)*redis.Pool{
	return &redis.Pool{
		MaxIdle: 64,
		MaxActive: 1000,
		IdleTimeout: 240*time.Second,
		Dial: func() (redis.Conn, error) {
			conn,err := redis.Dial("tcp",addr)
			if err!=nil{
				return nil,err
			}
			if _,err := conn.Do("AUTH",password);err!=nil{
				conn.Close()
				return nil,err
			}
			return conn,err
		},
		// 连接测试,生产环境注释掉
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err := c.Do("PING")
			return err
		},
	}
}

func (m *RedisSessionMgr)CreateSession()(session inf.Session,err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()

	id:= uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建session
	session = session2.NewRedisSession(sessionId,m.pool)
	m.sessionMap[sessionId] = session
	return
}

func (m *RedisSessionMgr)Get(sessionId string)(session inf.Session,err error){
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()

	session,ok := m.sessionMap[sessionId]
	if !ok{
		err = errors.New("session not exist")
		return
	}
	return
}