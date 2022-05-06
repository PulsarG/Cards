package moduls

type sessionData struct {
	Username string
}

type Session struct {
	data map[string]*sessionData
}

func NewSession() *Session {
	s := new(Session)

	s.data = make(map[string]*sessionData)

	return s
}

func (s *Session) Init(username string) string {
	sessionId := GenerateId()

	data := &sessionData{Username: username}

	s.data[sessionId] = data

	return sessionId
}
