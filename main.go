package main

import (
	"fmt"
	"net/http"

	ConsUSer "./models/modelUser"
)

func main() {
	//ConsUSer.ConsultaUsuarios()
	fmt.Println("hola mundo mi server")
	http.HandleFunc("/usuarios", ConsUSer.ConsultaUsuarios)
	http.ListenAndServe(":8070", nil)
}
