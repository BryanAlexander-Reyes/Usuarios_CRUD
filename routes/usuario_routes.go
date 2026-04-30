package routes

import (
	"Usuarios_CRUD/controllers"
	"github.com/gorilla/mux"
)

//Registro de ruutas para la tabla usuario

func RegisterUserRoutes(r *mux.Router){
	r.HandleFunc("/usuario", controllers.GetALLUsuario).Methods("GET")
	r.HandleFunc("/usuario/{id}", controllers.GetUserByID).Methods("GET")
	r.HandleFunc("/usuario", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/usuario/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/usuario/{id}", controllers.DeleteUser).Methods("DELETE")
}