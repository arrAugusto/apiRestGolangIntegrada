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

	//INGRESO DE MERCADERIA  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/ingresosGeneral", IngresosGene.CtrIngGeneral).Methods("POST")
	router.HandleFunc("/ingresosGeneral/{idIng}", IngresosGene.CtrIngGeneralUpdate).Methods("PUT")
	router.HandleFunc("/ingresosGeneral/{idIng}", IngresosGene.CtrIngGeneralRemove).Methods("DELETE")

	//INGRESO DE DETALLES  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/detalleDeProductos", IngresosGene.CtrNewProductGeneral).Methods("POST")
	router.HandleFunc("/detalleDeProductos/{idDetalle}", IngresosGene.CtrDetallesGenUpdate).Methods("PUT")
	router.HandleFunc("/detalleDeProductos/{idIng}", IngresosGene.CtrDetallesGenRemove).Methods("DELETE")

	router.HandleFunc("/incidenciaDesGeneral", IngresosGene.CtrIncidenciaDescGeneral).Methods("POST")
	router.HandleFunc("/incidenciaDesGeneral/{idIncidencia}", IngresosGene.CtrIncDescGeneralUpdate).Methods("PUT")
	router.HandleFunc("/incidenciaDesGeneral/{idIncidencia}", IngresosGene.CtrIncDescGeneralRemove).Methods("DELETE")

	router.HandleFunc("/metrajeMercaGeneral", IngresosGene.CtrMetrajeBodegaGeneral).Methods("POST")
	router.HandleFunc("/metrajeMercaGeneral/{idMetraje}", IngresosGene.CtrMetrajeBodegaGeneralUpdate).Methods("PUT")
	router.HandleFunc("/metrajeMercaGeneral/{idMetraje}", IngresosGene.CtrMetrajeBodegaGeneralRemove).Methods("DELETE")

	router.HandleFunc("/ubicacionBodega", IngresosGene.CtrUbicacionesMercaGeneral).Methods("POST")
	router.HandleFunc("/ubicacionBodega/{idUbica}", IngresosGene.CtrUbicacionesMercaGUpdate).Methods("PUT")
	router.HandleFunc("/ubicacionBodega/{idUbica}", IngresosGene.CtrUbicacionesMercaGRemove).Methods("DELETE")

	router.HandleFunc("/productosBodGeneral", IngresosGene.ProductosBodGeneral).Methods("POST")

	/**
		*	IMAGENES UPLOAD AND DOWNLOAD METODOS PARA GESTOR DE CONTENIDO
	**/

	//SUBIENDO IMAGEN O PDF DE LA CERTIFICACION O POLIZA, SOPORTE LEGAL DE LA DESCARGA.
	//GESTOR DE DOCUMENTOS

	//IMAGENES Y PDF DE CERTIFICACIONES CONTABLES
	router.HandleFunc("/ingresosGeneralUpload", IngresosGene.CtrUploadSoportLegal).Methods("POST")
	router.HandleFunc("/docSopport/{idDoc}", IngresosGene.CtrDowloadSoportLegalRemove).Methods("DELETE")
	router.HandleFunc("/docSopport", IngresosGene.CtrDowloadSoportLegal).Methods("POST")

	//IMAGENES DE CONTENEDOR DE DESCARGA
	router.HandleFunc("/containerDescargaGeneral", IngresosGene.CtrContainerDescargaG).Methods("POST")
	router.HandleFunc("/containerDescargaGeneral/{idImg}", IngresosGene.CtrRemoveIMGContainer).Methods("DELETE")

	//IMAGENES DE MERCADERIAS RECIBIDAS
	router.HandleFunc("/mercaderiasRecibidas", IngresosGene.CtrIMGDescMerca).Methods("POST")
	router.HandleFunc("/mercaderiasRecibidas/{idImg}", IngresosGene.CtrIMGDescMercarRemove).Methods("DELETE")

	/**
		*	ABRIENDO EL CANAL Y PUERTO DE COMUNICACIÓN
	**/
	http.ListenAndServe(":3001", handlers.CORS(headers, origins, methods)(router))
}
