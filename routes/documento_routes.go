package routes

import (
	"Usuarios_CRUD/controllers"
	"github.com/gorilla/mux"
)

//Registro de ruutas para la tabla usuario

func RegisterDocumentoRoutes(r *mux.Router){
	r.HandleFunc("/document", controllers.GetALLDocuments).Methods("GET")
	r.HandleFunc("/document/{id}", controllers.GetDocumentsByID).Methods("GET")
	r.HandleFunc("/document", controllers.CreateDocuments).Methods("POST")
	r.HandleFunc("/document/{id}", controllers.UpdateDocumnts).Methods("PUT")
	r.HandleFunc("/document/{id}", controllers.DeleteDocuments).Methods("DELETE")
}