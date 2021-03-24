package controllerUser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	Auth "../../authentication"
	Consult "../../models/modelUser"
	packetJWT "github.com/dgrijalva/jwt-go"
)

/**
*	ESTRUCTURA `UsuarioReq` DE DATOS SE UTILIZA PARA TOMAR LOS VALORES DEL SP EL CUAL FORMARA PARTE DEL PAYLOAD DE JWT
**/

type DataRequerida struct {
	UsuarioReq  int    `json: usuarioReq`
	PasswordReq string `json: passwordReq`
	packetJWT.StandardClaims
}

type UserNotFound struct {
	Msg string `json: msg`
}

/**
*	CONTROLADORA EN EL CUAL SE CASTEA LOS DATOS Y SE RESPONDE AL QUE CONSUME EL API REST
**/
func CtrConsultUser(w http.ResponseWriter, r *http.Request) {
	// Read body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var msg DataRequerida
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	usuarioCast := msg.UsuarioReq
	passwordCast := msg.PasswordReq
	//TOMANDO VARIABLE DEL USUARIO DE SISTEMA

	//TOMANDO CONTRASEÑA DEL USUARIO
	usuarioData, err := Consult.MdlConsultaUsuarios(usuarioCast, passwordCast)

	JWTResponse := Auth.CrearJWTUser(usuarioData)

	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
		return
	}

	//	CrearJWTUser(usuarioData)
	if len(usuarioData) == 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(JWTResponse)

	} else {

		var message UserNotFound
		message.Msg = "Usuario o Contraseña no existe"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(message)
		return
	}
}
