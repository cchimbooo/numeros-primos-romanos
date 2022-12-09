package romanos

import (
	"errors"
	"strings"
)

// Gerar converte um inteiro para um número romano.
// A função trabalha apenas com M = 1000,D = 500, C = 100, L = 50, X = 10, V = 5 e I = 1
// Consequentemente o maior número aceito é 3999
func Gerar(numero int) (string, error) {
	// Valida se é um número até 3999
	if numero > 3999 || numero < 1 {
		return "", errors.New("numero inválido, o mesmo deve estar entre 1 e 3999")
	}
	
	// resultado criado num builder
	romanoBuilder := strings.Builder{}

	// 2 arrays que representam os números romanos e os casts para inteiros (mapa não garantem ordem)
	// Os valors 900, 400, 40, 9 e 4 são os únicos valores distintos em romano, que não apresentam a regra de "somar"
	// Em vez de fazer uma lógica de botar na direita ou esquerda é mais fácil mapear os mesmos
	var valores = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var representacao = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// faz um index para controlar por fora posição
	index := 0
	// Enquanto o número não for 0 roda
	for numero > 0 {
		// Enquanto o número for menor que o número da posição escreve o número e subtrai
		for valores[index] <= numero {
			// Escreve a representação romana
			romanoBuilder.WriteString(representacao[index])
			// Subtrai o valor
			numero -= valores[index]
		}
		// Saindo  do for (while) reduz o index e passa novamente no próximo número
		index += 1
	}

	return romanoBuilder.String(), nil
}
