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
		http.Error(w, "Error en el envio de datos", 500)
		return
	}
	//VALIDANDO LA DATA QUE CUMPLAN CON LO SOLICITADO
	RevisionStruct, err := govalidator.ValidateStruct(ObjDataIngGeneral)
	if err != nil {
		http.Error(w, "Error en la validación del formulario "+err.Error(), 500)
		return
	}

	//GUARDANDO INGRESO
	fmt.Println(RevisionStruct)
	IdUser := ObjDataIngGeneral.IdUser
	IdBod := ObjDataIngGeneral.IdBod
	IdNit := ObjDataIngGeneral.IdNit
	CantBlts := ObjDataIngGeneral.CantBlts
	ValTotal := ObjDataIngGeneral.ValTotal

	respuestaDB := Respuesta.MdlNuevoIngresoGeneral(IdUser, IdBod, IdNit, CantBlts, ValTotal)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

/**
	* CtrIngGeneral Sirve modificar la data del ingreso
**/

func CtrIngGeneralUpdate(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	params := mux.Vars(r)
	idIng := params["idIng"]
	fmt.Println(idIng)
	//SETEANDO LA DATA EN EL STRUCT
	var ObjDataIngGeneral StructDB.ObjIdIng
	err = json.Unmarshal(b, &ObjDataIngGeneral)
	fmt.Println(ObjDataIngGeneral)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("respuestaDB")
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

	//LEYENDO LA VARIABLE EN UR
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

	//LEYENDO LA VARIABLE EN UR
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

//Creando nuevos productos en bodega general

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
	//GUARDANDO INGRESO

	IdIng := ObjProduct.IdIng
	IdProduct := ObjProduct.IdProduct
	Bultos := ObjProduct.Bultos
	ValorUnitario := ObjProduct.ValorUnitario
	tokenString := ObjProduct.TokenReq
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
	var ObjProduct StructDB.JwtRead

	r.ParseMultipartForm(2000)
	IngIMG := r.FormValue("idIng")
	IdIngGeneral, err := strconv.Atoi(IngIMG)
	if err != nil {
		log.Fatal(err)
		return
	}
	file, fileInfo, err := r.FormFile("fileSoport")
	split := strings.Split(fileInfo.Filename, ".")
	if split[1] == "PDF" || split[1] == "pdf" || split[1] == "PNG" || split[1] == "png" || split[1] == "JPG" || split[1] == "jpg" {
		fmt.Println("Es es pedf")
	} else {
		fmt.Println("No es pedf")
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
	tokenString := ObjProduct.Token
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	spExecute := "spNewDocDescarga"
	resp := Respuesta.MdlNewDocSistema(IdIngGeneral, IdUserInt, spExecute, Ruta)

	fmt.Println(resp)
	//fmt.Println(respuestaDB)
	fmt.Fprintf(w, fileInfo.Filename)
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
	fmt.Println(rutaImagen)
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
		fmt.Println("Es es pedf")
	} else {
		fmt.Println("No es pedf")
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
		fmt.Println("Es es pedf")
	} else {
		fmt.Println("No es pedf")
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
