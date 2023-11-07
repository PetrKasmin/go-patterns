package pkg

import "log"

type HpComputer struct {
	Memory int
	Cpu    int
}

func (pc HpComputer) PrintDetails() {
	log.Printf("[HP] Pc Cpu: [%d], Memory: [%d]", pc.Cpu, pc.Memory)
}
