package controllers

import (
	"API_GO_CRUD/models"
	"API_GO_CRUD/config"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" 
)

func GetALLDocuments(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_documento, tipo_documento, activo, fecha _creacion FROM documento  WHERE 1=1"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}

	defer rows.Close()

	var Documento []models.documento

	for rows.Next() {
		var d models.documento

		rows.Scan(&d.ID_documento &d.Tipo_documento,&d.Activo,&d.Fecha_creacion)
		Documento = append(Documento, d)

	}

	respondJSON(w, 200, Documento)
	
}

