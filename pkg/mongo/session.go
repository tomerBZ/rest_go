package mongo

import (
	"github.com/globalsign/mgo"
	"github.com/tomerBZ/web/pkg/logs"
	"os"
	"time"
)

type Session struct {
	session *mgo.Session
}

func NewSession() (*Session, error) {
	maxWait := time.Duration(5 * time.Second)
	mongoHost := os.Getenv("MONGO_HOST")
	mongoPort := os.Getenv("MONGO_PORT")
	logs.Info.Println("connecting to mongo at:", mongoHost+":"+mongoPort)
	session, sessionErr := mgo.DialWithTimeout(mongoHost+":"+mongoPort, maxWait)
	if sessionErr != nil {
		logs.Error.Printf("Unable to connect to local mongo instance! at %s:%s", mongoHost, mongoPort)
		return nil, sessionErr
	}
	return &Session{session}, sessionErr
}

func (s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func (s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
