package main

import (
	"fmt"
	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type Commands struct {
	cmd map[string]func(*state, command) error
}

func (c *Commands) run(s *state, cmd command) error {
	cmdFunc, exists := c.cmd[cmd.name]
	if exists == false {
		return fmt.Errorf("Command does not exist")
	}
	return cmdFunc(s, cmd)
}

func (c *Commands) register(name string, f func(*state, command) error) {
	_, exists := c.cmd[name]
	if exists == false {
		c.cmd[name] = f
	}
	return
}

func loginHandler(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments provided")
	}

	if len(cmd.args) > 1 {
		return fmt.Errorf("Too many arguments")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Set user %s in config file\n", s.cfg.User)

	return nil
}
