package main

import (
	"fmt"
	"time"
)

func testTime() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

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

var java, python, c int = 1, 2, 3

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

const PI = 3.14

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	var valJava, valGolan = true, "SI!"
	fmt.Println(java, python, c, valGolan, valJava)
	var name string
	var age int64 = 29
	var gender uint8 = 0
	name = "omar" + " barra"
	var price float64 = 3.3
	//explicit
	var programmer bool = true
	//implicit
	var normal = false
	fmt.Printf("%T", normal)

	nickname := "obarra"
	fmt.Println(nickname)

	fmt.Println("Hola Mundo", "in Golang", name, age, gender, price, programmer, "%T", normal)

	var out = fmt.Sprintf("Number: \n \t %07d the best", 404)

	fmt.Println(out)

	var number float32 = 1
	var otherNumber float32 = 3
	result := number / otherNumber
	fmt.Printf("%g", result)

	var numberint int = 1
	var otherNumberInt int = 3
	resultInt := numberint / otherNumberInt
	fmt.Printf("%d", resultInt)

	isTrue := (false || true) && !false && 10-5 == 5
	fmt.Printf("%T", isTrue)

	var herAge = 29
	if herAge >= 29 {
		fmt.Println("Major than 29")
	} else if herAge >= 14 {
		fmt.Println("Major than 14")
	} else {
		fmt.Println("Major than menor than 14")
	}

	if der := "golan"; der == "golan" {
		fmt.Println("This is GOLAN")
	} else {
		fmt.Println("this is JAVA")
	}

	for x := 0; x <= 10; x++ {
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)
	}

	environment := "test"
	switch {
	case environment == "test", environment == "dev":
		fmt.Println("no prod")
	case environment == "prod":
		fmt.Println("prod")
	default:
		fmt.Println("no case")
	}

	switch environment {
	case "test", "dev":
		fmt.Println("no prod")
	case "prod":
		fmt.Println("prod")
	default:
		fmt.Println("no case")
	}

	arr := [3]int{1, 2, 3}
	for d := 0; d < len(arr); d++ {
		fmt.Println(arr[d])
	}

	arr2D := [3][2]int{{1, 2}, {4, 3}, {5, 6}}
	for d := 0; d < len(arr2D); d++ {
		fmt.Println(arr2D[d])
	}

	arraWithOutLimit := []int{3, 4}
	arraWithOutLimit = append(arraWithOutLimit, 10, 12, 13)

	slice := arraWithOutLimit[2:4]
	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	maked := make([]int, 5)
	fmt.Printf("%T", maked)
	fmt.Println(maked)

	for i, element := range arr {
		fmt.Println(i, element)
	}

	for _, element := range arr {
		fmt.Println(element)
	}

	myMap := map[string]int{
		"apple":  5,
		"orange": 3,
		"pear":   1,
	}

	fmt.Println(myMap)
	fmt.Println(myMap["pear"])
	myMap["pear"] = 10
	val, ok := myMap["pear"]
	fmt.Println(val, ok)
	delete(myMap, "pear")
	vald, okd := myMap["pear"]
	fmt.Println(vald, okd)

	added, subs, mul := performNumbers(4, 2)
	fmt.Println(added, subs, mul)

	neg := func(x int) int {
		return x * -1
	}
	fmt.Println(neg(9))
	ioc(neg)

	add := func(x int) int {
		return x + 1
	}
	ioc(add)

	returnFunc("maru")()

	//change both so it is mutable
	var x []int = []int{3, 4, 5}
	//the same pointer
	y := x
	y[0] = 100

	fmt.Println(x, y)

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

	fmt.Println(PI)

	testTime()

	//Nil slices
	var nilSlices []string
	fmt.Println(nilSlices, len(nilSlices), cap(nilSlices))

	if nilSlices == nil {
		fmt.Println(nilSlices)
	}
}

func performNumbers(x, y int) (w, z, d int) {
	defer fmt.Println("it is finally")
	w = x + y
	z = x - y
	d = x * x
	return
}

func ioc(myfunc func(int) int) {
	fmt.Println(myfunc(4))
}

func returnFunc(word string) func() {
	var newWord = "Hi" + word
	return func() {
		fmt.Println(newWord)
	}
}
