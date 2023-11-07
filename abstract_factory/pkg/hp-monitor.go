package pkg

import "log"

type HpMonitor struct {
	Size int
}

func (pc HpMonitor) PrintDetails() {
	log.Printf("[HP] Monitor Size: [%d]", pc.Size)
}
