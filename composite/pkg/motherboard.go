package pkg

import "log"

type Motherboard struct {
	Name        string
	Description string
	Components  []Component
}

func (m *Motherboard) Search(name string) {
	if m.Name == name {
		log.Printf("Component: [%s] is find %s", m.Name, m.Description)
	}

	for _, component := range m.Components {
		component.Search(name)
	}
}

func (m *Motherboard) GetName() string {
	return m.Name
}
