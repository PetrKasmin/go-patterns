package types

import "log"

type Hp struct {
}

func (s *Hp) ScanFile() {
	log.Println("Epson scan file")
}
