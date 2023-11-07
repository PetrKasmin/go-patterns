package pkg

import (
	"log"
)

type Computer struct {
	Core          int
	Brand         string
	Memory        int
	Monitor       int
	GraphicalCard int
}

func (c Computer) Print() {
	log.Printf(
		"%s Core:[%d], Memory: [%d], Monitor: [%d], GraphicalCard: [%d]\n",
		c.Brand,
		c.Core,
		c.Memory,
		c.Monitor,
		c.GraphicalCard,
	)
}
