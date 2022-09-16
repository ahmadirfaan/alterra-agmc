package main

import (
	"alterra-agmc-day4/app"
	"alterra-agmc-day4/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
