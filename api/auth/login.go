package auth

import (
	"encoding/json"
	"github.com/Palantir/palantir/model/auth"
	"io/ioutil"
	"log"
	"net/http"
)

type mockLoginHandler struct{}

func (h *mockLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here is a string"))
}

type loginHandler struct{}

func (h *loginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	account := &auth.Account{}
	requestBody, streamErr := ioutil.ReadAll(r.Body)
	if streamErr != nil {
		log.Fatal(streamErr.Error())
	}
	if parseErr := json.Unmarshal(requestBody, account); parseErr != nil {
		log.Fatal(parseErr.Error())
	}
	account.Id = 1
	account.Password = ""

	response := make(map[string]string)
	response["token"] = "grojieorign043909jg03g04iern3w4f34"
	responseBody, responseErr := json.Marshal(response)
	if responseErr != nil {
		log.Fatal(responseErr.Error())
	}

	w.Write(responseBody)

}