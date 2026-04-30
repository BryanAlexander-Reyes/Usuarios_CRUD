package routes

import (
	"Usuarios_CRUD/controllers"
	"github.com/gorilla/mux"
)

//Registro de ruutas para la tabla usuario

func RegisterContrasenaRoutes(r *mux.Router){
	r.HandleFunc("/password", controllers.GetALLUsuario).Methods("GET")
	r.HandleFunc("/password/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/password", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/password/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/password/{id}", controllers.DeleteUser).Methods("DELETE")
}