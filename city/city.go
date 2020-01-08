package city

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type CityMap struct {
	Cities map[string]*City
}

type City struct {
	Name     string
	Invaders map[int]*Alien
	North    *City
	East     *City
	South    *City
	West     *City
}

type Alien struct {
	id          int
	currentCity *City
	nbrOfMoves  int
}

func (cityMap CityMap) HasCity(name string) bool {
	_, ok := cityMap.Cities[name]
	return ok
}

func (cityMap CityMap) GetCity(name string) *City {
	for _, city := range cityMap.Cities {
		if city.Name == name {
			return city
		}
	}
	return nil
}

/*
Moves an alien in a direction to another city, if present.
*/
func (alien *Alien) move(direction int) {
	var destination *City

	switch direction {
	case 0:
		destination = alien.currentCity.North
	case 1:
		destination = alien.currentCity.East
	case 2:
		destination = alien.currentCity.South
	case 3:
		destination = alien.currentCity.West

	}

	if destination != nil {
		destination.Invaders[alien.id] = alien
		delete(alien.currentCity.Invaders, alien.id)
		alien.currentCity = destination
	}
	alien.nbrOfMoves++
}

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

/*
Creates nbrOfAliens aliens and assigns them to random cities.
*/
func assignInvaders(nbrOfAliens int, cityMap CityMap) map[int]*Alien {
	aliens := map[int]*Alien{}

	for i := 0; i < nbrOfAliens; i++ {
		var startingCity *City
		alien := &Alien{i, startingCity, 0}

		for k := range cityMap.Cities {
			startingCity = cityMap.Cities[k]
			break
		}

		startingCity.Invaders[i] = alien

		alien.currentCity = startingCity
		aliens[i] = alien
	}
	return aliens
}

func (cityMap CityMap) Invade(nbrOfAliens int) {
	aliens := assignInvaders(nbrOfAliens, cityMap)
	for checkMoveCount(aliens) && len(aliens) > 1 {
		iterateInvasion(&cityMap.Cities, &aliens)
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
