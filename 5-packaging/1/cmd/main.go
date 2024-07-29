package main

import (
	"curso-go/full-cycle/5-packaging/1/math"
	"fmt"
)

func main() {
	m := math.Math{A: 3, B: 2}
	fmt.Println(m.Add())
}

// replace no go:
// import "curso-go/full-cycle/5-packaging/1/math"
// por:
// import "../math"
// no terminal:
// go mod edit -replace=curso-go/full-cycle/5-packaging/1/math=../math
// essa solução é paia, pois o caminho é relativo ao meu sistema
// quando eu for compartilhar o código, o caminho não vai funcionar