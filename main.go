package main

import (
	"net/http"

	IngresosGene "./controllers/controllerIngGeneral"
	ConsUSer "./controllers/controllerUser"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	// do all your routes declaration

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	router.HandleFunc("/usuarios", ConsUSer.CtrConsultUser).Methods("POST")

	//INGRESO NUEVOS DE MERCADERIA
	router.HandleFunc("/ingresosGeneral", IngresosGene.CtrIngGeneral).Methods("POST")

	http.ListenAndServe(":3001", handlers.CORS(headers, origins, methods)(router))
}
