package invasion

type City struct {
	Name     string
	Invaders map[int]*Alien
	North    *City
	East     *City
	South    *City
	West     *City
}
