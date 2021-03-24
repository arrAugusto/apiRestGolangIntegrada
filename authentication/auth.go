package authentication

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

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
}
type Claim struct {
	MyUser int `json: myUser`
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
	jwtGenerado := GenerateJWT(JwtResp, []Claim{})
	fmt.Print(jwtGenerado)
	return JwtResp

}

func init() {
	//
	privateBytes, err := ioutil.ReadFile("authentication/private.rsa")
	if err != nil {
		log.Fatal("No se pudo leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("authentication/public.rsa.pub")
	if err != nil {
		log.Fatal("No se pudo leer el archivo publico")
	}

	privateKey, err = packetJWT.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("No se puede hacer el parse a privatekey")

	}
	publicKey, err = packetJWT.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("No se puede hacer el parse a publickey")

	}
}

func GenerateJWT(data []JWTAuth, claim []Claim) string {

	claims := Claim{
		MyUser: data[0].IdUserJWT,
		StandardClaims: packetJWT.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "taller",
		},
	}
	token := packetJWT.NewWithClaims(packetJWT.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
		fmt.Printf(result)

	}
	return result
}
