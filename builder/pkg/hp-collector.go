package pkg

type HpCollector struct {
	Core          int
	Brand         string
	Memory        int
	Monitor       int
	GraphicalCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Memory = 1
}

func (collector *HpCollector) SetGraphicalCard() {
	collector.GraphicalCard = 2
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:          collector.Core,
		Brand:         collector.Brand,
		Memory:        collector.Memory,
		GraphicalCard: collector.GraphicalCard,
	}
}
