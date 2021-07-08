package controllerNit

import (
	"encoding/json"
	"fmt"
	"net/http"

	Auth "../../authentication"
	Respuesta "../../models/modelNit"
	"github.com/gorilla/mux"
)

/**
	* CtrMostrarNit muestra nit y datos generales del cliente
**/

func CtrMostrarNit(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	params := mux.Vars(r)
	NumeroNit := params["numNit"]
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	tokenString := r.FormValue("TokenReq")
	fmt.Println(tokenString)
	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}

	respuestaDB := Respuesta.MdlConsultaNit(NumeroNit)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}

/**
	* CtrMostrarNit muestra nit y datos generales del cliente
**/

func CtrBodegasInfo(w http.ResponseWriter, r *http.Request) {
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	params := mux.Vars(r)
	tokenString := params["TokenReq"]	
	// LEYENDO LA DATA PROVENIENTE DEL CLIENTE
	ValidaJWT := Auth.TokenValid(tokenString)

	if ValidaJWT != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error en la validación del token " + ValidaJWT.Error())
		return
	}
	JWTResponse := Auth.ReadPyloadJWT(tokenString)
	IdUser := JWTResponse["IdDeBodegaJWT"]
	var IdDeBodegaJWT int = int(IdUser.(float64))

	respuestaDB := Respuesta.MdlBodegasInfo(IdDeBodegaJWT)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}
