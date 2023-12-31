package pkg

import "log"

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computer {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	log.Printf("%s Core:[%d], Mem:[%d], Monitor: [%v]", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
