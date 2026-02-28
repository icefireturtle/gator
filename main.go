package main

import (
	"fmt"
	"os"

	"gator/internal/config"
)

func main() {
	fmt.Println("Welcome to Gator!")
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config file: %s", err)
		return
	}

	s := state{
		cfg: &cfg,
	}

	cmds := Commands{
		cmd: make(map[string]func(*state, command) error),
	}

	cmds.register("login", loginHandler)

	if len(os.Args) < 2 {
		fmt.Printf("need program name and argument\n")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Printf("error running command: %s\n", err)
		os.Exit(1)
	}

}
