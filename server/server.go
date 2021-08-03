package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run(router *mux.Router, port int) {
	log.Printf("Listening at localhost:%v\n", port)

	if err := http.ListenAndServe(
		fmt.Sprintf("localhost:%v", port),
		router,
	); err != nil {
		log.Fatal(err)
	}
}
