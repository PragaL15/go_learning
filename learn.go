package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter input: ")
	input, _ := reader.ReadString('\n')

	// Trim the newline character
	input = strings.TrimSpace(input)

	// Convert input string to a float64
	numRat, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Added 1 to your rating:", numRat+1)
	}
}
