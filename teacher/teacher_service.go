package teacher

import (
	"fmt"
	"main/utils"
)

var teacherList []Teacher

func TeacherManagement() {

	for {
		utils.ClearScreen()
		fmt.Println("-----------------Teacher Management-----------------")
		fmt.Println("1.Add Teacher")
		fmt.Println("2.View Teachers")
		fmt.Println("3.Update Teacher")
		fmt.Println("4.Delete Teacher")
		fmt.Println("5.Back to Main Menu")

		choice := utils.GetPositiveIntInput("Enter your choice: ")

		switch choice {
		case 1:
			utils.ClearScreen()
			addTeacher()
		case 2:
			viewTeachers()
		case 3:
			updateTeacher()
		case 4:
			deleteTeacher()
		case 5:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		utils.ReadInput("Press Enter to continue...")
	}
}

func addTeacher() {
	name := utils.GetNonEmptyString("Enter teacher name: ")
	subject := utils.GetNonEmptyString("Enter teacher subject: ")
	numberOfClasses := utils.GetPositiveIntInput("Enter number of class: ")

	var classes []string
	for i := 1; i <= numberOfClasses; i++ {
		class := utils.GetNonEmptyString(fmt.Sprintf("Enter class name %d: ", i))
		classes = append(classes, class)
	}
	t := Teacher{
		Name:    name,
		Class:   classes,
		Subject: subject,
	}
	tc := New(t)
	teacherList = append(teacherList, *tc)
	fmt.Println("Student added successfully: ", *tc)
}
func viewTeachers() {
	if len(teacherList) == 0 {
		fmt.Println("No data")
		return
	}
	for _, teacher := range teacherList {

		teacher.GetTeacherInfo()
	}
}
func updateTeacher() {
	fmt.Println("Updating a teacher...")
	// Implementation for updating a teacher
}
func deleteTeacher() {
	fmt.Println("Deleting a teacher...")
	// Implementation for deleting a teacher
}
