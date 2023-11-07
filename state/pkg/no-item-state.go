package pkg

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (v *NoItemState) AddItem(i int) error {
	v.vendingMachine.incrementItemCount(i)
	v.vendingMachine.setState(v.vendingMachine.hasItem)
	return nil
}

func (v *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Item out of stock")
}

func (v *NoItemState) DispenseItem() error {
	return fmt.Errorf("Item out of stock")
}

func (v *NoItemState) RequestItem() error {
	return fmt.Errorf("Item out of stock")
}
