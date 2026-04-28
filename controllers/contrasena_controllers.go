package controllers

import (
	"API_GO_CRUD/models"
	"API_GO_CRUD/config"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" 
)

func GetALLContrasena(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_contrasena, contrasena, hash_contrasena, id_usuario, activo FROM contrasena  WHERE 1=1"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}

	defer rows.Close()

	var Contrasena []models.contrasena

	for rows.Next() {
		var c models.contrasena

		rows.Scan(&c.ID_contrasena, &c.Contrasena,&c.Hash_contrasena, &c.ID_usuario, &c.Activo)
		Contrasena = append(Contrasena, d)

	}

	respondJSON(w, 200, Contrasena)
	
}

