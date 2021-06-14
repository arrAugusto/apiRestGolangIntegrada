package authentication

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	StructUser "../structures/structuresUser"
	"github.com/dgrijalva/jwt-go"
	packetJWT "github.com/dgrijalva/jwt-go"
)

//Jwt Read
type JwtReadRevision struct {
	TokenReq string `json: tokenReq`
}

//guaradando variables de javes publicas y privadas
var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

//Nuevo objeto a partir del recibido en el retorno y validacion del login
type JWTAuth struct {
}

//Objeto que guarda el JWT a retornar
type Claim struct {
	JWTAuth `json: myUser`
	packetJWT.StandardClaims
}

//Array que retorna el objeto JWT ala vista
type JWTRespStarting struct {
	TokenStarting     string `json: tokenStarting`
	Status            string `json: status`
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

//inicializando los metodos a utilizar
var JwtResp []JWTAuth
var tokenStart []JWTRespStarting

/**
*
CREANDO JWT SEGUN EL LA STRUCTURA DEL JWTAUTH
**/
func CrearJWTUser(data []StructUser.Usuario) []JWTRespStarting {
	//haciendo publico el uso del objeto JWTAuth / JWTRespStarting

	var jwt JWTAuth

	var jwtStart JWTRespStarting
	//Seteando y limpiando objetos a retornar
	jwtStart.Status = ""
	jwtStart.TokenStarting = ""

	JwtResp = nil
	//Haciendo set al objeto JWTAuth
	jwtStart.IdUserJWT = data[0].IdUser
	jwtStart.NombreJWT = data[0].Nombre
	jwtStart.ApellidosJWT = data[0].Apellidos
	jwtStart.Fecha_creacionJWT = data[0].Fecha_creacion
	jwtStart.CelularJWT = data[0].Celular
	jwtStart.EmailJWT = data[0].Email
	jwtStart.NivelesJWT = data[0].Niveles
	jwtStart.DependenciaJWT = data[0].Dependencia
	jwtStart.IdDeBodegaJWT = data[0].IdDeBodega
	jwtStart.FotoJWT = data[0].Foto
	jwtStart.EstadoJWT = data[0].Estado
	jwtStart.DepartamentosJWT = data[0].Departamentos

	//Cargando variables al objeto
	JwtResp = append(JwtResp, jwt)
	//Invocando el metodo que crea el JWT
	jwtGenerado := GenerateJWT(JwtResp, []Claim{})

	//Asignando string JWT a el objeto a retornar como array
	jwtStart.TokenStarting = jwtGenerado
	jwtStart.Status = "conectado"

	//Cargando al metodo jwt
	TokenStarting := append(tokenStart, jwtStart)

	//Retornando el string token de conexión
	return TokenStarting

}

/**
*	CREANDO EL JWT
**/
func GenerateJWT(data []JWTAuth, claim []Claim) string {
	//utilizando el metodo claim y crando el jwt
	claims := Claim{
		JWTAuth: data[0],
		StandardClaims: packetJWT.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "Incio de sesión",
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

/**
*	CARGANDO LOS FICHEROS PRIVADOS Y PUBLICOS PARA HACER EL PARSEO DEL JWT
**/
func init() {
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

//LEYENDO EL PAYLOAD DEL JSON WEB TOKEN
func ReadPyloadJWT(tokenString string) jwt.MapClaims {

	token, err := jwt.Parse(tokenString, nil)
	if token == nil {
		fmt.Println("Error validate jwt " + err.Error())
	}
	claims, _ := token.Claims.(jwt.MapClaims)

	return claims

}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil

}

func TokenValid(TokeView string) error {
	token, err := ValidateToken(TokeView)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}
