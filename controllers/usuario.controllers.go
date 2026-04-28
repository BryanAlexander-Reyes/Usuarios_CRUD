package controllers

import (
	"API_GO_CRUD/models"
	"API_GO_CRUD/config"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux" 
)

func GetALLUsuario(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_usuario, nombre, apellido, numero_documento, email, telefono, fecha_nacimiento,id_documento, fecha_registro, activo FROM usuario WHERE 1=1"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}

	defer rows.Close()

	var Usuario []models.Usuario

	for rows.Next() {
		var u models.usuario

		rows.Scan(&u.ID_usuario, &u.Nombre, &u.Apellido, &u.Numero_documento, &u.Email, &u.Telefono, &u.Fecha_nacimiento, &u.ID_documento, &u.Fecha_registro, &u.Activo)
		Usuario = append(Usuario, u)

	}

	respondJSON(w, 200, Usuario)
	
}

//GET PARA userbyid

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var u models.usuario

	err := config.DB.QueryRow(
		"SELECT id_usuario, nombre, apellido, numero_documento, email, telefono, fecha_nacimiento,id_documento, fecha_registro, activo FROM usuario WHERE id = $1",
		id,
	).Scan(&u.ID_usuario, &u.Nombre, &u.Apellido, &u.Numero_documento, &u.Email, &u.Telefono, &u.Fecha_nacimiento, &u.ID_documento, &u.Fecha_registro, &u.Activo)

	if err == sql.ErrNoRows{
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, u)
}

