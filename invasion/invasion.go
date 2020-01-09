package invasion

import (
	"bufio"
	"os"
	"strings"
)

/*
Builds a city map from a .txt file, located at filePath.
*/
func BuildCityMap(filePath string) CityMap {
	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	cityMap := NewCityMap()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, " ")
		cityName := lineData[0]
		var neighborNorth *City
		var neighborEast *City
		var neighborSouth *City
		var neighborWest *City

		for i := 1; i < len(lineData); i++ {
			neighborData := strings.Split(lineData[i], "=")
			direction := neighborData[0]
			neighborCityName := neighborData[1]

			var neighborCity *City

			if !cityMap.HasCity(neighborCityName) {
				cityMap.AddCity(NewCity(neighborCityName))
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
			newCity := NewCity(cityName)
			newCity.North = neighborNorth
			newCity.East = neighborEast
			newCity.South = neighborSouth
			newCity.West = neighborWest
			cityMap.AddCity(newCity)
		} else {
			cityMap.Cities[cityName].North = neighborNorth
			cityMap.Cities[cityName].East = neighborEast
			cityMap.Cities[cityName].South = neighborSouth
			cityMap.Cities[cityName].West = neighborWest
		}
	}

	return cityMap
}

func checkMoveCount(aliens Aliens) bool {
	for _, alien := range aliens {
		if alien != nil {
			if alien.nbrOfMoves < 10000 {
				return true
			}
		}
	}
	return false
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
