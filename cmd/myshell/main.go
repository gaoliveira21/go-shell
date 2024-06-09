package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLine() string {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(line)
}

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	cmd := readLine()

	fmt.Printf("%s: command not found", cmd)
}
