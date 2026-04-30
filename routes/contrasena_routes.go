package routes

import (
	"Usuarios_CRUD/controllers"
	"github.com/gorilla/mux"
)

//Registro de ruutas para la tabla usuario

func RegisterContrasenaRoutes(r *mux.Router){
	r.HandleFunc("/password", controllers.GetALLContrasena).Methods("GET")
	r.HandleFunc("/password/{id}", controllers.GetContrasenaByID).Methods("GET")
	r.HandleFunc("/password", controllers.CreateContrasena).Methods("POST")
	r.HandleFunc("/password/{id}", controllers.UpdateContrasena).Methods("PUT")
	r.HandleFunc("/password/{id}", controllers.DeleteContrasena).Methods("DELETE")
}