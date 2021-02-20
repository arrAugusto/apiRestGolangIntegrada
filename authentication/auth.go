package authentication

import (
	"fmt"

	StructUser "../structures/structuresUser"
)

type TodoApi struct{}

func NewTodoApi() *TodoApi {
	return &StructUser.Usuario{}
	return &StructUser{}

}

type JWTAuth struct {
	IdUserJWT int `json: id`
}

var JwtResp []JWTAuth

/**
*
CREANDO JWT SEGUN EL LA STRUCTURA DEL JWTAUTH
**/
func CrearJWTUser(data []StructUser.Usuario) (ta TodoApi) {

	var jwt JWTAuth
	jwt.IdUserJWT = data[0].IdUser
	JwtResp = append(JwtResp, jwt)
	fmt.Printf(data[0].Apellidos)
	var err error
	return JWTAuth, err

}
