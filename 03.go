package main

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	Aluno := Aluno{
		nome:  "fulano",
		idade: 20,
		notas: []float64{1, 2, 3, 4, 5, 6},
	}

	calcularMediaAluno(Aluno)
}

func calcularMediaAluno(a Aluno) float64 {
	var somatorio float64
	for _, nota := range a.notas {
		somatorio += nota
	}
	return somatorio / float64(len(a.notas))
}
