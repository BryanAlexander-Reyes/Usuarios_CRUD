package controllers

import (
	"API_GO_CRUD/models"
	"API_GO_CRUD/config"
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

	var Historial_acceso []models.historial_acceso

	for rows.Next() {
		var h models.historial_acceso

		rows.Scan(&h.ID_historial_acceso, &h.Fecha_intento,&h.Exitoso,&h.Ip_origen,&h.Fallo_motivo, &h.ID_usuario,&h.Id_contrasena &h.Activo)
		Historial_acceso = append(Historial_acceso, h)

	}

	respondJSON(w, 200, Historial_acceso)
	
}

//GET PARA contrasenabyid

func GetHistorial_accesoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var h models.historial_acceso

	err := config.DB.QueryRow(
		"SELECT id_historial_acceso, fecha_intento, exitoso,ip_fallo, fallo_motivo, id_usuario, id_contraseña, activo FROM historial_acceso WHERE id = $1",
		id,
	).Scan(&h.ID_historial_acceso, &h.Fecha_intento,&h.Exitoso,&h.Ip_origen,&h.Fallo_motivo, &h.ID_usuario,&h.Id_contrasena &h.Activo)

	if err == sql.ErrNoRows{
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, h)
}

