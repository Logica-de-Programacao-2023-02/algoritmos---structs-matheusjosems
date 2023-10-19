package main

import (
	"fmt"
)

//Crie uma struct chamada Pessoa com os campos "nome", "idade" e
//"endereço". O campo "endereço" deve ser uma outra struct com os
//campos "rua", "número", "cidade" e "estado". Escreva uma função
//que receba uma Pessoa como parâmetro e imprima seu endereço completo.

type Pessoa struct {
	nome     string
	idade    int
	endereço Endereço
}
type Endereço struct {
	rua    string
	numero int
	cidade string
	estado string
}

func ImprimirEndereço(pessoa Pessoa) {
	fmt.Println("nome:", pessoa.nome)
	fmt.Println("idade:", pessoa.idade)
	fmt.Println("endereço:", pessoa.endereço)
	fmt.Println("rua:", pessoa.endereço.rua)
	fmt.Println("numero:", pessoa.endereço.numero)
	fmt.Println("cidade:", pessoa.endereço.cidade)
	fmt.Println("estado:", pessoa.endereço.estado)

}
func main() {

	y := Endereço{
		rua:    "lago norte",
		numero: 14,
		cidade: "Brasilia",
		estado: "DF",
	}
	x := Pessoa{
		nome:     "matheus",
		idade:    22,
		endereço: y,
	}

	ImprimirEndereço(x)

}
