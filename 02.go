package main

import "fmt"

type Livro struct {
	titulo string
	autor  string
	ano    int
}

func main() {
	l := Livro{
		titulo: "livro",
		autor:  "autor",
		ano:    2023,
	}
	imprimeLivro(l)

}

func imprimeLivro(l Livro) {
	fmt.Println("titulo:", l.titulo)
	fmt.Println("autor:", l.autor)
	fmt.Println("ano:", l.ano)
}