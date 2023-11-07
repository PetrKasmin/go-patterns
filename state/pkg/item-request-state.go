package pkg

import "fmt"

type ItemRequestState struct {
	vendingMachine *VendingMachine
}

func (v *ItemRequestState) RequestItem() error {
	return fmt.Errorf("Item already request")
}

func (v *ItemRequestState) AddItem(i int) error {
	return fmt.Errorf("Item dispense in progress")
}

func (v *ItemRequestState) InsertMoney(money int) error {
	if v.vendingMachine.itemPrice > money {
		return fmt.Errorf("Insert money is less. Please insert %d", v.vendingMachine.itemPrice)
	}

	fmt.Println("Insert money is ok")
	v.vendingMachine.setState(v.vendingMachine.hasMoney)
	return nil
}

func (v *ItemRequestState) DispenseItem() error {
	return fmt.Errorf("Please insert money first")
}
