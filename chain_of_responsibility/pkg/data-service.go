package pkg

import "log"

type DataService struct {
	Name string
	Next Service
}

func (u *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		log.Println("Data not update")
		u.Next.Execute(data)
		return
	}

	log.Println("Data save")
}

func (u *DataService) SetNext(svc Service) {
	u.Next = svc
}
