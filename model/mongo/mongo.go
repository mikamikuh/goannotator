package mongo

import "gopkg.in/mgo.v2"
import "github.com/mikamikuh/goannotator/model"

type Session interface {
	Close()
	GetAnnotations(uri string) ([]model.Annotation, error)
}

type MongoSession struct {
	session    *mgo.Session
	collection *mgo.Collection
}

func NewSession(url, database, collection string) (Session, error) {
	session, err := mgo.Dial(url)

	if err != nil {
		return nil, err
	}

	c := session.DB(database).C(collection)

	return MongoSession{
		session:    session,
		collection: c,
	}, nil
}

func (s MongoSession) Close() {
	s.session.Close()
}
