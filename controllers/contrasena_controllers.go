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

//GET PARA contrasenabyid

func GetContrasenaByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var c models.contrasena

	err := config.DB.QueryRow(
		"SELECT id_contrasena, contrasena, hash_contrasena, id_usuario, activo FROM contrasena WHERE id = $1",
		id,
	).Scan(&c.ID_contrasena, &c.Contrasena,&c.Hash_contrasena, &c.ID_usuario, &c.Activo)

	if err == sql.ErrNoRows{
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, c)
}

// POST create nuevos datos

func CreateContrasena(w http.ResponseWriter, r *http.Request) {
	var c models.contrasena

	json.NewDecoder(r.Body).Decode(&c)

	err := config.DB.QueryRow(
		"INSERT INTO contrasena(id_usuario,contrasena, hash_contrasena) VALUES ($1, $2, $3) RETURNING id",
		c.ID_usuario,c.Contrasena, c.Hash_contrasena,
	).Scan(&c.ID)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 201, map[string]string{"Message": "Dato Creado"})
}

// UpdateUser actualizar 
func UpdateContrasena(w http.ResponseWriter, r *http.Request){
		id:=mux.Vars(r)["id"]

		var c models.contrasena
		json.NewDecoder(r.Body).Decode(&c)

		_, err := config.DB.Exec(
			"UPDATE Contrasena SET contrasena id=$2",
			c.Password, id,
		)

	if err!=nil {
		respondJSON(w,500,map[string]string{"Error":err.Error()})
		return
	}
	respondJSON(w,200,map[string]string{"Message":"Dato actualizado"})

}
// DeleteUser eliminar 
func DeleteContrasena(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	
	_, err := config.DB.Exec("DELETE FROM contrasena WHERE id = $1", id)

	if err != nil {
		respondJSON(w, 500,map[string]string{"Error":"Error"})
		return
	}
	respondJSON(w, 200, map[string]string{"Message": "Dato eliminado"})
}
