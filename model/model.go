package model

type Annotation struct {
	quote string
}

type Session interface {
	Close()
	GetAnnotations(uri string) ([]Annotation, error)
}
