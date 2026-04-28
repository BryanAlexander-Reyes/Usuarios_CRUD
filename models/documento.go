package models

import ( 
	"Time"
)

type Documento struct{
	ID int `json:"id_documento"`
	Tipo_documento string `json:"tipo_documento"`
	Activo bool `json:"activo"`
	Fecha_creacion time.Time `json:"fecha_creacion"`
	Fecha_modificacion time.Time `json:"fecha_modificacion"`
}