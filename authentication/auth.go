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
	IdUserJWT         int    `json: idJWT`
	NombreJWT         string `json: nombreJWT`
	ApellidosJWT      string `json: apellidosJWT`
	Fecha_creacionJWT string `json: fechaCreacionJWT`
	CelularJWT        string `json: telefonoJWT`
	EmailJWT          string `json: emailJWT`
	NivelesJWT        string `json: nivelesJWT`
	DependenciaJWT    int    `json: dependenciaJWT`
	IdDeBodegaJWT     int    `json: idBodegaJWT`
	FotoJWT           string `json: fotoJWT`
	EstadoJWT         int    `json: estadoJWT`
	DepartamentosJWT  string `json: departamentosJWT`
}
type Claim struct {
	JWTAuth `json: myUser`
	packetJWT.StandardClaims
}
type JWTRespStarting struct {
	TokenStarting string `json: tokenStarting`
}

var JwtResp []JWTAuth
var tokenStart []JWTRespStarting

/**
*
CREANDO JWT SEGUN EL LA STRUCTURA DEL JWTAUTH
**/
func CrearJWTUser(data []StructUser.Usuario) []JWTRespStarting {
	var jwt JWTAuth
	var jwtStart JWTRespStarting
	jwt.IdUserJWT = data[0].IdUser

	jwt.NombreJWT = data[0].Nombre
	jwt.ApellidosJWT = data[0].Apellidos
	jwt.Fecha_creacionJWT = data[0].Fecha_creacion
	jwt.CelularJWT = data[0].Celular
	jwt.EmailJWT = data[0].Email
	jwt.NivelesJWT = data[0].Niveles
	jwt.DependenciaJWT = data[0].Dependencia
	jwt.IdDeBodegaJWT = data[0].IdDeBodega
	jwt.FotoJWT = data[0].Foto
	jwt.EstadoJWT = data[0].Estado
	jwt.DepartamentosJWT = data[0].Departamentos

	JwtResp = append(JwtResp, jwt)
	jwtGenerado := GenerateJWT(JwtResp, []Claim{})

	jwtStart.TokenStarting = jwtGenerado
	TokenStarting := append(tokenStart, jwtStart)

	return TokenStarting

}

func GenerateJWT(data []JWTAuth, claim []Claim) string {

	claims := Claim{
		JWTAuth: data[0],
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
