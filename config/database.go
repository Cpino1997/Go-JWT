package config

import (
	"app/modelos"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Declaramos a GORM
var Instance *gorm.DB
var dbError error

// Funcion para conectar la base de datos
func Connect(connectionString string) {
	Instance, dbError = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("No ha sido posible conectarse a la base de datos")
	}
	log.Println("Conectado a la base de datos!")
}
func Migrate() {
	err := Instance.AutoMigrate(&modelos.User{})
	if err != nil {
		return
	}
	log.Println("Actualizacion de la base de datos realizada con exito!")
}
