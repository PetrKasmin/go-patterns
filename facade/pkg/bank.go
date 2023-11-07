package pkg

import (
	"errors"
	"log"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (b Bank) CheckBalance(cardNumber string) error {
	log.Printf("[Банк] Получение остатка по карте %s\n", cardNumber)
	time.Sleep(time.Millisecond * 300)

	for _, card := range b.Cards {
		if card.Name != cardNumber {
			continue
		}

		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}

	log.Println("[Банк] Остаток положительный")
	return nil
}
