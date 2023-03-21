package session

import (
	sessions "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// 全域變數
var Store = Session{}

type Session struct {
	requestSession sessions.Session
}

// Get 取得 session 值
func (s Session) Get(key string) interface{} {
	return s.requestSession.Get(key)
}

// Set 設置並儲存 session 值
func (s Session) Set(key string, value interface{}) {
	s.requestSession.Set(key, value)
}

// PreLoad 在 middleware 中將 session load 進去 Store 全域變數
func (s *Session) PreLoad(requestSession sessions.Session) {
	s.requestSession = requestSession
}

// 創建 session 引擎的方法, 可以依據需求改變其他引擎
func (s *Session) NewStore() sessions.Store {
	//基於cookie的 引擎， 参数是用於加密的 key
	store := cookie.NewStore([]byte("OWZHMGE3YZQTN2EXMS0ZN2JILTG4MDCTYJUYODM1ZMZJODM1"))
	return store
}
