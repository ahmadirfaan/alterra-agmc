package main

import (
	"alterra-agmc-dynamic-crud/app"
	"alterra-agmc-dynamic-crud/cli"
	"os"
)

func main() {
	c := cli.NewCli(os.Args)
	c.Run(app.Init())

}
