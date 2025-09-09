package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPositiveIntInput(prompt string) int {

	for {
		input := ReadInput(prompt)
		value, err := strconv.Atoi(input)
		if err != nil || value <= 0 {
			fmt.Println("Invalid input. Please enter a positive integer.")
			continue
		}
		return value
	}
}
func GetPositiveFloatInput(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil || value <= 0 {
			fmt.Println("Invalid input. Please enter a positive number.")
			continue
		}
		return value
	}
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func GetNonEmptyString(prompt string) string {

	for {
		input := ReadInput(prompt)

		if input != "" {
			return input
		}
		fmt.Println("Not allowed empty string ")
	}
}

func GenerateID() string {
	return uuid.New().String()
}
