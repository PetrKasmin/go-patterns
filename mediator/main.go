package main

import (
	"mediator/pkg"
	"time"
)

func main() {
	stationManager := pkg.NewStationManager()
	passengerBas := pkg.Passenger{
		Dispatcher: stationManager,
	}

	cargo := pkg.Cargo{
		Dispatcher: stationManager,
	}

	passengerBas.Arrive()
	time.Sleep(time.Second * 1)
	cargo.Arrive()
	time.Sleep(time.Second * 1)
	passengerBas.Go()
}
