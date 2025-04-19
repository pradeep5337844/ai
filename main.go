package main

import (
	"fmt"
	inventory "inventory_management/inventory"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Inventory Management System")
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [command] [arguments]")
		fmt.Println("Usage: available commands: add, view, update, delete")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go add name quantity")
			return
		}
		name := os.Args[2]
		quantity, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid quantity:", err)
			return
		}
		inventory.AddItem(inventory.Item{ID: len(inventory.Inventory) + 1, Name: name, Quantity: quantity})

	case "view":
		inventory.ViewInventory()

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go update itemID newQuantity")
			return
		}
		itemID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid item ID:", err)
			return
		}
		newQuantity, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid quantity:", err)
			return
		}
		inventory.UpdateQuantity(itemID, newQuantity)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go delete itemID")
			return
		}
		itemID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid item ID:", err)
			return
		}
		inventory.DeleteItem(itemID)

	default:
		fmt.Println("Unknown command:", command)
	}
}
