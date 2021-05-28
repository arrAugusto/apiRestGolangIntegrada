package main

import (
	"net/http"

	IngresosGene "./controllers/controllerIngGeneral"
	NitEmpresas "./controllers/controllerNit"
	RetirosGene "./controllers/controllerRetGeneral"

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
	/*

	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   									***** AUTOR AUGUSTO GOMEZ
	   									***** MODULO DE INGRESOS GENERALES Y SUS ENDPOINT
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	*/

	//INGRESO DE MERCADERIA  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/ingresosGeneral", IngresosGene.CtrIngGeneral).Methods("POST")
	router.HandleFunc("/ingresosGeneral/{idIng}", IngresosGene.CtrIngGeneralUpdate).Methods("PUT")
	router.HandleFunc("/ingresosGeneral/{idIng}", IngresosGene.CtrIngGeneralRemove).Methods("DELETE")

	//INGRESO DE DETALLES  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/detalleDeProductos", IngresosGene.CtrNewProductGeneral).Methods("POST")
	router.HandleFunc("/detalleDeProductos/{idDetalle}", IngresosGene.CtrDetallesGenUpdate).Methods("PUT")
	router.HandleFunc("/detalleDeProductos/{idIng}", IngresosGene.CtrDetallesGenRemove).Methods("DELETE")
	router.HandleFunc("/detalleDeProductosAll/{TokenReq}", IngresosGene.CtrDetallesGenAll).Methods("GET")

	//INGRESO DE INCIDENCIA  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/incidenciaDesGeneral", IngresosGene.CtrIncidenciaDescGeneral).Methods("POST")
	router.HandleFunc("/incidenciaDesGeneral/{idIncidencia}", IngresosGene.CtrIncDescGeneralUpdate).Methods("PUT")
	router.HandleFunc("/incidenciaDesGeneral/{idIncidencia}", IngresosGene.CtrIncDescGeneralRemove).Methods("DELETE")
	//INGRESO DE METRAJE MERCADERIA  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/metrajeMercaGeneral", IngresosGene.CtrMetrajeBodegaGeneral).Methods("POST")
	router.HandleFunc("/metrajeMercaGeneral/{idMetraje}", IngresosGene.CtrMetrajeBodegaGeneralUpdate).Methods("PUT")
	router.HandleFunc("/metrajeMercaGeneral/{idMetraje}", IngresosGene.CtrMetrajeBodegaGeneralRemove).Methods("DELETE")
	//INGRESO DE UBICACION BODEGA  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/ubicacionBodega", IngresosGene.CtrUbicacionesMercaGeneral).Methods("POST")
	router.HandleFunc("/ubicacionBodega/{idUbica}", IngresosGene.CtrUbicacionesMercaGUpdate).Methods("PUT")
	router.HandleFunc("/ubicacionBodega/{idUbica}", IngresosGene.CtrUbicacionesMercaGRemove).Methods("DELETE")

	router.HandleFunc("/productosBodGeneral", IngresosGene.ProductosBodGeneral).Methods("POST")
	router.HandleFunc("/productosBodGeneral/{producto}", IngresosGene.ConsultaProductos).Methods("GET")
	router.HandleFunc("/productosBodGeneralAll", IngresosGene.ConsultaProductosAll).Methods("GET")

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

	/*

	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   									***** AUTOR AUGUSTO GOMEZ
	   									***** MODULO DE INGRESOS GENERALES Y SUS ENDPOINT
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	*/

	//INGRESO DE MERCADERIA  **** POST=INGRESO **** PUT UPDATE **** DELETE ANULACIÓN
	router.HandleFunc("/detalleRebajaRetGen", RetirosGene.CtrNewDetRetGeneral).Methods("POST")
	router.HandleFunc("/detalleRebajaRetGen/{idDetalle}", RetirosGene.CtrDetalleGeneralRemove).Methods("PUT")
	router.HandleFunc("/detalleRebajaRetGen/{idDetalle}", RetirosGene.CtrDetalleGeneralRemove).Methods("DELETE")

	router.HandleFunc("/retirosGeneral", RetirosGene.CtrNewRetiroGeneral).Methods("POST")
	router.HandleFunc("/retirosGeneral/{idRet}", RetirosGene.CtrRetiroGeneralRemove).Methods("DELETE")

	router.HandleFunc("/IMGRetMercaGeneral", RetirosGene.CtrGuardarIMGGeneral).Methods("POST")
	router.HandleFunc("/IMGRetMercaGeneral/{idIMGRet}", RetirosGene.CtrIMGRetGeneralRemove).Methods("DELETE")

	/*

	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   									***** AUTOR AUGUSTO GOMEZ
	   									***** CONSULTAS REQUERIDAS POR GET
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	   *****
	*/
	router.HandleFunc("/nitEmpresa/{numNit}", NitEmpresas.CtrMostrarNit).Methods("GET")
	router.HandleFunc("/bodegasInfo/{TokenReq}", NitEmpresas.CtrBodegasInfo).Methods("GET")

	/**
		*	ABRIENDO EL CANAL Y PUERTO DE COMUNICACIÓN
	**/

	http.ListenAndServe(":3001", handlers.CORS(headers, origins, methods)(router))
}
