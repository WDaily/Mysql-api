package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"db_mysql/pkg"
)

func main(){

	r := mux.NewRouter()
	rts.RoutesFunc(r)
	http.Handle("/",r)

	log.Fatal(http.ListenAndServe("localhost:8080",r))
}