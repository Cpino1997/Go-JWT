package controladores

import (
	"app/config"
	"app/modelos"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Registrar Usuarios
func RegisterUser(context *gin.Context) {
	var user modelos.User
	// Recibimos la req en formato json
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// Encriptamos la password
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// Traemos la instancia de la bd para actualizar el usuario.
	record := config.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	// Si todo salio bien devolvemos el mensaje con los datos del nuevo usuario creado!
	context.JSON(http.StatusCreated, gin.H{"id": user.ID, "correo": user.Correo, "nombre de usuario": user.Usuario})
}

// Obtener Todos los usuarios
func GetUsuarios(context *gin.Context) {
	// Creamos la Instancia de la bd
	db, err := gorm.Open(config.Instance.Config)
	// Si algo falla lanzamos un error
	if err != nil {
		panic(err)
	}
	// Creamos una lista de usuarios
	var users []modelos.User
	// Buscamos en la bd si existen usuarios para agregarlos a la lista de lo contrario lanzamos un error
	if err := db.Find(&users).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error al listar los usuarios": err.Error()})
		return
	}
	// Devolvemos los usuarios en la res
	context.JSON(http.StatusOK, gin.H{"data": users})
}

// Actualizar usuarios
func UpdateUser(c *gin.Context) {
	// Obtener el usuario
	var user modelos.User
	if err := config.Instance.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario no encontrado!"})
		return
	}
	// Validar ingreso
	var input modelos.UpdateUsuario
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Encriptacion de la password
	if err := input.UpdateHashPassword(input.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	config.Instance.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"Nuevos Datos": input})
}

// Borrar Usuarios
func BorrarUser(c *gin.Context) {
	// Declaramos el usuario
	var user modelos.User
	// Llamamos a la instancia de la bd para consultar si el usuario existe sino lanzamos un error :c
	if err := config.Instance.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Error": "usuario no encontrado!"})
		return
	}
	// Eliminamos el usuario y devolvemos un mensaje al usuario!
	config.Instance.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"Mensaje": "Usuario Eliminado con exito!"})
}
