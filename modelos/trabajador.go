package modelos

import "gorm.io/gorm"

type Trabajador struct {
	gorm.Model
	IdTrabajador int    `gorm:"primary_key;unique_index;AUTO_INCREMENT;column:id" json:"idTrabajador"`
	Nombre       string `json:"nombre"`
	Apellido     string `json:"apellido"`
	Rut          string `json:"rut" gorm:"unique"`
	Correo       string `json:"correo"`
	IdAfp        int    `json:"idAfp"`
	AFP          AFP    `gorm:"foreignkey:id_afp" json:"AFP"`
}

type AFP struct {
	gorm.Model
	IdAfp     int     `gorm:"primary_key;unique_index;AUTO_INCREMENT;column:id_afp" json:"idAfp"`
	Nombre    string  `json:"nombre"`
	Descuento float32 `json:"descuento"`
}
