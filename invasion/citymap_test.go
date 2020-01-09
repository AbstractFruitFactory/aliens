package invasion

import (
	"testing"
)

func TestAssignInvaders(t *testing.T) {
	cityMap := setup()

	aliens := cityMap.assignInvaders(2)

	expectedCities := map[int]string{
		0: "Baz",
		1: "Shiz",
	}

	for i := range aliens {
		if aliens[i].currentCity.Name != expectedCities[i] {
			t.Errorf("Expected currentCity: %v, got currentCity: %v.", expectedCities[i], aliens[i].currentCity.Name)
		}
	}

}

func TestIterateInvasion(t *testing.T) {
	cityMap := setup()

	aliens := cityMap.assignInvaders(5)
	expectedCities := map[int]string{
		1: "Bee",
		2: "Waz",
		3: "Pellio",
	}

	nbrOfAliensBefore := aliens.count()
	nbrOfCitiesBefore := len(cityMap.CityNames)

	cityMap.iterateInvasion(&aliens)

	nbrOfAliensAfter := aliens.count()
	nbrOfCitiesAfter := len(cityMap.CityNames)

	if nbrOfAliensAfter != nbrOfAliensBefore-2 {
		t.Errorf("Expected %v number of aliens after iteration, got %v.", nbrOfAliensBefore-2, nbrOfAliensAfter)
	}

	if nbrOfCitiesAfter != nbrOfCitiesBefore-1 {
		t.Errorf("Expected %v number of cities after iteration, got %v.", nbrOfCitiesBefore-1, nbrOfCitiesAfter)
	}

	for i, alien := range aliens {
		if alien == nil {
			continue
		}
		if alien.currentCity.Name != expectedCities[i] {
			t.Errorf("Expected currentCity: %v, got currentCity: %v.", expectedCities[i], alien.currentCity.Name)
		}
	}
}

func setup() CityMap {
	SetRandomSeed(1)
	return newCity()
}

func newCity() CityMap {
	return BuildCityMap("../test_input/test1.txt")
}
