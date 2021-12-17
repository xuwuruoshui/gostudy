package category

import (
	"errors"
	"gostudy/32gin/06session/session/inf"
	"sync"
)

type MemorySession struct {

	sessionId string

	// 存kv
	data map[string]interface{}
	rwlock sync.RWMutex
}


func NewMemorySession(id string) inf.Session {
	return &MemorySession{
		sessionId: id,
		data: make(map[string]interface{},16),
	}
}

func (m *MemorySession)Set(key string,value interface{})(err error){
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}

func (m *MemorySession)Get(key string)(value interface{},err error){
	m.rwlock.RLock()
	defer m.rwlock.RUnlock()
	value,ok := m.data[key]
	if !ok{
		err = errors.New("key no exist in session")
	}
	return
}

func (m *MemorySession)Del(key string)(err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data,key)
	return
}

func (m *MemorySession)Save()(err error) {
	return
}