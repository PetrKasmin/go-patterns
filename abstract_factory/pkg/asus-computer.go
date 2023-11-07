package pkg

import "log"

type AsusComputer struct {
	Memory int
	Cpu    int
}

func (pc AsusComputer) PrintDetails() {
	log.Printf("[Asus] Pc Cpu: [%d], Memory: [%d]", pc.Cpu, pc.Memory)
}
