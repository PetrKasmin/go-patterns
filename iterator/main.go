package main

import (
	"iterator/pkg"
	"log"
)

var routes = pkg.Routes{
	Routes: []pkg.Route{
		{
			Name:       "Route-1",
			TravelTime: 110,
		},
		{
			Name:       "Route-2",
			TravelTime: 50,
		},
		{
			Name:       "Route-3",
			TravelTime: 60,
		},
		{
			Name:       "Route-4",
			TravelTime: 40,
		},
	},
}

func main() {
	for routes.HasNext() {
		route := routes.GetNext()
		log.Println(route.Name, route.TravelTime)
	}
}
