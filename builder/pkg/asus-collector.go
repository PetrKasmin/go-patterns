package pkg

type AsusCollector struct {
	Core          int
	Brand         string
	Memory        int
	Monitor       int
	GraphicalCard int
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *AsusCollector) SetGraphicalCard() {
	collector.GraphicalCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:          collector.Core,
		Brand:         collector.Brand,
		Memory:        collector.Memory,
		GraphicalCard: collector.GraphicalCard,
	}
}
