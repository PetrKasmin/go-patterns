package main

import "strategy/pkg"

var (
	start      = 10
	end        = 100
	strategies = []pkg.Strategy{
		&pkg.WalkStrategy{},
		&pkg.RoadStrategy{},
		&pkg.PublicTransportStrategy{},
	}
)

func main() {
	nav := pkg.Navigator{}
	for _, s := range strategies {
		nav.SetStrategy(s)
		nav.Route(start, end)
	}
}
