package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func MathExamples() {

	p := fmt.Println

	// We need to have a seed for our random
	// Or we will get the same random every time
	rand.Seed(time.Now().UnixMicro())
	p(rand.Intn(100))

	// Finding the absolute value of a number
	p(math.Abs(-300))

	// Find the square root of a float64
	p(math.Sqrt(64))

	// Rounding a float 64
	p(math.Round(10.5))

}
