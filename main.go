package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "beer-mug"
	app.Usage = "measures your works. I help your beer happy drinking!!"

	app.Commands = []cli.Command{
		initCommand(),
		startCommand(),
		endCommand(),
		statusCommand(),
	}

	app.Run(os.Args)
}
