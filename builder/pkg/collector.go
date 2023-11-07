package pkg

const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicalCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}
