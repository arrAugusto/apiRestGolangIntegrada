package main

import (
	"fmt"
	"net/http"

	ConsUSer "./controllers/controllerUser"
)

func main() {
	//ConsUSer.ConsultaUsuarios()
	fmt.Println("hola mundo mi server")
	http.HandleFunc("/usuarios", ConsUSer.CtrConsultUser)
	http.ListenAndServe(":8070", nil)
}
