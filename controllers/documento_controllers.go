package controllers

import (
	"Usuarios_CRUD/models"
	"Usuarios_CRUD/config"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" 
)

func GetALLDocuments(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_documento, tipo_documento, activo, fecha_creacion, fecha_modificacion FROM documento  WHERE 1=1"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}

	defer rows.Close()

	var Documento []models.Documento

	for rows.Next() {
		var d models.Documento

		rows.Scan(&d.ID_documento, &d.Tipo_documento,&d.Activo,&d.Fecha_creacion, &d.Fecha_modificacion)
		Documento = append(Documento, d)

	}

	respondJSON(w, 200, Documento)
	
}

//GET PARA Documentsbyid

func GetDocumentsByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var d models.Documento

	err := config.DB.QueryRow(
		"SELECT id_documento, tipo_documento, activo, fecha_creacion, fecha_modificacion  FROM documento WHERE id_documento = $1",
		id,
	).Scan(&d.ID_documento, &d.Tipo_documento,&d.Activo,&d.Fecha_creacion, &d.Fecha_modificacion)

	if err != nil{
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, d)
}

// POST create nuevos datos

func CreateDocuments(w http.ResponseWriter, r *http.Request) {
	var c models.Documento

	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		"INSERT INTO documento (tipo_documento ) VALUES ($1) RETURNING id_documento",
		c.Tipo_documento,
	).Scan(&c.ID_documento)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 201, map[string]string{"Message": "Dato Creado"})
}

// UpdateUser actualizar 
func UpdateDocumnts(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var d models.Documento
	json.NewDecoder(r.Body).Decode(&d)

	_, err := config.DB.Exec(
		"UPDATE documento SET  tipo_documento= $1, activo = $2 WHERE id_documento = $3",
		d.Tipo_documento, d.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato Actualizado"})
}
// DeleteUser eliminar 
func DeleteDocuments(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	
	_, err := config.DB.Exec("DELETE FROM documento WHERE id_documento = $1", id)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato eliminado"})
}
