package invasion

import (
	"fmt"
	"math/rand"
	"time"
)

var randomGenerator *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type CityMap struct {
	CityNames []string
	Cities    map[string]*City
}

func NewCityMap() CityMap {
	return CityMap{[]string{}, map[string]*City{}}
}

func SetRandomSeed(seed int64) {
	randomSource := rand.NewSource(seed)
	randomGenerator = rand.New(randomSource)
}

func (cityMap *CityMap) AddCity(city *City) {
	cityMap.Cities[city.Name] = city
	cityMap.CityNames = append(cityMap.CityNames, city.Name)
}

func (cityMap *CityMap) RemoveCity(cityName string) {
	delete(cityMap.Cities, cityName)
	for index, name := range cityMap.CityNames {
		if name == cityName {
			cityMap.CityNames = removeItem(cityMap.CityNames, index)
		}
	}
}

func (cityMap CityMap) HasCity(name string) bool {
	_, ok := cityMap.Cities[name]
	return ok
}

func (cityMap CityMap) GetCity(name string) *City {
	return cityMap.Cities[name]
}

func (cityMap *CityMap) Invade(nbrOfAliens int) {
	aliens := cityMap.assignInvaders(nbrOfAliens)
	for checkMoveCount(aliens) && len(aliens) > 1 {
		cityMap.iterateInvasion(&aliens)
	}
}

/*
Creates nbrOfAliens aliens and assigns them to random cities.
*/
func (cityMap *CityMap) assignInvaders(nbrOfAliens int) map[int]*Alien {
	aliens := map[int]*Alien{}

	for i := 0; i < nbrOfAliens; i++ {
		var startingCity *City
		alien := &Alien{i, startingCity, 0}

		randomNbr := randomGenerator.Intn(len(cityMap.Cities))

		startingCity = cityMap.Cities[cityMap.CityNames[randomNbr]]
		startingCity.Invaders[i] = alien

		alien.currentCity = startingCity
		aliens[i] = alien
	}
	return aliens
}

/*
Moves each alien a random direction to another city (if a road in that direction exists).
Then, checks for every city if it has more than 1 invader present. If so, that city and present invaders are destroyed.
*/
func (cityMap *CityMap) iterateInvasion(aliens *map[int]*Alien) {
	for _, alien := range *aliens {
		randNbr := randomGenerator.Intn(3)
		alien.move(Direction(randNbr))
	}

	for _, city := range cityMap.Cities {
		if len(city.Invaders) > 1 {
			fmt.Printf("City %s was destroyed by aliens %v!\n", city.Name, city.Invaders)
			cityMap.RemoveCity(city.Name)

			for _, alien := range city.Invaders {
				delete(*aliens, alien.id)
			}
		}
	}
}

func removeItem(arr []string, index int) []string {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
