package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	http.Handle("/", router)

	router.HandleFunc("validuser/response", ReceiveCustomerIsValidHook)
}

type PayHookResponse struct {
	Event string     `json:"event"`
	Data  DataStruct `json:"data"`
}

type DataStruct struct {
	Customer_id    string         `json:"customer_id"`
	Customer_code  string         `json:"customer_code"`
	Email          string         `json:"email"`
	Identification Identification `json:"Identification"`
}

type Identification struct {
	Country        string `json:"country"`
	Type           string `json:"type"`
	Bvn            string `json:"bvn"`
	Account_number string `json:"account_number"`
	Bank_code      string `json:"bank_code"`
}

func ReceiveCustomerIsValidHook(w http.ResponseWriter, req *http.Request) {
	var response PayHookResponse
	if req.Method == "POST" {

		respByte, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
		}
		e := json.Unmarshal(respByte, &response)
		if e != nil {
			log.Println(e)
		}
		log.Println(response)
	}
}
