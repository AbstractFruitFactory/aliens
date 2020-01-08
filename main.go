package main

import (
	"bufio"
	"errors"
	"github.com/alexmedkex/aliens/city"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"strings"
)

var app = cli.NewApp()

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

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

		cityMap := buildCityMap("./example.txt")
		cityMap.Invade(nbrOfAliens)

		return nil
	}
}

func buildCityMap(filePath string) city.CityMap {
	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	cityMap := city.CityMap{}
	cityMap.Cities = map[string]*city.City{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, " ")
		cityName := lineData[0]
		var neighborNorth *city.City
		var neighborEast *city.City
		var neighborSouth *city.City
		var neighborWest *city.City

		for i := 1; i < len(lineData); i++ {
			neighborData := strings.Split(lineData[i], "=")
			direction := neighborData[0]
			neighborCityName := neighborData[1]

			var neighborCity *city.City

			if !cityMap.HasCity(neighborCityName) {
				cityMap.Cities[neighborCityName] = &city.City{Name: neighborCityName}
				cityMap.Cities[neighborCityName].Invaders = map[int]*city.Alien{}
			}

			neighborCity = cityMap.Cities[neighborCityName]

			switch direction {
			case "north":
				neighborNorth = neighborCity
			case "east":
				neighborEast = neighborCity
			case "south":
				neighborSouth = neighborCity
			case "west":
				neighborWest = neighborCity
			}
		}

		if !cityMap.HasCity(cityName) {
			cityMap.Cities[cityName] = &city.City{
				Name:  cityName,
				North: neighborNorth,
				East:  neighborEast,
				South: neighborSouth,
				West:  neighborWest,
			}
			cityMap.Cities[cityName].Invaders = map[int]*city.Alien{}
		} else {
			cityMap.Cities[cityName].North = neighborNorth
			cityMap.Cities[cityName].East = neighborEast
			cityMap.Cities[cityName].South = neighborSouth
			cityMap.Cities[cityName].West = neighborWest
		}
	}

	return cityMap
}

func main() {
	setInfo()
	err := app.Run(os.Args)
	checkErr(err)
}
