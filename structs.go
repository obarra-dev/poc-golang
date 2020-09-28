package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type shape interface {
	getDoubleRadius() float32
	setDoubleRadius(redius float32)
}

type Circle struct {
	center *Point
	redius float32
}

type CircleFlatten struct {
	*Point
	redius float32
}

func (c Circle) getDoubleRadius() float32 {
	return c.redius * 2
}

func (c *Circle) setDoubleRadius(radius float32) {
	c.redius = radius * 2
}

func (c CircleFlatten) getDoubleRadius() float32 {
	return c.redius * 2
}

func (c *CircleFlatten) setDoubleRadius(radius float32) {
	c.redius = radius * 2
}

func xmain() {

	circle := Circle{&Point{y: 4, x: 9}, 3.3}
	fmt.Println(circle.center.x, circle.center.y)

	circleFlatten := CircleFlatten{&Point{6, 7}, 5.5}
	fmt.Println(circleFlatten.x, circleFlatten.y)

	fmt.Println("first", circle.getDoubleRadius())

	circle.setDoubleRadius(7)
	fmt.Println(circle.getDoubleRadius())

	shapes := []shape{&circle, &circleFlatten}

	for _, shape := range shapes {
		fmt.Println(shape.getDoubleRadius())
	}

}
