package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "jottings"
	app.Usage = "note taking in terminal"
	app.Action = func(c *cli.Context) error {
		CreateFile("")
		return nil
	}
	app.Run(os.Args)
}
