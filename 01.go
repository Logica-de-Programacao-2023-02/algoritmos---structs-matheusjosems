package main

import (
	"fmt"
	"math"
)

//Crie uma struct chamada Circulo com o campo "raio".
//Escreva uma função que receba um Circulo como parâmetro
//e calcule a área do círculo (área = pi * raio * raio).
//Dica: use a constante math.Pi para representar o número pi.

type Circulo struct {
	raio float64
}

func CalcularAreaDoCirculo(circulo Circulo) float64 {
	return math.Pi * circulo.raio * circulo.raio
}

func main() {
	X := Circulo{raio: 4.0}
	area := CalcularAreaDoCirculo(X)
	fmt.Println("a area do circulo :", area)
}
