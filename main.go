package main

import (
	"log"
	"net/http"

	ConsUSer "./controllers/controllerUser"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/usuarios", ConsUSer.CtrConsultUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":3001", router))
}
