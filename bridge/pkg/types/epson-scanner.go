package types

import "log"

type Epson struct {
}

func (s *Epson) ScanFile() {
	log.Println("Epson scan file")
}
