package invasion

type CityMap struct {
	Cities map[string]*City
}

func (cityMap CityMap) HasCity(name string) bool {
	_, ok := cityMap.Cities[name]
	return ok
}

func (cityMap CityMap) GetCity(name string) *City {
	return cityMap.Cities[name]
}

func (cityMap CityMap) Invade(nbrOfAliens int) {
	aliens := cityMap.assignInvaders(nbrOfAliens)
	for checkMoveCount(aliens) && len(aliens) > 1 {
		iterateInvasion(&cityMap.Cities, &aliens)
	}
}

/*
Creates nbrOfAliens aliens and assigns them to random cities.
*/
func (cityMap CityMap) assignInvaders(nbrOfAliens int) map[int]*Alien {
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
