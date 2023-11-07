package pkg

import "log"

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Computer {
	return Notebook{
		Type:    NotebookType,
		Core:    6,
		Memory:  8,
		Monitor: true,
	}
}

func (pc Notebook) GetType() string {
	return pc.Type
}

func (pc Notebook) PrintDetails() {
	log.Printf("%s Core:[%d], Mem:[%d], Monitor: [%v]", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
