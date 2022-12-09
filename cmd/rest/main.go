package main

import (
	"log"
	"prova-programacao/pkg/core/menorPrimoRomano"
	"prova-programacao/pkg/helpers/primos"
	"prova-programacao/pkg/servers/rest"
)

func main() {
	sievePrimos := primos.Sieve3999()
	romanoInterface := menorPrimoRomano.CriarProcessarStringInterface(sievePrimos)
	if romanoInterface == nil {
		log.Fatalln("interface vazia")
	}
	if err := rest.SetUpRoute(romanoInterface); err != nil {
		log.Fatalln("Falha ao inicializar as rotas")
	}

}
