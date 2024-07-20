package main

// não reclama de não ser usada pois é um escopo global
// var(
// 	a bool		= true
// 	b int			= 1
// 	c string	= "Hello, World!"
// 	d float64	= 1.0
// )

func main() {
	// outra forma de declarar variáveis
	f := 1 // dessa forma o compilador infere o tipo da variável e ele já atribui o valor
	// f := 2 da erro pois a variável já foi declarada
	println(f)
}
