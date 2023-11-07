package pkg

import (
	"log"
	"time"
)

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (c Card) CheckBalance() error {
	log.Println("[Карта] Запрос в банк для проверки остатка")
	time.Sleep(time.Millisecond * 800)
	return c.Bank.CheckBalance(c.Name)
}
