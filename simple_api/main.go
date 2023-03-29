package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const(
	WEBPORT = ":8080"
)

func homefunc(w http.ResponseWriter, r *http.Request){

	fmt.Fprint(w,"We have Recived Request")

}

func main() {
	fmt.Println("My Application is starting")

	router := mux.NewRouter()

	http.Handle("/",router)


	router.HandleFunc(/*this url use in postman*/"/home",homefunc)
	http.ListenAndServe(WEBPORT,nil)
}