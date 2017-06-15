package mongo

import "gopkg.in/mgo.v2/bson"
import "github.com/mikamikuh/goannotator/model"

func (s MongoSession) GetAnnotations(uri string) ([]model.Annotation, error) {
	results := []model.Annotation{}
	err := s.collection.Find(bson.M{"uri": uri}).All(&results)

	return results, err
}
