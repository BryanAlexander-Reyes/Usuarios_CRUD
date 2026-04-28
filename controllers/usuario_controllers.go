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

// POST create nuevos datos

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var c models.usuario

	json.NewDecoder(r.Body).Decode(c)

	err := config.DB.QueryRow(
		"INSERT INTO usuario (nombre, apellido, numero_documento, email, telefono, fecha_necimiento, id_documento fecha_registro, activo) VALUES ($1, $2, $3, $4) RETURNING id",
		c.Nombre, c.Apellido, c.Numero_documento, c.Email, c.Fecha_nacimineto, c.ID_documento, c.Fecha_registro, c.Activo,
	).Scan(&c.ID)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 201, map[string]string{"Message": "Dato Creado"})
}

// UpdateUser actualizar usuario
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var u models.usuario
	json.NewDecoder(r.Body).Decode(&u)

	_, err := config.DB.Exec(
		"UPDATE usuario SET telefono = $1, email = $2, activo = $3 WHERE id = $4",
		u.Telefono, u.Email, u.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato Actualizado"})
}

// DeleteUser eliminar usuario
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	
	_, err := config.DB.Exec("DELETE FROM usuario WHERE id = $1", id)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato eliminado"})
}