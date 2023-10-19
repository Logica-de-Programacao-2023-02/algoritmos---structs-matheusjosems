package main

//Crie uma struct chamada Playlist com os campos "nome" e "músicas".
//O campo "músicas" deve ser um slice de structs, cada uma representando
//uma música com os campos "título", "artista" e "duração". Escreva uma
//função que receba uma Playlist como parâmetro e imprima o nome da
//playlist, o tempo total da playlist e as informações de cada música.

type Playlist struct {
	nume    string
	Musicas []musica
}
type musica struct {
	titulo  string
	artista string
	duraçao float64
}

func main() {

}

func Playlistinfo(playlist Playlist) {

}
