package pkg

import "log"

type AsusMonitor struct {
	Size int
}

func (pc AsusMonitor) PrintDetails() {
	log.Printf("[Asus] Monitor Size: [%d]", pc.Size)
}
