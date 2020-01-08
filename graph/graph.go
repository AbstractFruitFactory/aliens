package graph

import (
	"math/rand"
)

type City struct {
	Name  string
	North *City
	East  *City
	South *City
	West  *City
}

func (cityMap CityMap) HasCity(name string) bool {
	_, ok := cityMap.Cities[name]
	return ok
}


type CityMap struct {
	Cities map[string]*City
}
