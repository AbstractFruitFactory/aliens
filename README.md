# aliens

Simulation of an alien invasion!

Reads a map file (see example.txt) that specifies cities and roads to other cities in different directions (north, east, south, west).

Choose how many aliens should invade. Each alien will spawn in a random city, and then travel random directions to other cities. If two aliens meet in a city, they start a war and kill each other, which in turn destroys the city as well.

Run the project:

`go build`

`./aliens <numberOfAliens> <nameOfMapFile>`

Run all tests:

`go test ./... -v`
