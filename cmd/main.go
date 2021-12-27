package main

import(
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func main(){
	router := mux.NewRouter()
	
	router.HandleFunc("/", getFeed)

	http.ListenAndServe(":4000", router)
}

func getFeed(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}