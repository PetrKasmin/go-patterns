package pkg

import "fmt"

type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (v *HasMoneyState) RequestItem() error {
	return fmt.Errorf("Item dispense in progress")
}

func (v *HasMoneyState) AddItem(i int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (v *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (v *HasMoneyState) DispenseItem() error {
	fmt.Errorf("Dispense Item")
	v.vendingMachine.itemCount = v.vendingMachine.itemCount - 1
	if v.vendingMachine.itemCount == 0 {
		v.vendingMachine.setState(v.vendingMachine.noItem)
	} else {
		v.vendingMachine.setState(v.vendingMachine.hasItem)
	}
	return nil
}
