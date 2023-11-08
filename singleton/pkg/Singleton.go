package pkg

import "log"

type Singleton struct {
	Type string
}

func (s Singleton) Print() {
	log.Printf("Type %s\n", s.Type)
}

func NewSingleton(item *Singleton, typeName string) *Singleton {
	if item == nil {
		return &Singleton{
			Type: typeName,
		}
	}

	log.Printf("Type %s - is created\n", item.Type)

	return item
}
