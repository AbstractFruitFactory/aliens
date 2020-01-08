package main

import (
	"bufio"
	"errors"
	"github.com/alexmedkex/aliens/graph"
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
	app.Name = "Tendermint Challenge"
	app.Usage = "Solution to the Tendermint challenge."
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		input := c.Args().Get(0)
		nbrOfAliens, err := strconv.Atoi(input)

		if err != nil || nbrOfAliens < 1 {
			panic(errors.New("Invalid input. Must be an positive integer."))
		}

		return nil
	}
}

func buildCityMap(filePath string) graph.CityMap {
	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	cityMap := graph.CityMap{}
	cityMap.Cities = map[string]*graph.City{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, " ")
		cityName := lineData[0]
		var neighborNorth *graph.City
		var neighborEast *graph.City
		var neighborSouth *graph.City
		var neighborWest *graph.City

		for i := 1; i < len(lineData); i++ {
			neighborData := strings.Split(lineData[i], "=")
			direction := neighborData[0]
			neighborCityName := neighborData[1]

			var neighborCity *graph.City

			if !cityMap.HasCity(neighborCityName) {
				cityMap.Cities[neighborCityName] = &graph.City{Name: neighborCityName}
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
			cityMap.Cities[cityName] = &graph.City{
				Name:  cityName,
				North: neighborNorth,
				East:  neighborEast,
				South: neighborSouth,
				West:  neighborWest,
			}
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
	//cityMap := buildCityMap("./example.txt")
	setInfo()

	err := app.Run(os.Args)
	checkErr(err)
}
