package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/bff/v1/sign_in", func(w http.ResponseWriter, r *http.Request) {
		reqBody, _ := ioutil.ReadAll(r.Body)

		var user User
		json.Unmarshal(reqBody, &user)
		json.NewEncoder(w).Encode(user)

		newData, err := json.Marshal(user)

		println(newData)
		println(err)

		// Chama a camada de serviço e retorna uma resposta adequada.
		// myData := User{Email: "teste", Password: "teste"}
		json.NewEncoder(w).Encode(newData)
	})

	log.Fatal(http.ListenAndServe(":9090", r))
}
