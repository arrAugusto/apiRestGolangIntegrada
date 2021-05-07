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
	Respuesta "../../models/modelRetGeneral"
	StructDB "../../structures/structuresRetGeneral"
)

type ListaIngGeneralGO struct {
	ListaIngGeneral []string
}

/**
	* CtrIngGeneral Sirve para guardar el ingreso en la db
**/

func CtrNewDetRetGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.DetalleRetGen
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

	//GUARDANDO INGRESO
	fmt.Println(RevisionStruct)
	IdRet := ObjProduct.IdRet
	IdDetalle := ObjProduct.IdDetalle
	TotalBultos := ObjProduct.TotalBultos
	tokenString := ObjProduct.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlNewDetRetGeneral(IdRet, IdDetalle, IdUserInt, TotalBultos)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

func CtrNewRetiroGeneral(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//SETEANDO LA DATA EN EL STRUCT
	var ObjProduct StructDB.NewRetiroGeneral
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

	//GUARDANDO INGRESO
	fmt.Println(RevisionStruct)
	TotalBultos := ObjProduct.TotalBultos
	tokenString := ObjProduct.TokenReq
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))
	respuestaDB := Respuesta.MdlNewRetiroGeneral(TotalBultos, IdUserInt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

func CtrGuardarIMGGeneral(w http.ResponseWriter, r *http.Request) {
	//SETEANDO LA DATA EN EL STRUCT

	r.ParseMultipartForm(2000)
	IdRet := r.FormValue("idRet")
	tokenString := r.FormValue("tokenReq")
	idRetGeneral, err := strconv.Atoi(IdRet)
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
	os.Mkdir("./docSopport/IMGRetGeneral/", 0777)

	f, err := os.OpenFile("./docSopport/IMGRetGeneral/"+"goIMGRetGeneral"+IdRet+numAleatorioString+timeNano+"."+split[1], os.O_WRONLY|os.O_CREATE, 0666)
	Ruta := "docSopport/IMGRetGeneral/" + "goIMGRetGeneral" + IdRet + numAleatorioString + timeNano + "." + split[1]
	io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	spExecute := "spNewIMGRet"
	resp := Respuesta.MdlMGRetiroGeneral(idRetGeneral, IdUserInt, spExecute, Ruta)

	fmt.Println(resp)
	//fmt.Println(respuestaDB)
	fmt.Fprintf(w, fileInfo.Filename)

}

///debaja de retiros
/**
	* CtrIngGeneral remover de la vista el ingreso dar de debaja
**/

func CtrRetiroGeneralRemove(w http.ResponseWriter, r *http.Request) {
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
	idRet := params["idRet"]
	IdRetGeneral, err := strconv.Atoi(idRet)
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
	respuestaDB := Respuesta.MdlAnulaTransaccionesRet(IdRetGeneral, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIngGeneral remover de la vista el ingreso dar de debaja
**/

func CtrDetalleGeneralRemove(w http.ResponseWriter, r *http.Request) {
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
	idDetalle := params["idDetalle"]
	idDetalleAnl, err := strconv.Atoi(idDetalle)
	fmt.Println(idDetalleAnl)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaRebajaGeneral"
	respuestaDB := Respuesta.MdlAnulaTransaccionesRet(idDetalleAnl, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIngGeneral remover de la vista el ingreso dar de debaja
**/

func CtrIMGRetGeneralRemove(w http.ResponseWriter, r *http.Request) {
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
	idIMGRet := params["idIMGRet"]
	idIMGRetAnl, err := strconv.Atoi(idIMGRet)
	if err != nil {
		http.Error(w, "Error en el envio de datos", 500)
		return
	}

	motivo := ObjDataIngGeneral.Motivo
	tokenString := ObjDataIngGeneral.TokenReq

	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdUserJWT"]
	var IdUserInt int = int(IdUser.(float64))

	storeProduce := "spDebajaIMGRet"
	respuestaDB := Respuesta.MdlAnulaTransaccionesRet(idIMGRetAnl, IdUserInt, motivo, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return
}

/**
	* CtrIncDescGeneralUpdate Sirve la incendencia de la mercaderia recibida
**/
/*
func CtrDetalleGeneralUpdate(	) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//LEYENDO LA VARIABLE EN UR
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
*/
