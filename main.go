package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Aragon Challenge"
	app.Usage = "Solution to the Aragon challenge."
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %q", c.Args().Get(0))
		return nil
	}
}

func main() {
	info()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
