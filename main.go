package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// This is a placeholder for the main function.
	var name string

	fmt.Println("Nhap ten cua ban: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name = scanner.Text()
	}
	fmt.Printf("Hello, %s!\n", name)
}
