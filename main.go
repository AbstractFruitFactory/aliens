package main

import (
	"errors"
	"github.com/alexmedkex/aliens/city"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

var app = cli.NewApp()

func setInfo() {
	app.Name = "Alien invasion simulator"
	app.Usage = "\n\nSimulate an alien invasion!\n1. Specify the city map in a separate .txt file in the current folder.\n2. Run the program with number of alien invaders as argument."
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		input := c.Args().Get(0)
		nbrOfAliens, err := strconv.Atoi(input)

		if err != nil || nbrOfAliens < 1 {
			panic(errors.New("Invalid input. Must be a positive integer."))
		}

		cityMap := city.BuildCityMap("./example.txt")
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
