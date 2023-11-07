package pkg

import "log"

type Subscriber struct {
	Name string
}

func (s *Subscriber) Update(pubName string) {
	log.Printf("Sending to subscribe %s from publisher %s\n", s.GetName(), pubName)
}

func (s *Subscriber) GetName() string {
	return s.Name
}
