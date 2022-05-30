package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/test/{query}",controllerfunc)
	http.ListenAndServe("0.0.0.0:8080",router)
}
func controllerfunc(w http.ResponseWriter, r *http.Request){
	if r.Body!=nil{
		defer r.Body.Close()
	}
	vars:=mux.Vars(r)
	q:=vars["query"]
	fmt.Print("qqqqqqqqq",q)
	fmt.Fprintf(w,"controller returns %v",q)
}