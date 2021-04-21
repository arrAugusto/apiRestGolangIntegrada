package controllerIngGeneral

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"

	Respuesta "../../models/modelIngGeneral"
	StructDB "../../structures/structuresIngGeneral"
)

type ListaIngGeneralGO struct {
	ListaIngGeneral []string
}

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
