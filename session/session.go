package session
import "crypto/rand"
import "fmt"
func GenerateId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
type sessionData struct {
	Username string
}
type Session struct {
	data map[string]*sessionData
}
func NewSession() *Session {
	s := new (Session)
	s.data = make(map[string]*sessionData)
	return s
}
func (s *Session) Init(username string) string {
	sessionId := GenerateId()
	data := &sessionData{Username: username}
	s.data[sessionId] = data
	return sessionId
}

func (s *Session) Get(sessionId string) string {
	data := s.data[sessionId]
	if data == nil {
		return ""
	}
	return data.Username

}
