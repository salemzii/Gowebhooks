package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		router := gin.Default()

		router.GET("/", welcome)
		router.POST("valid/response", ReceiveCustomerIsValidHook)

		router.Run()
	*/
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
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

func ReceiveCustomerIsValidHook(c *gin.Context) {
	var response PayHookResponse
	if c.Request.Method == "POST" {
		c.BindJSON(&response)
		log.Println(response)
	}
}

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}
