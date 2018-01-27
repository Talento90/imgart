package mongo

import mgo "gopkg.in/mgo.v2"

// Session mongo client
type Session struct {
	*mgo.Session
	Database string
}

// NewSession creates a new mongo session
func NewSession(c Configuration) (*Session, error) {
	session, err := mgo.Dial(c.MongoURL)

	if err != nil {
		return nil, err
	}

	return &Session{Session: session, Database: c.Database}, nil
}

// Check mongo health
func (c *Session) Check() error {
	return c.Ping()
}
