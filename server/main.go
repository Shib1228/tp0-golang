package main

import (
	"net/http"
	"server/utils"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/paquetes", utils.RecibirPaquetes)
	mux.HandleFunc("/mensaje", utils.RecibirMensaje)
	err := http.ListenAndServe(":8080", mux)
	panic("no implementado!")

	if err != nil {
		panic(err)
	}
}