package main

import "composite/pkg"

var (
	cpu = pkg.Cpu{
		Name:        "Cpu-1",
		Description: "Процессор-1",
	}

	cpu2 = pkg.Cpu{
		Name:        "Cpu-2",
		Description: "Процессор-2",
	}

	card = pkg.GraphicsCard{
		Name:        "Radeon",
		Description: "Видеокарта-1",
	}

	card2 = pkg.GraphicsCard{
		Name:        "GeForce",
		Description: "Видеокарта-2",
	}

	motherboard = pkg.Motherboard{
		Name:        "Gigabyte",
		Description: "Материнская плата",
		Components: []pkg.Component{
			&cpu,
			&cpu2,
			&card,
			&card2,
		},
	}

	pc = pkg.Pc{
		Name:        "PC",
		Description: "Компьютер",
		Components: []pkg.Component{
			&motherboard,
		},
	}
)

func main() {
	pc.Search("Gigabyte")
	pc.Search("Radeon")
	pc.Search("Cpu-2")
}
