package routes

import (
	"Usuarios_CRUD/controllers"
	"github.com/gorilla/mux"
)

//Registro de ruutas para la tabla usuario

func RegisterHistorial_accesoRouetes (r*mux.Router){
	r.HandleFunc("/access", controllers.GetALLHistorial_acceso).Methods("GET")
	r.HandleFunc("/access/{id}", controllers.GetHistorial_accesoByID).Methods("GET")
	r.HandleFunc("/access", controllers.CreateHistorial_acceso).Methods("POST")
	r.HandleFunc("/access/{id}", controllers.UpdateHistorial_acceso).Methods("PUT")
	r.HandleFunc("/access/{id}", controllers.DeleteHistorial_acceso).Methods("DELETE")
}