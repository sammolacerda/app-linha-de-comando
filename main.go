package main

import (
	"aplicacoa-linha-de-comando/app"
	"fmt"

	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()

	if erro := aplicacao.Run(os.Args);erro != nil{
		log.Fatal(erro)
	}
}