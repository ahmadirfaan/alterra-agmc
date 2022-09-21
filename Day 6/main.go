package main

import (
	"alterra-agmc-day6/app"
	"alterra-agmc-day6/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
