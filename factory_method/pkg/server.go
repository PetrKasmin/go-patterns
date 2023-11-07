package pkg

import "log"

type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Computer {
	return Server{
		Type:   ServerType,
		Core:   32,
		Memory: 256,
	}
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrintDetails() {
	log.Printf("%s Core:[%d], Mem:[%d]", pc.Type, pc.Core, pc.Memory)
}
