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
	router.HandleFunc("/ingNewProductGeneral", IngresosGene.CtrNewProductGeneral).Methods("POST")

	//SUBIENDO IMAGEN O PDF DE LA CERTIFICACION O POLIZA, SOPORTE LEGAL DE LA DESCARGA.
	//GESTOR DE DOCUMENTOS
	router.HandleFunc("/ingresosGeneralUpload", IngresosGene.CtrUploadSoportLegal).Methods("POST")
	router.HandleFunc("/docSopport", IngresosGene.CtrDowloadSoportLegal)

	http.ListenAndServe(":3001", handlers.CORS(headers, origins, methods)(router))
}
