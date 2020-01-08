package invasion

type Alien struct {
	id          int
	currentCity *City
	nbrOfMoves  int
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
