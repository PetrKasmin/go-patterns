package pkg

import "log"

type Pc struct {
	Name        string
	Description string
	Components  []Component
}

func (p *Pc) Search(name string) {
	if p.Name == name {
		log.Printf("Component: [%s] is find %s", p.Name, p.Description)
	}

	for _, component := range p.Components {
		component.Search(name)
	}
}

func (p *Pc) GetName() string {
	return p.Name
}
