package main

import (
	"errors"
	"github.com/alexmedkex/aliens/invasion"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

var app = cli.NewApp()

func setInfo() {
	app.Name = "Alien invasion simulator"
	app.Usage = "\n\nSimulate an alien invasion!\n1. Specify the city map in a separate .txt file in the current folder.\n2. Run the program with number of alien invaders and file name as arguments."
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		arg1 := c.Args().Get(0)
		arg2 := c.Args().Get(1)

		nbrOfAliens, err := strconv.Atoi(arg1)

		if err != nil || nbrOfAliens < 1 {
			panic(errors.New("Invalid input. Must be a positive integer."))
		}

		cityMap := invasion.BuildCityMap("./" + arg2)
		cityMap.Invade(nbrOfAliens)

		return nil
	}
}

func main() {
	setInfo()
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
