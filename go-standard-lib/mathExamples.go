package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MathExamples() {
	rand.Seed(time.Now().UnixMicro())
	fmt.Println("Math")
	fmt.Println(rand.Intn(100))
}
