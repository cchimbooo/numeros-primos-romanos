package rest

import (
	"github.com/gin-gonic/gin"
	"prova-programacao/pkg/core/menorPrimoRomano"
)

func SetUpRoute(romanoInterface menorPrimoRomano.ProcessarStringInterface) error {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.POST("/search", Search(romanoInterface))
	return r.Run(":8080")

}
