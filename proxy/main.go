package main

import (
	"log"
	"proxy/pkg"
)

var (
	admin = "admin"
	user  = "user"

	users = map[string]bool{
		admin: true,
		user:  false,
	}
)

func main() {
	proxy := pkg.ProxyDatabase{
		Users: users,
		Db:    &pkg.Database{},
	}

	data, err := proxy.GetData(admin)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(data)

	data, err = proxy.GetData(user)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(data)
}
