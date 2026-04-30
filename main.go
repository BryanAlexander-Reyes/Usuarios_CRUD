package main

import (
	"log"
	"net/http"

	"Usuarios_CRUD/config"
	"Usuarios_CRUD/routes"

	"github.com/gorilla/mux")

// middleware CORS
func enableCORS(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// Permiten cualquier origen de la petición
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == "OPTIONS" {
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func main(){
	config.ConnectDB() // Conexión a la base de datos

	r := mux.NewRouter()

	// registro de rutas 
	routes.RegisterUserRoutes(r)
	routes.RegisterContrasenaRoutes(r)
	routes.RegisterDocumentoRoutes(r)
	routes.RegisterHistorial_accesoRouetes(r)
	

	// Rutas de la api
	
	log.Println("Servidor operando en el puerto 8080")

	log.Fatal(http.ListenAndServe(":8080", enableCORS(r)))
}