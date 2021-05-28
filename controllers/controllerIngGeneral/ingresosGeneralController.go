package controllerIngGeneral

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"

	Auth "../../authentication"
	Respuesta "../../models/modelIngGeneral"
	StructDB "../../structures/structuresIngGeneral"
)

type ListaIngGeneralGO struct {
	ListaIngGeneral []string
}

/**
	* CtrIngGeneral Sirve para guardar el ingreso en la db
**/

func CtrIngGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.IngresoGeneralDat
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en la recepción de su objeto JSON "+err.Error(), 500)
		return
	}

	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en la validación de la data "+err.Error(), 500)
		return
	}

	//Si la validación es falsa se retorna el error y termina la ejecución deL codigo
	if RevisionStruct == false {
		http.Error(w, "Error en la validación de la data "+err.Error(), 500)
		return
	}
	//getter de los valores de la estructura
	IdBod := ObjDataIngGeneral.IdBod
	IdNit := ObjDataIngGeneral.IdNit
	CantBlts := ObjDataIngGeneral.CantBlts
	ValTotal := ObjDataIngGeneral.ValTotal
	fechaRealIng := ObjDataIngGeneral.FechaIngreso
	tokenString := ObjDataIngGeneral.TokenReq

	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlNuevoIngresoGeneral(IdUserInt, IdBod, IdNit, CantBlts, ValTotal, fechaRealIng)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

/**
	* CtrIngGeneral remover de la vista el ingreso dar de debaja
**/

func CtrIngGeneralRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)
	idIng := params["idIng"]
	IdIngGeneral, err := strconv.Atoi(idIng)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaIngGeneral"
	respuestaDB := Respuesta.MdlRemoveEstadosActivos(IdIngGeneral, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrMetrajeBodegaGeneralRemove remover de la vista el ingreso dar de debaja
**/

func CtrMetrajeBodegaGeneralRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idMetraje, err := strconv.Atoi(params["idMetraje"])
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaMetraje"
	respuestaDB := Respuesta.MdlRemoveEstadosActivos(idMetraje, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrUbicacionesMercaGRemove remover la ubicacion de mercaderia
**/

func CtrUbicacionesMercaGRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idUbica, err := strconv.Atoi(params["idUbica"])
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaUbicacion"
	respuestaDB := Respuesta.MdlRemoveEstadosActivos(idUbica, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrDowloadSoportLegalRemove remover la visualización de la certificación
**/

func CtrDowloadSoportLegalRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idDoc, err := strconv.Atoi(params["idDoc"])
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaDocDescarga"
	respuestaDB := Respuesta.MdlRemoveEstadosActivos(idDoc, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIngGeneral remover de la vista el ingreso dar de debaja a los detalles de mercaderia
**/

func CtrDetallesGenRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionDetallesGeneral
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URLL
	params := mux.Vars(r)
	idIng := params["idIng"]
	IdIngGeneral, err := strconv.Atoi(idIng)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	idDetalle := ObjDataIngGeneral.IdDetalle
	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaProducto"
	respuestaDB := Respuesta.MdlRemoveEstdDetGeneral(IdIngGeneral, idDetalle, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrRemoveIMGContainer cancelar visualmente el estado de la fotografia del contenedor
**/

func CtrRemoveIMGContainer(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idImg, err := strconv.Atoi(params["idImg"])
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaIMGCont"
	respuestaDB := Respuesta.MdlRemoveEstadosActivos(idImg, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

/**
	* CtrIMGDescMercarRemove cancelar visualmente el estado de la fotografia de las mercaderias recibidas
**/

func CtrIMGDescMercarRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idImg, err := strconv.Atoi(params["idImg"])
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaIMGMeca"
	respuestaDB := Respuesta.MdlRemoveEstadosActivos(idImg, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando nuevos detalles de certificaciones en bodega general

func CtrNewProductGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.NewProductGeneral
	err = json.Unmarshal(b, &ObjProduct)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjProduct)
	if err != nil {
		http.Error(w, "Error en la validación del formulario "+err.Error(), 500)
		return
	}
	if RevisionStruct {

	}
	IdIng := ObjProduct.IdIng
	IdProduct := ObjProduct.IdProduct
	Bultos := ObjProduct.Bultos
	ValorUnitario := ObjProduct.ValorUnitario
	tokenString := ObjProduct.TokenReq
	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	respuestaDB := Respuesta.MdlNewProductGeneral(IdIng, IdUserInt, IdProduct, Bultos, ValorUnitario)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando nuevos productos en bodega general

func ProductosBodGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PR	OVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.NewProducto
	err = json.Unmarshal(b, &ObjProduct)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjProduct)
	if err != nil {
		http.Error(w, "Error en la validación del formulario "+err.Error(), 500)
		return
	}
	if RevisionStruct {

	}
	//GUARDANDO INGRESO
	IdProduct := ObjProduct.Producto
	tokenString := ObjProduct.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	respuestaDB := Respuesta.MdlNewProducto(IdProduct, IdUserInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando incidencia de lo recibido en bodega

func CtrIncidenciaDescGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PR	OVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.IncidenciaDescarga
	err = json.Unmarshal(b, &ObjProduct)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjProduct)
	if err != nil {
		http.Error(w, "Error en la validación del formulario "+err.Error(), 500)
		return
	}
	if RevisionStruct {

	}
	//GUARDANDO INGRESO
	IdIng := ObjProduct.IdIngReq
	IdDetalle := ObjProduct.IdDetalleReq
	Descripcion := ObjProduct.DescripcionReq
	tokenString := ObjProduct.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlIncidenciaDesGenerla(IdIng, IdDetalle, IdUserInt, Descripcion)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando el metodo para guardar metros del lo recibido en bodega

func CtrMetrajeBodegaGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PR	OVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.NewMetraje
	err = json.Unmarshal(b, &ObjProduct)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjProduct)
	if err != nil {
		http.Error(w, "Error en la validación del formulario "+err.Error(), 500)
		return
	}
	if RevisionStruct {

	}
	//GUARDANDO INGRESO
	IdIng := ObjProduct.IdIngReq
	IdDetalle := ObjProduct.IdDetalleReq
	IdAreaBod := ObjProduct.IdAreaBodReq
	Metros := ObjProduct.MetrosReq
	Posiciones := ObjProduct.PosicionesReq
	PromedioTarima := ObjProduct.PromedioTarimaReq
	MetrosStock := ObjProduct.MetrosStockReq
	PosicionesStock := ObjProduct.PosicionesStockReq

	tokenString := ObjProduct.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlNewMetrajeBodGeneral(IdIng, IdDetalle, IdAreaBod, IdUserInt, Metros, Posiciones, PromedioTarima, MetrosStock, PosicionesStock)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando el metodo para guardar metros del lo recibido en bodega

func CtrUbicacionesMercaGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PR	OVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.NewUbicaciones
	err = json.Unmarshal(b, &ObjProduct)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjProduct)
	if err != nil {
		http.Error(w, "Error en la validación del formulario "+err.Error(), 500)
		return
	}
	if RevisionStruct {

	}
	//GUARDANDO INGRESO
	IdIng := ObjProduct.IdIngReq
	IdDetalle := ObjProduct.IdDetalleReq
	IdAreaBod := ObjProduct.IdAreaBodReq
	Pasillo := ObjProduct.Pasillo
	Columna := ObjProduct.Columna
	Comentarios := ObjProduct.Comentarios

	tokenString := ObjProduct.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlNewUbicacionBodegaGeneral(IdIng, IdDetalle, IdAreaBod, IdUserInt, Pasillo, Columna, Comentarios)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

/**
	* CtrUploadSoportLegal Sirve para guardar la imagen o pdf del soporte legal de la descarga
**/

func CtrUploadSoportLegal(w http.ResponseWriter, r *http.Request) {
	//SETEANDO LA DATA EN EL STRUCT

	r.ParseMultipartForm(2000)
	IngIMG := r.FormValue("idIng")
	fmt.Println(IngIMG)
	TokenReq := r.FormValue("TokenReq")
	IdIngGeneral, err := strconv.Atoi(IngIMG)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, fileInfo, err := r.FormFile("fileSoport")
	split := strings.Split(fileInfo.Filename, ".")
	fmt.Println(file)
	if split[1] == "PDF" || split[1] == "pdf" || split[1] == "PNG" || split[1] == "png" || split[1] == "JPG" || split[1] == "jpg" {
		fmt.Println("Es es pdf")
	} else {
		fmt.Println("No es pdf")
		return
	}
	aleatorio := rand.Intn(9999)
	numAleatorioString := strconv.Itoa(aleatorio)
	dt := time.Now()
	nanos := dt.UnixNano()
	timeNano := strconv.Itoa(int(nanos))

	// save Picture
	// save Picture
	os.Mkdir("./docSopport/", 0777)
	os.Mkdir("./docSopport/docDescargaLegal/", 0777)

	f, err := os.OpenFile("./docSopport/docDescargaLegal/"+"goCertificacion"+IngIMG+numAleatorioString+timeNano+"."+split[1], os.O_WRONLY|os.O_CREATE, 0666)
	Ruta := "docSopport/docDescargaLegal/" + "goCertificacion" + IngIMG + numAleatorioString + timeNano + "." + split[1]
	io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
		return
	}
	tokenString := TokenReq
	//Validando token
	ValidaJWT := Auth.TokenValid(tokenString)
	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}
	//Loading token
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	spExecute := "spNewDocDescarga"
	resp := Respuesta.MdlNewDocSistema(IdIngGeneral, IdUserInt, spExecute, Ruta)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(resp)
	return
}

func CtrDowloadSoportLegal(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjIding StructDB.ObjIdIng
	err = json.Unmarshal(b, &ObjIding)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	IdIngG := ObjIding.IdIngReq
	respuestaDB := Respuesta.MdlRutaImagen(IdIngG)

	if err != nil {
		log.Fatal(err)
		return
	}
	rutaImagen := respuestaDB[0].RutaImg
	file, err := os.Open("." + "/" + rutaImagen)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Header().Set("Content-Disposition", rutaImagen)
	w.Write(buff)

}

/**
	* CtrUploadSoportLegal Sirve para guardar la imagen o pdf del soporte legal de la descarga
**/

func CtrContainerDescargaG(w http.ResponseWriter, r *http.Request) {
	//SETEANDO LA DATA EN EL STRUCT

	r.ParseMultipartForm(2000)
	IngIMG := r.FormValue("idIng")
	tokenString := r.FormValue("tokenReq")
	IdIngGeneral, err := strconv.Atoi(IngIMG)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, fileInfo, err := r.FormFile("fileSoport")
	split := strings.Split(fileInfo.Filename, ".")
	if split[1] == "PDF" || split[1] == "pdf" || split[1] == "PNG" || split[1] == "png" || split[1] == "JPG" || split[1] == "jpg" {
		fmt.Println("Es es pdf")
	} else {
		fmt.Println("No es pdf")
		return
	}
	aleatorio := rand.Intn(9999)
	numAleatorioString := strconv.Itoa(aleatorio)
	dt := time.Now()
	nanos := dt.UnixNano()
	timeNano := strconv.Itoa(int(nanos))
	// save Picture
	// save Picture
	os.Mkdir("./docSopport/", 0777)
	os.Mkdir("./docSopport/containerIMGGeneral/", 0777)

	f, err := os.OpenFile("./docSopport/containerIMGGeneral/"+"goContainerIMGGeneral"+IngIMG+numAleatorioString+timeNano+"."+split[1], os.O_WRONLY|os.O_CREATE, 0666)
	Ruta := "docSopport/containerIMGGeneral/" + "goContainerIMGGeneral" + IngIMG + numAleatorioString + timeNano + "." + split[1]
	io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	spExecute := "spNewIMGContainer"
	resp := Respuesta.MdlNewDocSistema(IdIngGeneral, IdUserInt, spExecute, Ruta)

	fmt.Println(resp)
	//fmt.Println(respuestaDB)
	fmt.Fprintf(w, fileInfo.Filename)
}

/**
	* CtrUploadSoportLegal Sirve para guardar la imagen o pdf del soporte legal de la descarga
**/

func CtrIMGDescMerca(w http.ResponseWriter, r *http.Request) {
	//SETEANDO LA DATA EN EL STRUCT

	r.ParseMultipartForm(2000)
	IngIMG := r.FormValue("idDetalle")
	tokenString := r.FormValue("tokenReq")
	idDetalleGeneral, err := strconv.Atoi(IngIMG)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, fileInfo, err := r.FormFile("fileSoport")
	split := strings.Split(fileInfo.Filename, ".")
	if split[1] == "PDF" || split[1] == "pdf" || split[1] == "PNG" || split[1] == "png" || split[1] == "JPG" || split[1] == "jpg" {
		fmt.Println("Es es pdf")
	} else {
		fmt.Println("No es pdf")
		return
	}
	aleatorio := rand.Intn(9999)
	numAleatorioString := strconv.Itoa(aleatorio)
	dt := time.Now()
	nanos := dt.UnixNano()
	timeNano := strconv.Itoa(int(nanos))
	// save Picture
	// save Picture
	os.Mkdir("./docSopport/", 0777)
	os.Mkdir("./docSopport/IMGMercaDescargada/", 0777)

	f, err := os.OpenFile("./docSopport/IMGMercaDescargada/"+"goIMGDescMerca"+IngIMG+numAleatorioString+timeNano+"."+split[1], os.O_WRONLY|os.O_CREATE, 0666)
	Ruta := "docSopport/IMGMercaDescargada/" + "goIMGDescMerca" + IngIMG + numAleatorioString + timeNano + "." + split[1]
	io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	spExecute := "spNewImgProduct"
	resp := Respuesta.MdlNewDocSistema(idDetalleGeneral, IdUserInt, spExecute, Ruta)

	fmt.Println(resp)
	//fmt.Println(respuestaDB)
	fmt.Fprintf(w, fileInfo.Filename)
}

// Unified error output interface
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

/**
	* CtrIngGeneralUpdate Sirve modificar la data del ingreso
**/

func CtrIngGeneralUpdate(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idIng, err := strconv.Atoi(params["idIng"])

	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.UpdateIng
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	bultos := ObjDataIngGeneral.BultosTotal
	ValorTotal := ObjDataIngGeneral.ValorTotal
	Motivo := ObjDataIngGeneral.Motivo

	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlUpdateIngresoGeneral(idIng, IdUserInt, bultos, ValorTotal, Motivo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrDetallesGenUpdate Sirve modificar la data de los detalles de ingresos, el detalle sera modificado si solo si
	* el ingreso esta en estado antes de finalizado, de lo contrario solo el supervisor puede modificar
**/

func CtrDetallesGenUpdate(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idDetalle, err := strconv.Atoi(params["idDetalle"])

	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.UpdateDetalle
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	IdIngReq := ObjDataIngGeneral.IdIngReq
	IdProduct := ObjDataIngGeneral.IdProduct
	PUnitario := ObjDataIngGeneral.PUnitario
	Motivo := ObjDataIngGeneral.Motivo
	Bultos := ObjDataIngGeneral.Bultos
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlUpdateDetallesGeneral(idDetalle, IdIngReq, IdUserInt, IdProduct, Bultos, PUnitario, Motivo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrDetallesGenUpdate Sirve modificar la data de los detalles de ingresos, el detalle sera modificado si solo si
	* el ingreso esta en estado antes de finalizado, de lo contrario solo el supervisor puede modificar
**/

func CtrIncDescGeneralRemove(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idIncidencia, err := strconv.Atoi(params["idIncidencia"])

	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.AnulacionFormas
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	Motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	respuestaDB := Respuesta.MdlDetallesGeneralRemove(idIncidencia, IdUserInt, Motivo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIncDescGeneralUpdate Sirve la incendencia de la mercaderia recibida
**/

func CtrIncDescGeneralUpdate(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idIncidencia, err := strconv.Atoi(params["idIncidencia"])

	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.UpdateIncidencia
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	Motivo := ObjDataIngGeneral.Motivo

	Descripcion := ObjDataIngGeneral.Descripcion

	tokenString := ObjDataIngGeneral.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlIncDescGeneralUpdate(idIncidencia, IdUserInt, Descripcion, Motivo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIncDescGeneralUpdate Sirve la incendencia de la mercaderia recibida
**/

func CtrMetrajeBodegaGeneralUpdate(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idMetraje, err := strconv.Atoi(params["idMetraje"])

	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.UpdateMetraje
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	IdAreaBodReq := ObjDataIngGeneral.IdAreaBodReq
	Metros := ObjDataIngGeneral.Metros
	Posiciones := ObjDataIngGeneral.Posiciones
	Promedio := ObjDataIngGeneral.Promedio
	Motivo := ObjDataIngGeneral.Motivo

	tokenString := ObjDataIngGeneral.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlMetrajeBodegaGeneralUpdate(idMetraje, IdAreaBodReq, IdUserInt, Metros, Posiciones, Promedio, Motivo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIncDescGeneralUpdate Sirve la incendencia de la mercaderia recibida
**/

func CtrUbicacionesMercaGUpdate(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN URL
	params := mux.Vars(r)

	idUbica, err := strconv.Atoi(params["idUbica"])

	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.UpdateUbicacion
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	IdAreaBodReq := ObjDataIngGeneral.IdAreaBodReq
	Pasillo := ObjDataIngGeneral.Pasillo
	Columna := ObjDataIngGeneral.Columna
	Comentario := ObjDataIngGeneral.Comentario
	Motivo := ObjDataIngGeneral.Motivo

	tokenString := ObjDataIngGeneral.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlUbicacionesMercaGUpdate(idUbica, IdAreaBodReq, IdUserInt, Pasillo, Columna, Comentario, Motivo)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

//Creando nuevos productos en bodega general

func ConsultaProductos(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	params := mux.Vars(r)
	Producto := params["producto"]
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	tokenString := r.FormValue("TokenReq")
	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}

	respuestaDB := Respuesta.MdlConsultaProducto(Producto)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando nuevos productos en bodega general

func ConsultaProductosAll(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE

	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	tokenString := r.FormValue("TokenReq")
	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}

	respuestaDB := Respuesta.MdlConsultaProductoAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

//Creando nuevos productos en bodega general

func CtrDetallesGenAll(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE

	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	params := mux.Vars(r)

	tokenString := params["TokenReq"]

	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {	
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}

	respuestaDB := Respuesta.MdlDetallesGenAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}
