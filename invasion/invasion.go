package invasion

import (
	"bufio"
	"os"
	"strings"
)

/*
Builds a city map from a .txt file, located at filePath.

Cities that have roads leading to other cities, i.e listed in the leftmost column in the file, are
saved as major cities.
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
		majorCityName := lineData[0]

		city := &City{}
		var neighborNorth *City
		var neighborEast *City
		var neighborSouth *City
		var neighborWest *City

		if cityMap.hasCity(majorCityName) {
			city = cityMap.getCity(majorCityName)
		} else {
			city = NewCity(majorCityName)
			cityMap.addCity(city)
		}

		cityMap.addMajorCity(majorCityName)

		for i := 1; i < len(lineData); i++ {
			neighborData := strings.Split(lineData[i], "=")
			direction := neighborData[0]
			neighborCityName := neighborData[1]

			var neighborCity *City

			if cityMap.hasCity(neighborCityName) {
				neighborCity = cityMap.getCity(neighborCityName)
			} else {
				neighborCity = NewCity(neighborCityName)
				cityMap.addCity(neighborCity)
			}

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

		city.North = neighborNorth
		city.East = neighborEast
		city.South = neighborSouth
		city.West = neighborWest

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
