package main

import (
	"decorator/pkg"
	"log"
)

var (
	basePc = pkg.BasePc{}
	homePc = pkg.HomePc{
		Cpu:           4,
		GraphicalCard: 1,
		Wrapper:       basePc,
	}
	serverPc = pkg.ServerPc{
		Cpu:     12,
		Memory:  256,
		Wrapper: basePc,
	}
)

func main() {
	log.Println(basePc.GetPrice(), homePc.GetPrice(), serverPc.GetPrice())
}
