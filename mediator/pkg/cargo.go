package pkg

import "log"

type Cargo struct {
	Dispatcher
}

func (c *Cargo) PermitArrive() {
	log.Println("Грузовик: погрузка...")
	c.Arrive()
}

func (c *Cargo) Go() {
	log.Println("Грузовик: отправление!")
	c.Dispatcher.NotifyAboutGo()
}

func (c *Cargo) Arrive() {
	if !c.CanArrive(c) {
		log.Println("Грузовик: отправление задерживается...")
		return
	}
	log.Println("Грузовик: отправлен")
}
