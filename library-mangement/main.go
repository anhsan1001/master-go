package librarymangement

import (
	"fmt"

	"main/utils"
)

func Menu() {
	lib := NewLibrary()
	for {
		utils.ClearScreen()

		fmt.Println("---------------QUAN LY THU VIEN---------------")
		fmt.Println("1. Them sach")
		fmt.Println("2. Xem danh sach")
		fmt.Println("3. Them nguoi muon")
		fmt.Println("4. Xem danh sach nguoi muon")
		fmt.Println("5. Muon sach")
		fmt.Println("6. Xem lich su muon")
		fmt.Println("7. Tra sach")
		fmt.Println("8. Tim kiem sach")
		fmt.Println("9. Thoat")
		choice := utils.GetPositiveIntInput("Chon chuc nang: ")
		utils.ClearScreen()

		switch choice {
		case 1:
			if err := AddBook(lib); err != nil {
				fmt.Printf("Loi!: %v\n", err)
			}
		case 2:
			if err := ListBook(lib); err != nil {
				fmt.Printf("Loi!: %v\n", err)
			}
		case 3:

			if err := AddBorrower(lib); err != nil {

				fmt.Printf("Loi!: %v\n", err)
			}
		case 4:

			if err := ListBorrower(lib); err != nil {

				fmt.Printf("Loi!: %v\n", err)
			}
		case 5:
			if err := BorrowerBook(lib); err != nil {

				fmt.Printf("Loi!: %v\n", err)
			}
		case 6:

			if err := ListBorrowerHistory(lib); err != nil {

				fmt.Printf("Loi!: %v\n", err)
			}
		case 7:

			if err := ReturnBook(lib); err != nil {

				fmt.Printf("Loi!: %v\n", err)
			}
		case 8:

			if err := SearchBook(lib); err != nil {

				fmt.Printf("Loi!: %v\n", err)
			}
		case 9:
			return
		}

		utils.ReadInput("Press Enter to continue...")
	}
}
