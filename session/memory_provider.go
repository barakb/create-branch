package session

import "fmt"

func init() {
	Register("memory", memoryProvider{make(map[string]memorySession)})
}

type memoryProvider struct {
	sessions map[string]memorySession
}

func (p memoryProvider) SessionInit(sid string) (Session, error) {
	fmt.Printf("SessionInit %s\n", sid)
	return p.getSessionCreated(sid)
}

func (p memoryProvider) SessionRead(sid string) (Session, error) {
	return p.getSessionCreated(sid)
}

func (p memoryProvider) getSessionCreated(sid string) (Session, error) {
	res, ok := p.sessions[sid]
	if !ok {
		res = memorySession{make(map[interface{}]interface{}), sid}
		p.sessions[sid] = res
	}
	return res, nil
}
func (p memoryProvider) SessionDestroy(sid string) error {
	delete(p.sessions, sid)
	return nil
}

func (p memoryProvider) SessionGC(maxLifeTime int64) {
	fmt.Println("Should implement memory provider gs !!!!")
	//todo
}

type memorySession struct {
	m   map[interface{}]interface{}
	sid string
}

func (s memorySession) Set(key, value interface{}) error {
	s.m[key] = value
	return nil
}

func (s memorySession) Get(key interface{}) interface{} {
	return s.m[key]
}

func (s memorySession) Delete(key interface{}) error {
	delete(s.m, key)
	return nil

}

func (s memorySession) SessionID() string {
	return s.sid
}
