package main

import (
	"github.com/urfave/cli"
)

func statusCommand() cli.Command {
	return cli.Command{
		Name: "status",
		Action: runStatus,
	}
}

func runStatus(c *cli.Context) error {
	return runInitLook(c)
}