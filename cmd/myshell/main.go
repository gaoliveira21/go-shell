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
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd := readLine()

		fmt.Printf("%s: command not found\n", cmd)
	}
}
