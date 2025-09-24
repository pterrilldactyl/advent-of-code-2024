// My First Go project!

package main

import "fmt"

var foo int = 17

type coordinate struct {
	x int
	y int
}

func newCoordinate(x int, y int) *coordinate {
	return &coordinate{x, y}
}

func (c *coordinate) calculateArea() int {
	return c.x * c.y
}

func main() {
	bar := 17
	fmt.Println("Hello, world!")
	fmt.Printf("My number is: %d \n", add(foo, bar))

	coord := coordinate{}
	coord.x = 17
	coord.y = 17
	fmt.Printf("My coodinates is: %+v \n", coord)

}

func add(a int, b int) int {
	return a + b
}
