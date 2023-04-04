package modules

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Thauan/bff-api-gateway/handlers"
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

		fmt.Println(payload)

		res, err2 := http.Post(authApiUrl+"/sign_in", "application/json", payload)

		if err2 != nil {
			log.Fatalln(err2)
		}

		defer res.Body.Close()

		var resBody interface{}
		err3 := json.NewDecoder(res.Body).Decode(&resBody)

		if err3 != nil {
			log.Fatalln(err3)
		}

		data, err4 := json.Marshal(resBody)

		if err4 != nil {
			log.Fatalln(err)
		}

		fmt.Println(data)

		w.WriteHeader(http.StatusCreated)
		w.Write(data)
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

		fmt.Println(payload)

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

		fmt.Println(data)

		w.WriteHeader(http.StatusCreated)
		w.Write(data)
	}
}
