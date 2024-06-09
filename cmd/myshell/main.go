package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var builtIns map[string]func(args []string)

func exit(args []string) {
	code := 0
	if len(args) > 0 {
		c, err := strconv.Atoi(args[0])
		if err == nil {
			code = c
		}
	}

	os.Exit(code)
}

func echo(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func typeBuiltIn(args []string) {
	cmd := args[0]

	for k := range builtIns {
		if k == cmd {
			fmt.Printf("%s is a shell builtin\n", cmd)
			return
		}
	}

	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, ":")

	for _, path := range paths {
		entries, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		for _, e := range entries {
			if !e.IsDir() && e.Name() == cmd {
				fmt.Printf("%s is %s\n", cmd, fmt.Sprintf("%s/%s", path, cmd))
				return
			}
		}
	}

	fmt.Printf("%s: not found\n", cmd)
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
	}

	for {
		fmt.Fprint(os.Stdout, "$ ")
		line := readLine()
		cmd, args := splitLine(line)
		run(cmd, args)
	}
}
