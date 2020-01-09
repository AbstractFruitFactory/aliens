package invasion

import "testing"

func TestMove(t *testing.T) {
	city := NewCity("TestCity")
	city2 := NewCity("TestCity2")
	city.North = city2
	alien := Alien{1, city, 0}

	if alien.currentCity != city {
		t.Errorf("Expected current city to be %v, instead got %v.", city.Name, alien.currentCity.Name)
	}

	alien.move(NORTH)

	if alien.currentCity != city2 {
		t.Errorf("Expected current city to be %v, instead got %v.", city2.Name, alien.currentCity.Name)
	}

	alien.move(EAST)

	if alien.currentCity != city2 {
		t.Errorf("Expected current city to be %v, instead got %v.", city2.Name, alien.currentCity.Name)
	}

	if alien.nbrOfMoves != 2 {
		t.Errorf("Number of moves should be %v, instead got %v.", 2, alien.nbrOfMoves)
	}
}
