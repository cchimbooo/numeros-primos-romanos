package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"prova-programacao/pkg/core/menorPrimoRomano"
)

type SearchRequest struct {
	Text string `json:"text"`
}
type SearchResponse struct {
	Decimal int    `json:"value"`
	Romano  string `json:"number"`
}

// Search cria a rota HTTP
func Search(romanoInterface menorPrimoRomano.ProcessarStringInterface) func(c *gin.Context) {
	return func(c *gin.Context) {
		var request SearchRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		dados := romanoInterface.ProcessarString(request.Text)

		// Não sei como vocês preferem retornar, temos  a mania de retornar os campos vazios em vez de um null no response.
		// por isso esta retornando {"" e 0 se não encontrarmos nenhum número}. É intencional.
		response := SearchResponse{
			Decimal: dados.Decimal,
			Romano:  dados.Romano,
		}
		c.JSON(200, response)
	}
}
