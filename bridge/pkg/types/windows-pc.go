package types

import (
	"bridge/pkg/base"
	"log"
)

type WindowsPc struct {
	scanner base.Scanner
}

func (pc *WindowsPc) Scan() {
	log.Println("Scan pc.go windows system")
	pc.scanner.ScanFile()
}

func (pc *WindowsPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
