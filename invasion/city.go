package invasion

type City struct {
	Name      string
	Invaders  map[int]*Alien
	Destroyed bool
	North     *City
	East      *City
	South     *City
	West      *City
}

func NewCity(name string) *City {
	return &City{Name: name, Invaders: map[int]*Alien{}}
}
