package main

import (
	"github.com/urfave/cli"
	"time"
)

func startCommand() cli.Command {
	return cli.Command{
		Name: "start",
		Action: runStart,
	}
}

func runStart(*cli.Context) error {
	if appHome, err := getAppHome(); err != nil {
		return err;
	} else {
		return startTimer(appHome, time.Now())
	}
}