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
	router.HandleFunc("/ingresosGeneral/{idIng}", IngresosGene.CtrIngGeneralUpdate).Methods("PUT")
	router.HandleFunc("/ingresosGeneral/{idIng}", IngresosGene.CtrIngGeneralRemove).Methods("DELETE")

	router.HandleFunc("/detalleDeProductos", IngresosGene.CtrNewProductGeneral).Methods("POST")
	router.HandleFunc("/detalleDeProductos/{idIng}", IngresosGene.CtrDetallesGenRemove).Methods("DELETE")

	router.HandleFunc("/productosBodGeneral", IngresosGene.ProductosBodGeneral).Methods("POST")
	router.HandleFunc("/incidenciaDesGeneral", IngresosGene.CtrIncidenciaDescGeneral).Methods("POST")
	router.HandleFunc("/metrajeMercaGeneral", IngresosGene.CtrMetrajeBodegaGeneral).Methods("POST")
	router.HandleFunc("/ubicacionBodega", IngresosGene.CtrUbicacionesMercaGeneral).Methods("POST")

	/**
		*	IMAGENES UPLOAD AND DOWNLOAD METODOS PARA GESTOR DE CONTENIDO
	**/

	//SUBIENDO IMAGEN O PDF DE LA CERTIFICACION O POLIZA, SOPORTE LEGAL DE LA DESCARGA.
	//GESTOR DE DOCUMENTOS

	//IMAGENES Y PDF DE CERTIFICACIONES CONTABLES
	router.HandleFunc("/ingresosGeneralUpload", IngresosGene.CtrUploadSoportLegal).Methods("POST")
	router.HandleFunc("/ingresosGeneralUpload/{idImg}", IngresosGene.CtrUploadSoportLegal).Methods("DELETE")

	router.HandleFunc("/docSopport", IngresosGene.CtrDowloadSoportLegal)
	//IMAGENES DE CONTENEDOR DE DESCARGA
	router.HandleFunc("/containerDescargaGeneral", IngresosGene.CtrContainerDescargaG).Methods("POST")
	//IMAGENES DE MERCADERIAS RECIBIDAS
	router.HandleFunc("/mercaderiasRecibidas", IngresosGene.CtrIMGDescMerca).Methods("POST")

	/**
		*	ABRIENDO EL CANAL Y PUERTO DE COMUNICACIÓN
	**/
	http.ListenAndServe(":3001", handlers.CORS(headers, origins, methods)(router))
}
