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

func TestAdd(t *testing.T) {
	city := NewCity("TestCity")
	alien := Alien{1, city, 0}

	aliens := NewAliensList(2)

	aliens = aliens.add(&alien)

	if *aliens[alien.id] != alien {
		t.Errorf("Expected alien %v to be in aliens list, got %v", alien.id, aliens[alien.id].id)
	}

	if len(aliens) != 2 {
		t.Errorf("Expected length of aliens list to be %v, got %v", 2, len(aliens))
	}
}

func TestRemove(t *testing.T) {
	city := NewCity("TestCity")
	alien1 := Alien{0, city, 0}
	alien2 := Alien{1, city, 0}

	aliens := NewAliensList(2)

	aliens = aliens.add(&alien1)
	aliens = aliens.add(&alien2)

	countBefore := aliens.count()

	aliens = aliens.remove(0)

	countAfter := aliens.count()

	if countAfter != countBefore-1 {
		t.Errorf("Expected alien count to be %v after removing, got %v.", countBefore-1, countAfter)
	}

	if *aliens[1] != alien2 {
		t.Errorf("Expected aliens list to have alien %v after removal, got %v.", alien2.id, aliens[1].id)
	}

	if aliens[0] != nil {
		t.Errorf("Expected alien 0 to be nil, got alien with id %v.", aliens[0].id)
	}
}
