package pkg

import "fmt"

type HasItemState struct {
	vendingMachine *VendingMachine
}

func (v *HasItemState) RequestItem() error {
	if v.vendingMachine.itemCount == 0 {
		v.vendingMachine.setState(v.vendingMachine.noItem)
		return fmt.Errorf("No item present")
	}

	fmt.Printf("Item request\n")
	v.vendingMachine.setState(v.vendingMachine.itemRequest)
	return nil
}

func (v *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	v.vendingMachine.incrementItemCount(count)
	return nil
}

func (v *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first")
}

func (v *HasItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first")
}
