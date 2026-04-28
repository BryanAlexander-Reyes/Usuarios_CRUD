package models

import ( 
	"Time"
)

type Historial_acceso struct{
	Id_historial_acceso int `json:"id_historial_acceso"` 
	Fecha_intento time.Time `json:"fecha_intento"`
	Exitoso bool `json:"Exitoso"`
	Ip_origen string `json:"ip_origen"`
	Fallo_motivo string `json:"fallo_motivo"`
	Id_usuario int `json:"id_usuario"`
	Id_contrasena int `json:"id_contrasena"`
	Activo bool `json:"activo"`
	Fecha_creacion time.Time `json:"fecha_creacion"`
	Fecha_modificacion time.Time `json:"fecha_modificacion"`
}