package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func cd(args []string) {
	path := strings.Join(args, "")

	if err := os.Chdir(path); err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}
}
