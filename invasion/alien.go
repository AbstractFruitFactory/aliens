package invasion

type Alien struct {
	id          int
	currentCity *City
	nbrOfMoves  int
}

/*
Moves an alien in a direction to another city, if present.
*/
func (alien *Alien) move(direction Direction) {
	var destination *City

	switch direction {
	case NORTH:
		destination = alien.currentCity.North
	case EAST:
		destination = alien.currentCity.East
	case SOUTH:
		destination = alien.currentCity.South
	case WEST:
		destination = alien.currentCity.West

	}

	if destination != nil {
		destination.Invaders[alien.id] = alien
		delete(alien.currentCity.Invaders, alien.id)
		alien.currentCity = destination
	}
	alien.nbrOfMoves++
}
