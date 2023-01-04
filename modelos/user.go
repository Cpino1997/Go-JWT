package modelos

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// La estructura de nuestro usuario
type User struct {
	gorm.Model
	Nombre   string `json:"nombre"`
	Usuario  string `json:"usuario" gorm:"unique"`
	Correo   string `json:"correo" gorm:"unique"`
	Password string `json:"password"`
}

// Estructura de la actualizacion del usuario
type UpdateUsuario struct {
	gorm.Model
	Nombre   string `json:"nombre" binding:"required"`
	Usuario  string `json:"usuario" binding:"required"`
	Correo   string `json:"correo" gorm:"unique"`
	Password string `json:"password"`
}

// Funcion que recibe un usuario y le encripta la constraseña
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// Funcion para Checkear la password Encriptada
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

//Funcion para Encriptar la contraseña actualizada del usuario
func (u *UpdateUsuario) UpdateHashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

// Funcion para Checkear la encriptacion de la nueva contraseña
func (u *UpdateUsuario) UpdateCheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
