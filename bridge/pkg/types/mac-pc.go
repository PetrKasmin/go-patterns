package types

import (
	"bridge/pkg/base"
	"log"
)

type MacPc struct {
	scanner base.Scanner
}

func (pc *MacPc) Scan() {
	log.Println("Scan pc.go mac system")
	pc.scanner.ScanFile()
}

func (pc *MacPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
