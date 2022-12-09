package leitor

import "strings"

func verificaCaractereRomano(r rune) bool {
	return r == 'I' || r == 'V' || r == 'X' || r == 'C' || r == 'L' || r == 'M' || r == 'D'
}

//StringParaSequenciaDeRomanos recebe um texto e extrai treechos de números que pode ser romanos.
// ATENÇÃO: A função não extrai números romanos válidos, apenas sequencia de caracteres que podem ser um número romano
// dado o problema e a solução proposta optei por fazer isso, na próxima função deve ser comparado o retorno dessa
// com romanos primos, assim os números "errados" não vão estar no mapa de primos e consequentemente vão ser ignorados
func StringParaSequenciaDeRomanos(texto string) []string {
	// Poderiamos usar isso como o core e fazer o tratamento comentado para deixar mais rápido, mas acredito
	// que não valha o incomodo
	// Teria que receber um map[string]int e retornar um inteiro;

	// menor := 0

	// Array de strings para retornar
	retorno := []string{}

	// Determina se estamos numa sequencia de números romanos
	sequencia := false

	// Escreve o número
	textoAtual := strings.Builder{}

	// Converte o texto em runas, evitando problema com index fora de UTF-8,
	// ultimamente estou com a mania de iterar usando o for em vez do range (uma vez li e fiz uns testes e normalmente
	// é mais performático).
	textoRunas := []rune(texto)

	// Le cada runa.
	for i := 0; i < len(textoRunas); i++ {
		// Se for uma caractére romano:
		if verificaCaractereRomano(textoRunas[i]) {
			// Ativa a sequencia se a mesma não existe
			if !sequencia {
				sequencia = true
				textoAtual.WriteRune(textoRunas[i])
				continue
			}
			// se existe simplesmente escreve
			textoAtual.WriteRune(textoRunas[i])
			continue
			// Se não for um caractére romano:
		} else {
			// Se sequencia estiver vazia simplesmente ignora
			if !sequencia {
				continue
			}
			// Caso esteja em uma sequencia de romanos marca sequencia como falsa, salva o texto escrito no retorno
			// e limpa o buffer.
			sequencia = false
			retorno = append(retorno, textoAtual.String())
			// Aqui viria a lógica do menor.
			// Comentario pois não quero fazer assim, ia ficar muito confuso, só para mostrar aonde daria para por
			// if menor == 0 || menor > MapaRomanoIntPrimo[retorno] {
			//   menor = MapaRomanoIntPrimo[retorno]
			// }
			textoAtual.Reset()
		}
	}
	// Após iterar se os ultimos caracteres forem romanos adiciona o mesmo ao retorno
	if sequencia {
		// aqui teria a coisa do menor também
		retorno = append(retorno, textoAtual.String())
	}
	// return menor

	return retorno

}
