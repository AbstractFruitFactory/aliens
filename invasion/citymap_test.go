package invasion

import "testing"

var cityMap = BuildCityMap("../example.txt")

func TestAssignInvaders(t *testing.T) {
	SetRandomSeed(1)
	aliens := cityMap.assignInvaders(2)

	if aliens[0].currentCity.Name != "Baz" {
		t.Errorf("Expected currentCity: Baz, got currentCity: %v.", aliens[0].currentCity.Name)
	}
}
