package pkg

import "log"

const (
	ServerType           = "server"
	PersonalComputerType = "computer"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string
	PrintDetails()
}

func New(typeName string) Computer {
	switch typeName {
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	default:
		log.Printf("%s Не существующий тип объекта", typeName)
		return nil
	}
}
