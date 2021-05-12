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

	storeProduce := "spNitGeneralIng"
	respuestaDB := Respuesta.MdlConsultaNit(NumeroNit, storeProduce)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(respuestaDB)
	return

}
