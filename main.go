package main

import (
	"app/config"
	"app/controladores"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializamos la DB
	config.Connect("root:root@tcp(localhost:3306)/app?parseTime=true")
	config.Migrate()
	// Initialize Router
	router := initRouter()

	err := router.SetTrustedProxies([]string{"192.168.1.2"})
	if err != nil {
		return
	}
	// Arrancamos el servidor en el puerto 8080
	err = router.Run(":8080")
	if err != nil {
		return
	}
}

// Inicializamos el router y lo aseguramos con nuestro JWT
func initRouter() *gin.Engine {
	router := gin.Default()
	// El contexto en el que trabajara, Todas las solicitudes deben realizarse por HOST:PORT/API
	router.GET("/", controladores.Inicio)
	api := router.Group("/api")
	{
		api.GET("/", controladores.Inicio)
		// Endpoint para logearse y generar un token
		api.POST("/token", controladores.GenerarToken)
		// Endpoint para Registrar un usuario
		api.POST("/user/registro", controladores.RegisterUser)
		secured := api.Group("").Use(middlewares.Auth())
		{
			// Prueba de seguridad.
			secured.GET("/ping", controladores.Ping)
			// Get All Users
			secured.GET("/users", controladores.GetUsuarios)
			// Actualizar un usuario
			secured.PUT("/users/:id", controladores.UpdateUser)
			// Eliminar un usuario
			secured.DELETE("/users/:id", controladores.BorrarUser)
			//Get trabajadores
			secured.GET("/trabajadores", controladores.GetTrabajadores)
			secured.POST("/trabajadores", controladores.CrearTrabajadores)
		}
	}
	return router
}
