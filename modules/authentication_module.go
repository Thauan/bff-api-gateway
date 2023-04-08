package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Thauan/bff-api-gateway/handlers"
	user_credentials "github.com/Thauan/bff-api-gateway/shared/protobuf"
	"google.golang.org/protobuf/proto"
)

type UserCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	User UserCredentials `json:"user"`
}

func SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t UserRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Fatalln(err.Error())
		}

		authApiUrl := handlers.GetEnvWithKey("AUTHENTICATION_API_URL")

		userCred := &user_credentials.UserCredentials{
			User: &user_credentials.UserCredentials_User{
				Email:    t.User.Email,
				Password: t.User.Password,
			},
		}

		fmt.Println(userCred)

		// Codificar o objeto protobuf para enviar como corpo da solicitação
		body, err2 := proto.Marshal(userCred)

		fmt.Println(body)

		if err2 != nil {
			log.Fatal("Error marshaling request:", err2.Error())
		}

		payload := bytes.NewBuffer(body)

		// Enviar a solicitação HTTP POST
		res, err3 := http.Post(authApiUrl+"/sign_in", "application/octet-stream", payload)
		if err3 != nil {
			log.Fatalln("Erro ao ler resposta: " + err3.Error())
		}

		defer res.Body.Close()

		// Ler a resposta e decodificar em um objeto protobuf
		resBody := &user_credentials.AuthResponse{}
		data, err4 := ioutil.ReadAll(res.Body)

		if err4 != nil {
			log.Fatalln(err4.Error())
		}

		fmt.Println(data)
		err5 := proto.Unmarshal(data, resBody)

		if err5 != nil {
			log.Fatalln(err5.Error())
		}

		jsonBytes, err6 := json.Marshal(resBody)

		if err6 != nil {
			log.Fatalln(err5.Error())
		}

		statusCode := res.StatusCode

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(jsonBytes)
	}
}

func SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var t UserRequest
		err := decoder.Decode(&t)
		if err != nil {
			panic(err)
		}

		authApiUrl := handlers.GetEnvWithKey("AUTHENTICATION_API_URL")

		userReq := UserRequest{
			User: UserCredentials{
				Email:    t.User.Email,
				Password: t.User.Password,
			},
		}

		body, err := json.Marshal(userReq)

		if err != nil {
			log.Fatalln(err.Error())
		}

		payload := bytes.NewBuffer(body)

		res, err2 := http.Post(authApiUrl+"/sign_up", "application/json", payload)

		if err != nil {
			log.Fatalln(err2)
		}

		defer res.Body.Close()

		var resBody interface{}
		err3 := json.NewDecoder(res.Body).Decode(&resBody)
		if err3 != nil {
			log.Fatalln(err)
		}

		data, err4 := json.Marshal(resBody)

		if err4 != nil {
			log.Fatalln(err)
		}

		statusCode := res.StatusCode

		w.WriteHeader(statusCode)
		w.Write(data)
	}
}
