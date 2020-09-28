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

func main() {

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

	testTime()

}
