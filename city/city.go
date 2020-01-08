package city

import (
	"fmt"
	"math/rand"
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
		if alien.nbrOfMoves < 100 {
			return true
		}
	}
	return false
}

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

func (cityMap CityMap) Invade(nbrOfAliens int) {
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

	for checkMoveCount(aliens) && len(aliens) > 1 {
		iterateInvasion(&cityMap.Cities, &aliens)
	}
}
