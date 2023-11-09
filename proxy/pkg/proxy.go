package pkg

import "errors"

type ProxyDatabase struct {
	Users map[string]bool
	Db    *Database
}

func (p *ProxyDatabase) GetData(user string) ([]string, error) {
	if !p.Users[user] {
		return nil, errors.New("Forbidden")
	}
	return p.Db.GetData(user)
}
