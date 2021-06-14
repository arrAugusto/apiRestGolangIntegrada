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
	Msg    string `json: msg`
	Status string `json: status`
}

/**
*	CONTROLADOR EN EL CUAL SE CASTEA LOS DATOS Y SE RESPONDE AL QUE CONSUME EL API REST
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
	//SETEANDO USUARIO Y CONTRASEÑA DEL USUARIO
	usuarioCast := msg.UsuarioReq
	passwordCast := msg.PasswordReq
	//TOMANDO VARIABLE DEL USUARIO DE SISTEMA
	//TOMANDO CONTRASEÑA DEL USUARIO

	usuarioData, err := Consult.MdlConsultaUsuarios(usuarioCast, passwordCast)
	//MANEJO DE ERROR
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
		return
	}
	var message UserNotFound
	//	CrearJWTUser(usuarioData)
	if len(usuarioData) == 1 {
		message.Msg = "conectado"
		JWTResponse := Auth.CrearJWTUser(usuarioData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(JWTResponse[0])
		return
	} else {

		message.Msg = "Usuario o Contraseña no existe"
		message.Status = "NoAutorizado"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(message)
		return
	}
}
