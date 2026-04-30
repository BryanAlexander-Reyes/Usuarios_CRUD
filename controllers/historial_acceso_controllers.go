package controllers

import (
	"Usuarios_CRUD/models"
	"Usuarios_CRUD/config"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" 
)

func GetALLHistorial_acceso(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_historial_acceso, fecha_intento, exitoso,ip_fallo, fallo_motivo, id_usuario, id_contraseña, activo FROM historial_acceso  WHERE 1=1"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}

	defer rows.Close()

	var Historial_acceso []models.Historial_acceso

	for rows.Next() {
		var h models.Historial_acceso

		rows.Scan(&h.Id_historial_acceso, &h.Fecha_intento,&h.Exitoso,&h.Ip_origen,&h.Fallo_motivo, &h.Id_usuario,&h.Id_contrasena, &h.Activo)
		Historial_acceso = append(Historial_acceso, h)

	}

	respondJSON(w, 200, Historial_acceso)
	
}

//GET PARA contrasenabyid

func GetHistorial_accesoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var h models.Historial_acceso

	err := config.DB.QueryRow(
		"SELECT id_historial_acceso, fecha_intento, exitoso,ip_fallo, fallo_motivo, id_usuario, id_contraseña, activo FROM historial_acceso WHERE id = $1",
		id,
	).Scan(&h.Id_historial_acceso, &h.Fecha_intento,&h.Exitoso,&h.Ip_origen,&h.Fallo_motivo, &h.Id_usuario, &h.Id_contrasena, &h.Activo)

	if err == sql.ErrNoRows{
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, h)
}

// POST create nuevos datos

func CreateHistorial_acceso(w http.ResponseWriter, r *http.Request) {
	var h models.Historial_acceso

	json.NewDecoder(r.Body).Decode(&h)

	err := config.DB.QueryRow(
		"INSERT INTO historial_acceso( fecha_intento, exitoso,ip_fallo, fallo_motivo, id_usuario, id_contraseña VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		h.Fecha_intento, h.Exitoso, h.Ip_origen, h.Fallo_motivo, h.Id_usuario, h.Id_contrasena,
	).Scan(&h.Id_historial_acceso)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 201, map[string]string{"Message": "Dato Creado"})
}

// UpdateUser actualizar 
func UpdateHistorial_acceso(w http.ResponseWriter, r *http.Request){
		id:=mux.Vars(r)["id"]

		var h models.Historial_acceso
		json.NewDecoder(r.Body).Decode(&h)

		_, err := config.DB.Exec(
			"UPDATE historial_acceso SET fecha_intento=$1,existoso=$2, ip_origen=$3, fallo_motivo=$4 id=$5",
			h.Fecha_intento, h.Exitoso, h.Ip_origen, h.Fallo_motivo, h.Id_usuario, h.Id_contrasena, id,
		)

	if err!=nil {
		respondJSON(w,500,map[string]string{"Error":err.Error()})
		return
	}
	respondJSON(w,200,map[string]string{"Message":"Dato actualizado"})

}
// DeleteUser eliminar 
func DeleteHistorial_acceso(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	
	_, err := config.DB.Exec("DELETE FROM historial_acceso WHERE id = $1", id)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato eliminado"})
}
