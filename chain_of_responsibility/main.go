package main

import "chain_of_responsibility/pkg"

func main() {
	device := &pkg.Device{
		Name: "Device-1",
	}

	updateService := &pkg.UpdateDataService{
		Name: "Update-1",
	}

	dataService := &pkg.DataService{}

	device.SetNext(updateService)
	updateService.SetNext(dataService)

	data := &pkg.Data{}
	device.Execute(data)
}
