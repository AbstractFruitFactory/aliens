package invasion

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func checkMoveCount(aliens map[int]*Alien) bool {
	for _, alien := range aliens {
		if alien.nbrOfMoves < 10000 {
			return true
		}
	}
	return false
}

/*
Moves each alien a random direction to another city (if a road in that direction exists).
Then, checks for every city if it has more than 1 invader present. If so, that city and present invaders are destroyed.
*/
func iterateInvasion(cities *map[string]*City, aliens *map[int]*Alien) {
	for _, alien := range *aliens {
		direction := rand.Intn(3)
		alien.move(direction)
	}

	for _, city := range *cities {
		if len(city.Invaders) > 1 {
			fmt.Printf("City %s was destroyed by aliens %v!\n", city.Name, city.Invaders)
			delete(*cities, city.Name)
			for _, alien := range city.Invaders {
				delete(*aliens, alien.id)
			}
		}
	}
}

func BuildCityMap(filePath string) CityMap {
	file, err := os.Open(filePath)
	checkErr(err)
	defer file.Close()

	cityMap := CityMap{}
	cityMap.Cities = map[string]*City{}
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
				cityMap.Cities[neighborCityName] = &City{Name: neighborCityName}
				cityMap.Cities[neighborCityName].Invaders = map[int]*Alien{}
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
			cityMap.Cities[cityName] = &City{
				Name:  cityName,
				North: neighborNorth,
				East:  neighborEast,
				South: neighborSouth,
				West:  neighborWest,
			}
			cityMap.Cities[cityName].Invaders = map[int]*Alien{}
		} else {
			cityMap.Cities[cityName].North = neighborNorth
			cityMap.Cities[cityName].East = neighborEast
			cityMap.Cities[cityName].South = neighborSouth
			cityMap.Cities[cityName].West = neighborWest
		}
	}

	return cityMap
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
