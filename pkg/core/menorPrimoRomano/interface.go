package menorPrimoRomano

// ProcessarStringInterface é a interface exposta pelo core
// Não existe razão nenhuma para usar interface e uma extrutura complexa  para um projeto pequeno como este
// Estou fazendo assim exclusivamente pois é um teste e para mostrar que sei fazer :/
type ProcessarStringInterface interface {
	ProcessarString(texto string) MenorNumero
}
