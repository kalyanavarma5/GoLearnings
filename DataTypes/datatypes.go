package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	dataTypes()
}

func dataTypes() {
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)
	//Type conversion
	helo := 80             //int
	plo := 91.8            //float64
	sum := helo + int(plo) //int + float64 not allowed
	fmt.Println(sum)
}
