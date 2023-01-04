package controladores

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Inicio(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"mensaje": "Servidor Funcionando Sin Problemas!"})
}
