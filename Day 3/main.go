package main

import (
	"alterra-agmc-day3/app"
	"alterra-agmc-day3/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
