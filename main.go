package main

import (
	"fmt"
	"main/student"
	"main/teacher"
	"main/utils"
)

func main() {
	for {
		utils.ClearScreen()

		fmt.Println("-----------------Main Menu-----------------")
		fmt.Println("1.Student Management")
		fmt.Println("2.Teacher Management")
		fmt.Println("3.Exit")

		choice := utils.GetPositiveIntInput("Enter your choice: ")

		switch choice {

		case 1:
			student.StudentManagement()
		case 2:
			teacher.TeacherManagement()
		case 3:
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
