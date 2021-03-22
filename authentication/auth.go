package authentication

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"

	StructUser "../structures/structuresUser"
	packetJWT "github.com/dgrijalva/jwt-go"
)

//guaradando variables de javes publicas y privadas
var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type JWTAuth struct {
	IdUserJWT int `json: id`
	packetJWT.StandardClaims
}

var JwtResp []JWTAuth

/**
*
CREANDO JWT SEGUN EL LA STRUCTURA DEL JWTAUTH
**/
func CrearJWTUser(data []StructUser.Usuario) []JWTAuth {

	var jwt JWTAuth
	jwt.IdUserJWT = data[0].IdUser
	JwtResp = append(JwtResp, jwt)

	return JwtResp

}

func init() {
	//
	privateBytes, err := ioutil.ReadFile("authentication/private.rsa")
	fmt.Println("hola mudno")
	if err != nil {
		fmt.Println(err)
	}
	publicBytes, err := ioutil.ReadFile("authentication/public.rsa.pub")
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err = packetJWT.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		fmt.Println(err)

	}
	publicKey, err = packetJWT.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		fmt.Println(err)

	}
}
