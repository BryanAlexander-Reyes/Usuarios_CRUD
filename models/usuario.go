package models

import ( 
	"time"
)

type Usuario struct{
	ID_usuario int `json:"id_usuario"` // campo id 
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Numero_documento int `json:"numero_documento"`
	Email string `json:"email"`
	Telefono int `json:"telefono"`
	Fecha_nacimiento time.Time `json:"fecha_nacimiento"`
	ID_documento int `json:"id_documento"`
	Fecha_registro time.Time `json:"fecha_registro"`
	Activo bool `json:"activo"`
	Fecha_creacion time.Time `json:"fecha_creacion"`
	Fecha_modificacion time.Time `json:"fecha_modificacion"`
}