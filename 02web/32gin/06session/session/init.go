package session

import (
	"errors"
	"gostudy/32gin/06session/session/inf"
	"gostudy/32gin/06session/session/mgr"
)

var(
	sessionMgr inf.SessionMgr
)

func Init(provider string,addr string,options ...string)(err error){
	switch provider {
	case "memory":
		sessionMgr = mgr.NewMemorySessionMgr()
	case "redis":
		sessionMgr = mgr.NewRedisSessionMgr()
	default:
		err = errors.New("not support method")
		return
	}

	err = sessionMgr.Init(addr,options...)
	return
}
