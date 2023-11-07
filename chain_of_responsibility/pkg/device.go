package pkg

import "log"

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		log.Printf("Data from device %s already get", d.Name)
		d.Next.Execute(data)
		return
	}

	log.Printf("Get data from device %s", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(svc Service) {
	d.Next = svc
}
