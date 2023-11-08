package pkg

import "log"

type GraphicsCard struct {
	Name        string
	Description string
}

func (g *GraphicsCard) Search(name string) {
	if g.Name == name {
		log.Printf("Component: [%s] is find %s", g.Name, g.Description)
	}
}

func (g *GraphicsCard) GetName() string {
	return g.Name
}
