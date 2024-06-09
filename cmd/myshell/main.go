package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var builtIns map[string]func(args []string)

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
	fn, f := builtIns[command]

	if f {
		fn(args)
	} else {
		cmd := exec.Command(command, args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			if errors.Is(err, exec.ErrNotFound) {
				fmt.Printf("%s: command not found\n", command)
			} else {
				log.Println(err)
			}
		}
	}
}

func main() {
	builtIns = map[string]func(args []string){
		"exit": exit,
		"echo": echo,
		"type": typeBuiltIn,
		"cd":   cd,
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		line := readLine()
		cmd, args := splitLine(line)
		run(cmd, args)
	}
}
