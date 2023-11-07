package pkg

import "log"

type UpdateDataService struct {
	Name string
	Next Service
}

func (u *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		log.Printf("Data from device %s already update", u.Name)
		u.Next.Execute(data)
		return
	}

	log.Printf("Update data from device %s", u.Name)
	data.UpdateSource = true
	u.Next.Execute(data)
}

func (u *UpdateDataService) SetNext(svc Service) {
	u.Next = svc
}
