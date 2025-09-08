package librarymangement

import (
	"fmt"
	"main/utils"
)

func Menu() {
	for {
		utils.ClearScreen()
		fmt.Println("---------------LIBRARY MANAGEMENT---------------")
		fmt.Println("1. Add book")
		fmt.Println("2. Get Book List")
		fmt.Println("3. Add book renter")
		fmt.Println("4. Get book renter list")
		fmt.Println("5. Rent book")
		fmt.Println("6. Rent book history")
		fmt.Println("7. Return book")
		fmt.Println("8. Find book")
		fmt.Println("9. Exit")
		choice := utils.GetPositiveIntInput("Enter choose feature: ")

		switch choice {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("1")
		case 3:
			fmt.Println("1")
		case 4:
			fmt.Println("1")
		case 5:
			fmt.Println("1")
		case 6:
			fmt.Println("1")
		case 7:
			fmt.Println("1")
		case 8:
			fmt.Println("1")
		case 9:
			return
		}
	}
}
