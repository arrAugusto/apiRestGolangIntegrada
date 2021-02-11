package models

/**
*El uso de la libreria github.com/dgrijalva/jwt-go es para uso de manejo de tokens y generacion del mismo
*
**/
import jwt "github.com/dgrijalva/jwt-go"

//mi structura claim
type Claim struct {
	User `json: "user"`
	jwt.StandardClaims
}
