package pkg

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{
		Collector: collector,
	}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetBrand()
	factory.Collector.SetMemory()
	factory.Collector.SetMonitor()
	factory.Collector.SetGraphicalCard()
	return factory.Collector.GetComputer()
}
