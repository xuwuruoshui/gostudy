package mgr

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	session2 "gostudy/32gin/06session/session/category"
	"gostudy/32gin/06session/session/inf"
	"sync"
)

type MemorySessionMgr struct {
	rwlock sync.RWMutex
	sessionMap map[string]inf.Session
}

func NewMemorySessionMgr() inf.SessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]inf.Session,1024),
	}
}

func (m *MemorySessionMgr)Init(addr string,option ...string) (err error){
	return
}

func (m *MemorySessionMgr) CreateSession()(session inf.Session,err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	id := uuid.NewV4()

	// 转string
	sessionId := id.String()
	// 创建session
	session = session2.NewMemorySession(sessionId)
	return
}

func (m *MemorySessionMgr) Get(sessionId string)(session inf.Session,err error){
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	session,ok := m.sessionMap[sessionId]
	if !ok{
		err = errors.New("session is not exist")
		return
	}
	return
}