package invasion

import (
	"fmt"
	"math/rand"
	"time"
)

var randomGenerator *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

type CityMap struct {
	CityNames   []string
	MajorCities map[string]struct{}
	Cities      map[string]*City
}

func NewCityMap() CityMap {
	return CityMap{[]string{}, map[string]struct{}{}, map[string]*City{}}
}

func SetRandomSeed(seed int64) {
	randomSource := rand.NewSource(seed)
	randomGenerator = rand.New(randomSource)
}

func (cityMap *CityMap) Invade(nbrOfAliens int) {
	aliens := cityMap.assignInvaders(nbrOfAliens)
	for checkMoveCount(aliens) && len(aliens) > 1 {
		cityMap.iterateInvasion(&aliens)
	}
	fmt.Print("\nInvasion has concluded!\n\nRemains of civilization:\n\n")
	cityMap.printMap()
}

func (cityMap *CityMap) addCity(city *City) {
	cityMap.Cities[city.Name] = city
	cityMap.CityNames = append(cityMap.CityNames, city.Name)
}

func (cityMap *CityMap) addMajorCity(name string) {
	cityMap.MajorCities[name] = struct{}{}
}

func (cityMap *CityMap) removeCity(cityName string) {
	cityMap.Cities[cityName].Destroyed = true
	for index, name := range cityMap.CityNames {
		if name == cityName {
			cityMap.CityNames = removeItem(cityMap.CityNames, index)
		}
	}

}

func (cityMap CityMap) hasCity(name string) bool {
	_, ok := cityMap.Cities[name]
	return ok
}

func (cityMap CityMap) getCity(name string) *City {
	return cityMap.Cities[name]
}

/*
Creates nbrOfAliens aliens and assigns them to random cities.
*/
func (cityMap *CityMap) assignInvaders(nbrOfAliens int) Aliens {
	aliens := NewAliensList(nbrOfAliens)

	for i := 0; i < nbrOfAliens; i++ {
		var startingCity *City
		alien := &Alien{i, startingCity, 0}

		randomNbr := randomGenerator.Intn(len(cityMap.Cities))

		startingCity = cityMap.Cities[cityMap.CityNames[randomNbr]]
		startingCity.Invaders[i] = alien

		alien.currentCity = startingCity
		aliens = aliens.add(alien)
	}

	return aliens
}

/*
Moves each alien a random direction to another city (if a road in that direction exists).
Then, checks for every city if it has more than 1 invader present. If so, that city and its invaders are removed.
*/
func (cityMap *CityMap) iterateInvasion(aliens *Aliens) {
	for _, alien := range *aliens {
		if alien == nil {
			continue
		}
		randNbr := randomGenerator.Intn(3)
		alien.move(Direction(randNbr))
	}

	for _, city := range cityMap.Cities {
		if city.Destroyed != true {
			if len(city.Invaders) > 1 {
				fmt.Printf("City %s was destroyed by aliens %v!\n", city.Name, getAlienIds(city.Invaders))
				cityMap.removeCity(city.Name)

				for _, alien := range city.Invaders {
					*aliens = aliens.remove(alien.id)
				}
			}
		}
	}
}

/*
Logs the map in its current state, ignoring all destroyed cities.
*/
func (cityMap CityMap) printMap() {
	var currentCity *City

	for name, _ := range cityMap.MajorCities {
		currentCity = cityMap.Cities[name]

		if currentCity.Destroyed {
			continue
		}

		fmt.Printf("%v ", currentCity.Name)

		if currentCity.North != nil && !currentCity.North.Destroyed {
			fmt.Printf("north=%v ", currentCity.North.Name)
		}
		if currentCity.East != nil && !currentCity.East.Destroyed {
			fmt.Printf("east=%v ", currentCity.East.Name)
		}
		if currentCity.South != nil && !currentCity.South.Destroyed {
			fmt.Printf("south=%v ", currentCity.South.Name)
		}
		if currentCity.West != nil && !currentCity.West.Destroyed {
			fmt.Printf("west=%v ", currentCity.West.Name)
		}

		fmt.Println()
	}
}

func getAlienIds(aliens map[int]*Alien) []int {
	ids := make([]int, len(aliens))

	i := 0
	for k := range aliens {
		ids[i] = k
		i++
	}
	return ids
}

func removeItem(arr []string, index int) []string {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
