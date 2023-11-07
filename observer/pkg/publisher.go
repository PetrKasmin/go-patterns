package pkg

type Publisher struct {
	Name      string
	Consumers map[string]Consumer
}

func (p *Publisher) Subscribe(consumer Consumer) {
	p.Consumers[consumer.GetName()] = consumer
}

func (p *Publisher) UnSubscribe(consumer Consumer) {
	delete(p.Consumers, consumer.GetName())
}

func (p *Publisher) Notify() {
	for _, c := range p.Consumers {
		c.Update(p.Name)
	}
}
