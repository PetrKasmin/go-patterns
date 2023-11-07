package types

import (
	"bridge/pkg/base"
	"log"
)

type LinuxPc struct {
	scanner base.Scanner
}

func (pc *LinuxPc) Scan() {
	log.Println("Scan pc.go linux system")
	pc.scanner.ScanFile()
}

func (pc *LinuxPc) AddScanner(scanner base.Scanner) {
	pc.scanner = scanner
}
