package controladores

import (
	"app/config"
	"app/modelos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CrearTrabajadores(context *gin.Context) {
	var temp struct {
		IdTrabajador string `json:"idTrabajador"`
		Correo       string `json:"correo"`
		Nombre       string `json:"nombre"`
		IdAfp        string `json:"idAfp"`
		AFP          struct {
			IdAfp     string `json:"idAfp"`
			Nombre    string `json:"nombre"`
			Descuento string `json:"descuento"`
		} `json:"AFP"`
	}

	if err := context.ShouldBindJSON(&temp); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idTrabajador, err := strconv.Atoi(temp.IdTrabajador)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	idAfp, err := strconv.Atoi(temp.IdAfp)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	trabajador := modelos.Trabajador{
		IdTrabajador: idTrabajador,
		Correo:       temp.Correo,
		Nombre:       temp.Nombre,
		IdAfp:        idAfp,
	}
	/*
		var newAfp modelos.AFP
		newAfp.IdAfp = 1
		newAfp.Descuento = 11.5
		newAfp.Nombre = "Plan Vital"

		record := config.Instance.Create(&newAfp)
		if record.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
			context.Abort()
			return
		} */
	record := config.Instance.Create(&trabajador)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"id": trabajador.IdTrabajador, "correo": trabajador.Correo, "nombre": trabajador.Nombre, "AFP": trabajador.IdAfp})
}
func GetTrabajadores(context *gin.Context) {
	var trabajadores []modelos.Trabajador
	db, err := gorm.Open(config.Instance.Config)
	if err != nil {
		panic(err)
	}
	// Buscamos en la bd si existen usuarios para agregarlos a la lista de lo contrario lanzamos un error
	if err := db.Find(&trabajadores).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error al listar los usuarios": err.Error()})
		return
	}
	// Devolvemos los usuarios en la res
	context.JSON(http.StatusOK, gin.H{"data": trabajadores})
}
