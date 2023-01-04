package controladores

import (
	"app/config"
	"app/modelos"
	JwtUtils "app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenRequest struct {
	Correo   string `json:"correo"`
	Password string `json:"password"`
}

func GenerarToken(context *gin.Context) {
	var request TokenRequest
	var user modelos.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	record := config.Instance.Where("correo = ?", request.Correo).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales Invalidas =("})
		context.Abort()
		return
	}

	tokenString, err := JwtUtils.GenerarToken(user.Correo, user.Usuario)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
