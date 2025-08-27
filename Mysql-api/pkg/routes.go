package rts

import(
	"github.com/gorilla/mux"
	"db_mysql/ctrls"
)

var RoutesFunc = func (r*mux.router){
	r.HandleFunc("/class",ctrls.GetClasses).Methods("GET")
	r.HandleFunc("/class/{ID}",ctrls.GetClass).Methods("GET")
	r.HandleFunc("/class",ctrls.CreateClass).Methods("POST")
	r.HandleFunc("/class/{ID}",ctrls.UpdateClass).Methods("PUT")
	r.HandleFunc("/class/{ID}",ctrls.DeleteClass).Methods("DELETE")
}