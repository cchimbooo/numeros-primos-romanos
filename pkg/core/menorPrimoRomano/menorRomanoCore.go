package menorPrimoRomano

import (
	"log"
	"prova-programacao/pkg/helpers/leitor"
	"prova-programacao/pkg/helpers/romanos"
)

// NOVAMENTE: Não precisaria ser uma interface, só para mostrar que sei como fazer.

type coreMenorRomano struct {
	mapa map[string]int
}

// ProcessarString string implementa a lógixa paara descobrior se é o menor primo romano
// e satisfaz o ProcessarStringInterface
func (c coreMenorRomano) ProcessarString(texto string) MenorNumero {
	numerosRomanos := leitor.StringParaSequenciaDeRomanos(texto)
	// novamente, estou com a mania (por causa da instrução do meu trabalho atual) de evitar usar o range]

	menor := MenorNumero{
		Romano:  "",
		Decimal: 4000,
	}
	for i := 0; i < len(numerosRomanos); i++ {
		// verifica o número decimal do número romano.
		if numero, existe := c.mapa[numerosRomanos[i]]; existe {
			if numero < menor.Decimal {
				menor.Decimal = numero
				menor.Romano = numerosRomanos[i]
			}
		}
	}
	if menor.Decimal == 4000 {
		menor.Decimal = 0
	}
	return menor
}

// CriarProcessarStringInterface é uma pequena implementação de uma futura arquitetura exagonal, poderia receber tudo
// como plug and play (a função de converter para romanos, a função de gerar os primos, a função de leitura entre outros)
// eu acabei não fazendo assim pois acredito que é extremamente uma "overkill", a única razão de eu implementar uma
// interface foi para mostrar que sei como fazer, estou recebendo só os primos pois é algo que depende de um processamento
// prévio.
// Novamente: Se fosse fazer isso em um ambiente profissional ou para mim eu não implementaria nenhuma interface nem uma.
// arquitetura dividida em pastas, acredito que um projeto deste tamanho deveria ser feito tudo em uma única pasta.
// estou fazendo asism pois é um teste.
func CriarProcessarStringInterface(primos []int) ProcessarStringInterface {
	// percorre os primos mapeando os mesmos para um mapa de [primosRomanos]primosInteiro
	mapa := make(map[string]int, len(primos))
	for i := 0; i < len(primos); i++ {
		// Converte o número
		numeroRomano, err := romanos.Gerar(primos[i])
		if err != nil {
			// Não tem um tratamento melhor, se der erro tem que dar um panic mesmos.
			log.Fatal(err)
		}
		// Atribui ao mapa os valores romanos dos primos
		mapa[numeroRomano] = primos[i]
	}
	return coreMenorRomano{mapa: mapa}
}
