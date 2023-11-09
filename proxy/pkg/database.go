package pkg

type Database struct {
}

func (d *Database) GetData(user string) ([]string, error) {
	return []string{
		"String-1",
		"String-2",
	}, nil
}
