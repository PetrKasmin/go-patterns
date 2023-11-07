package pkg

type HpFactory struct {
}

func (factory *HpFactory) GetComputer() Computer {
	return HpComputer{
		Memory: 16,
		Cpu:    6,
	}
}

func (factory *HpFactory) GetMonitor() Monitor {
	return HpMonitor{
		Size: 24,
	}
}
