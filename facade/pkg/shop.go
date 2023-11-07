package pkg

import (
	"errors"
	"log"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	log.Println("[Магазин] Запрос к пользователью, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)

	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	log.Printf("[Магазин] Проверка - может ли %s пользователь купить товар\n", user.Name)

	for _, p := range shop.Products {
		if p.Name != product {
			continue
		}

		if p.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средст для покупки товара")
		}

		log.Printf("[Магазин] Товар %s куплен\n", p.Name)
	}

	return nil
}
