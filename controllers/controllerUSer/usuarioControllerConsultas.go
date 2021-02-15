package controllerUser

import (
	"encoding/json"
	"fmt"
	"net/http"

	Consult "../../models/modelUser"
)

func CtrConsultUser(w http.ResponseWriter, r *http.Request) {
	contactos, err := Consult.ConsultaUsuarios()
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contactos)

}
