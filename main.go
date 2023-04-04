package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Thauan/bff-api-gateway/handlers"
	"github.com/Thauan/bff-api-gateway/modules"

	"github.com/gorilla/mux"
)

func init() {
	handlers.LoadEnv()
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/bff/v1/sign_in", modules.SignIn()).Methods("POST")

	Port := handlers.GetEnvWithKey("PORT")

	fmt.Println("Running in http://localhost:" + Port)

	log.Fatal(http.ListenAndServe(":"+Port, r))
}
