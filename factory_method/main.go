package main

import "factory_method/pkg"

var types = []string{pkg.ServerType, pkg.NotebookType, pkg.PersonalComputerType, "monoblock"}

func main() {
	for _, t := range types {
		computer := pkg.New(t)
		if computer == nil {
			continue
		}

		computer.PrintDetails()
	}
}
