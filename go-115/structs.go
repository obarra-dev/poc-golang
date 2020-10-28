package main

import (
	"math"
)

type Language string

func (l Language) GetName() string {
	return "name: " + string(l)
}

type Point struct {
	x int
	y int
}

func (p Point) Abs() float64 {
	return math.Sqrt(float64(p.x*p.x + p.y*p.y))
}

func (p *Point) Scale(f int) {
	p.x = p.x * f
	p.y = p.y * f
}

type Circle struct {
	center *Point
	redius float32
}

type CircleFlatten struct {
	*Point
	redius float32
}

type shape interface {
	getDoubleRadius() float32
	setDoubleRadius(redius float32)
	hasImplementation() bool
}

func (c Circle) getDoubleRadius() float32 {
	return c.redius * 2
}

func (c *Circle) setDoubleRadius(radius float32) {
	c.redius = radius * 2
}

func (c *Circle) hasImplementation() bool {
	if c == nil {
		return false
	}
	return true
}

func (c CircleFlatten) getDoubleRadius() float32 {
	return c.redius * 2
}

func (c *CircleFlatten) setDoubleRadius(radius float32) {
	c.redius = radius * 2
}

func (c *CircleFlatten) hasImplementation() bool {
	if c == nil {
		return false
	}
	return true
}

//todos los tipos implementan la interfaz vacia
type emptyInterface interface{}

type ComplexStruct struct {
	name      string
	age       int
	nick      *string
	addresses []string
	phones    map[string]string
}

func (c *ComplexStruct) createCopy() ComplexStruct {
	nick := *c.nick
	addresses := make([]string, len(c.addresses), len(c.addresses))
	copy(addresses, c.addresses)

	phones := make(map[string]string)
	for k, v := range c.phones {
		phones[k] = v
	}
	return ComplexStruct{c.name, c.age, &nick, addresses, phones}
}
