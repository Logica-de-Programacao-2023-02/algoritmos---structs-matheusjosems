package main

import "fmt"

//Crie uma struct chamada Triângulo com os campos "base" e "altura".
//Escreva uma função que receba um Triângulo como parâmetro e calcule
//a área do triângulo (área = base * altura / 2).

type triangulo struct {
	base   float64
	altura float64
}

func CalcularAreaDoTriangulo(triangulo triangulo) float64 {
	return (triangulo.base * triangulo.altura) / 2
}

func main() {
	x := triangulo{
		base:   5.0,
		altura: 20.0,
	}
	area := CalcularAreaDoTriangulo(x)
	fmt.Println("area do treiangulo:", area)
}
