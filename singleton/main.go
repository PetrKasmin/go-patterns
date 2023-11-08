package main

import (
	"log"
	"singleton/pkg"
)

var singleton *pkg.Singleton

func init() {
	log.Println("Singleton initialisation")
	singleton = &pkg.Singleton{
		Type: "Singleton",
	}
}

func main() {
	for range []int{1, 2, 3} {
		singleton = pkg.NewSingleton(singleton, "Create Singleton")
	}
}
