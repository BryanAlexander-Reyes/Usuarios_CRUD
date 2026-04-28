package controllers

import (
	"API_GO_CRUD/models"
	"API_GO_CRUD/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" 
)

// GetALLUsers obtener todos los usuarios con filtros opcionales
func GetALLUsuario(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id_usuario, nombre, apellido, numero_documento, email, telefono, fecha_nacimiento,id_documento, fecha_registro, activo FROM userio WHERE 1=1"

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}

	defer rows.Close()

	var Usuario []models.User

	for rows.Next() {
		var u models.User

		rows.Scan(&u.ID_usuario, &u.Nombre, &u.Apellido, &u.Numero_documento, &u.Email, &u.Telefono, &u.Fecha_nacimiento, &u.ID_documento, &u.Fecha_registro, &u.Activo)
		Usuario = append(Usuario, u)

	}

	respondJSON(w, 200, Usuario)
	
}
