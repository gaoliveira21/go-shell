package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func exit(args []string) {
	code, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(code)
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	return strings.TrimSpace(line)
}

func splitLine(line string) (string, []string) {
	args := strings.Fields(line)

	return args[0], args[1:]
}

func run(command string, args []string) {
	switch command {
	case "exit":
		exit(args)
	default:
		fmt.Printf("%s: command not found\n", command)
	}
}

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		line := readLine()
		cmd, args := splitLine(line)
		run(cmd, args)
	}
}
