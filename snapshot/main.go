package main

import (
	"log"
	"snapshot/pkg"
)

func main() {
	storage := &pkg.Guardian{
		Items: make([]*pkg.Snapshot, 0),
	}

	creator := &pkg.Creator{
		State: "A",
	}
	log.Printf("Creator current State %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("B")
	log.Printf("Creator current State %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.SetState("C")
	log.Printf("Creator current State %s\n", creator.GetState())
	storage.Add(creator.Create())

	creator.Restore(storage.Get(1))
	log.Printf("Restored to State %s\n", creator.GetState())

	creator.Restore(storage.Get(0))
	log.Printf("Restored to State %s\n", creator.GetState())
}
