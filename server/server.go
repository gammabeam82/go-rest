package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func Run(router *mux.Router, port int) {
	log.Printf("Listening at localhost:%v\n", port)

	handler := cors.Default().Handler(router)

	log.Fatal(
		http.ListenAndServe(fmt.Sprintf("localhost:%v", port), handler),
	)
}
