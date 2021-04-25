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
	tokenString := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZFVzZXJKV1QiOjEwMzcsIk5vbWJyZUpXVCI6IlJPTlkgVklOSUNJTyIsIkFwZWxsaWRvc0pXVCI6IkFSUklPTEEgTMOTUEVaIiwiRmVjaGFfY3JlYWNpb25KV1QiOiIyMDIwLTEwLTAyVDAwOjAwOjAwWiIsIkNlbHVsYXJKV1QiOiIzMzQ3MjExNCIsIkVtYWlsSldUIjoiY2RvbmlzQGJpLmNvbS5ndCIsIk5pdmVsZXNKV1QiOiJCQUpPIiwiRGVwZW5kZW5jaWFKV1QiOjEsIklkRGVCb2RlZ2FKV1QiOjEsIkZvdG9KV1QiOiIiLCJFc3RhZG9KV1QiOjEsIkRlcGFydGFtZW50b3NKV1QiOiJCb2RlZ2FzIEZpc2NhbGVzIiwiZXhwIjoxNjE5MTgxMDA4LCJpc3MiOiJJbmNpbyBkZSBzZXNpw7NuIn0.jG-b8PiKG0S8B5GxM4BOHbD5CfGv0OZ02JivHEjAqsmHjJiUGBlg8UbaXOq1K4XkWMzNuZp_jGiNQtxHd47bD1qyQQBDAD81HxsdBwZn7enMAehf31posZ7hMWYR8NzDc1mNDUZ2M8p17pUUIt9vSLwxLOormhutQmeCLRJVNpY"
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	fmt.Println(JWTResponse["IdUserJWT"])
	IdUser := JWTResponse["IdUserJWT"]

	var IdUserInt int = int(IdUser.(float64))
	fmt.Println(IdUserInt)
	respuestaDB := Respuesta.MdlNewProductGeneral(IdIng, IdUserInt, IdProduct, Bultos, ValorUnitario)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

/**
	* CtrUploadSoportLegal Sirve para guardar la imagen o pdf del soporte legal de la descarga
**/

func CtrUploadSoportLegal(w http.ResponseWriter, r *http.Request) {
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

	f, err := os.OpenFile("./docSopport/docDescargaLegal/"+"go"+IngIMG+numAleatorioString+timeNano+"."+split[1], os.O_WRONLY|os.O_CREATE, 0666)
	Ruta := "docSopport/docDescargaLegal/" + "go" + IngIMG + numAleatorioString + timeNano + "." + split[1]
	io.Copy(f, file)
	if err != nil {
		log.Fatal(err)
		return
	}

	resp := Respuesta.MdlNewDocSistema(Ruta, IdIngGeneral)

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

// Unified error output interface
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
