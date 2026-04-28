package models

import ( 
	"Time"
)

type Contrasena struct{
	ID_contrasena int `json:"id_contrasena"` // campo id 
	Contrasena string `json:"contrasena"`
	Hash_contrasena string `json:"hash_contrasena"`
	ID_usuario int `json:"id_usuario"`
	Activo bool `json:"activo"`
	Fecha_creacion time.Time `json:"fecha_creacion"`
	Fecha_modificacion time.Time `json:"fecha_modificacion"`
}