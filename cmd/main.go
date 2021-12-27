package main

import(
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main(){
	router := mux.NewRouter()
	
	router.HandleFunc("/", getFeed)

	http.ListenAndServe(":4000", router)
	log.Print("I'm running")
}

func getFeed(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}