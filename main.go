package main

import (
	"github.com/kuramocheez/web-go/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/notes", controller.GetData).Method("GET")
	r.HandleFunc("/notes", controller.PostData).Method("POST")
	r.HandleFunc("/notes/{id}", controller.UpdateData).Method("PUT")
	r.HandleFunc("/notes/{id}", controller.DeleteData).Method("DELETE")
	http.ListenAndServe(":8000", r)
}


