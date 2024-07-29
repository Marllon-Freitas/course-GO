package main

import (
	"curso-go/full-cycle/5-packaging/1/math"
	"fmt"
)

func main() {
	m := math.Math{A: 3, B: 2}
	fmt.Println(m.Add())
}
