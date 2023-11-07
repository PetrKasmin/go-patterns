package main

import (
	"observer/pkg"
)

func main() {
	sub := &pkg.Subscriber{
		Name: "Sub-1",
	}

	sub2 := &pkg.Subscriber{
		Name: "Sub-2",
	}

	sub3 := &pkg.Subscriber{
		Name: "Sub-3",
	}

	subN := &pkg.Subscriber{
		Name: "Sub-n",
	}

	ch := pkg.Publisher{
		Name:      "Publisher chanel",
		Consumers: map[string]pkg.Consumer{},
	}

	ch.Subscribe(sub)
	ch.Subscribe(sub2)
	ch.Subscribe(sub3)
	ch.Subscribe(subN)

	ch.UnSubscribe(sub2)
	ch.Notify()
}
