package main

import (
	"alterra-agmc-day7/app"
	"alterra-agmc-day7/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
