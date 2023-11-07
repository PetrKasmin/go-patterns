package main

import (
	"abstract_factory/pkg"
	"log"
)

var brands = []string{pkg.Asus, pkg.Hp, "Dell"}

func main() {
	for _, b := range brands {
		factory, err := pkg.GetFactory(b)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		monitor := factory.GetMonitor()
		monitor.PrintDetails()

		computer := factory.GetComputer()
		computer.PrintDetails()
	}
}
