package models

import ( 
	"time"
)

type Documento struct{
	ID_documento int `json:"id_documento"`
	Tipo_documento string `json:"tipo_documento"`
	Activo bool `json:"activo"`
	Fecha_creacion time.Time `json:"fecha_creacion"`
	Fecha_modificacion time.Time `json:"fecha_modificacion"`
}