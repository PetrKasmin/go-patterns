package main

import "bridge/pkg/types"

var (
	hpScanner    = types.Hp{}
	epsonScanner = types.Epson{}

	macPc     = types.MacPc{}
	linuxPc   = types.LinuxPc{}
	windowsPc = types.WindowsPc{}
)

func main() {
	linuxPc.AddScanner(&hpScanner)
	linuxPc.Scan()

	linuxPc.AddScanner(&epsonScanner)
	linuxPc.Scan()

	macPc.AddScanner(&hpScanner)
	macPc.Scan()

	windowsPc.AddScanner(&epsonScanner)
	windowsPc.Scan()
}
