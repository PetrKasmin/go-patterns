package pkg

import "log"

type Passenger struct {
	Dispatcher
}

func (p *Passenger) PermitArrive() {
	log.Println("Пасажиры: занимайте места...")
	p.Arrive()
}

func (p *Passenger) Go() {
	log.Println("Пасажиры: отправление!")
	p.Dispatcher.NotifyAboutGo()
}

func (p *Passenger) Arrive() {
	if !p.CanArrive(p) {
		log.Println("Пасажиры: отправление задерживается...")
		return
	}
	log.Println("Пасажиры: занимайте места...")
}
