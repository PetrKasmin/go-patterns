package pkg

import "log"

type Cpu struct {
	Name        string
	Description string
}

func (c *Cpu) Search(name string) {
	if c.Name == name {
		log.Printf("Component: [%s] is find %s", c.Name, c.Description)
	}
}

func (c *Cpu) GetName() string {
	return c.Name
}
