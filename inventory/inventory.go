package inventory

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantity int
}

var Inventory []Item

func AddItem(item Item) {
	Inventory = append(Inventory, item)
	fmt.Println("added item to inventory %v", item)
}
func ViewInventory() {
	if len(Inventory) == 0 {
		fmt.Println("the inventory is empty")
		return
	}
	fmt.Println("\n -----Inventory-----")
	for _, item := range Inventory {
		fmt.Printf("ID: %d , Name: %s, Quantity: %d\n", item.ID, item.Name, item.Quantity)
	}
}

func UpdateQuantity(itemID, newQuantity int) {
	for i, item := range Inventory {
		if item.ID == itemID {
			Inventory[i].Quantity = newQuantity
			fmt.Printf("Updated item %d to new quantity %d\n", itemID, newQuantity)
			return
		}
	}
	fmt.Println("Item not found")
}

func DeleteItem(itemID int) {
	for i, item := range Inventory {
		if item.ID == itemID {
			Inventory = append(Inventory[:i], Inventory[i+1:]...)
			fmt.Printf("Deleted item %d\n", itemID)
			return
		}
	}
	fmt.Println("Item not found")
}
