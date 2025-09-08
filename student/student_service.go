package student

import (
	"fmt"
	"main/utils"
	"strings"
)

var studentList []Student

func StudentManagement() {

	for {
		utils.ClearScreen()
		fmt.Println("-----------------Student Management-----------------")
		fmt.Println("1.Add Student")
		fmt.Println("2.View Students")
		fmt.Println("3.Update Student")
		fmt.Println("4.Delete Student")
		fmt.Println("5.Back to Main Menu")

		choice := utils.GetPositiveIntInput("Enter your choice: ")

		switch choice {
		case 1:
			utils.ClearScreen()
			addStudent()
		case 2:
			viewStudents()
		case 3:
			updateStudent()
		case 4:
			deleteStudent()
		case 5:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		utils.ReadInput("Press Enter to continue...")
	}
}

func addStudent() {
	name := utils.GetNonEmptyString("Enter student name: ")
	class := utils.GetNonEmptyString("Enter student class: ")
	scoresCount := utils.GetPositiveIntInput("Enter number of scores: ")
	var scores []float64
	for i := 1; i <= scoresCount; i++ {
		score := utils.GetPositiveFloatInput(fmt.Sprintf("Enter score %d: ", i))
		scores = append(scores, score)
	}
	s := Student{
		Name:   name,
		Class:  class,
		Scores: scores,
	}
	stt := New(s)
	studentList = append(studentList, *stt)
	fmt.Println("Student added successfully:", *stt)
}
func viewStudents() {
	utils.ClearScreen()
	fmt.Println("Viewing all students...")
	if len(studentList) == 0 {
		fmt.Println("No data")
		return
	}
	for _, student := range studentList {
		student.GetStudentInfo()
	}
}
func updateStudent() {
	utils.ClearScreen()

	keyword := utils.GetNonEmptyString("Input name of student: ")

	for _, student := range studentList {
		if strings.Contains(student.Name, keyword) {
			name := utils.ReadInput("Input new name or Press enter to skip: ")
			class := utils.ReadInput("Input new class or Press Enter to skip: ")
			scoresCount := utils.GetPositiveIntInput("Input new number of scores: ")
			var scores []float64
			for i := 1; i <= scoresCount; i++ {
				score := utils.GetPositiveFloatInput(fmt.Sprintf("Old Score: %.2f | Enter score %d: ", student.Scores[i-1], i))
				scores = append(scores, score)
			}
			newInfo := Student{Name: name,
				Class: class, Scores: scores}

			student.Update(newInfo)
		}
	}
}
func deleteStudent() {
	fmt.Println("Deleting a student...")
	// Implementation for deleting a student
}
