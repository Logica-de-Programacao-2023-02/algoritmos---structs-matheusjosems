package main

import "fmt"

type Retangulo struct {
	Largura int
	Altura  int
}

func main() {
	r := Retangulo{Largura: 5, Altura: 10}
	resultado := areaRetangulo(r)
	fmt.Println(resultado)
}

func areaRetangulo(r Retangulo) int {
	return r.Largura * r.Altura

}
