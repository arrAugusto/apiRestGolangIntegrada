package authentication

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	models "../models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	//llave privada
	privateKey *rsa.PrivateKey
	//lave public
	publicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("No se puede leer el archivo privado")
	}
	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("No se puede leer el archivo publico")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("no se cargo el archiv privado")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("no se cargo el archiv public")
	}
}

func GenerateJWT(user models.User) string {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			Audience:  "",
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "JWT Integrada",
			NotBefore: 0,
			Subject:   "",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)

	if err != nil {
		log.Fatal("No se pudo firmar token")
	}
	return result
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintln(w, "error al leer el usuario")
		return
	}
	if user.Name == "Augusto" && user.Password == "Gomez" {
		user.Password = ""
		user.Role = "admin"

		token := GenerateJWT(user)
		result := models.ResponseToken{token}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			fmt.Fprintln(w, "error al convertir el json")
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "aplication/json")
		w.Write(jsonResult)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "usuario no valido")
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil { /*
				switch.err(type){
			case *jwt.ValidationError:
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					fmt.Fprintln(w, "su token expiro")
					return

				case jwt.ValidateErrorSignatureInvalid:
					fmt.Fprintln(w, "la firma esta corrupta")
				default:
				}
				default:
				fmt.Fprintln(w, "la firma esta corrupta")

				}*/
		fmt.Fprintln(w, "su token no es valido")
		return

	}
	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Bienvenido al sistema")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Su token no es valido")

	}

}
