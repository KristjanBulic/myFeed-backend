package main

import(
	"os"
	"encoding/json"
	"log"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"github.com/ungerik/go-rss"
)

var url string

func init(){
	url = os.Getenv("URL")
}

func main(){
	router := mux.NewRouter()
	
	router.HandleFunc("/", getFeed)

	http.ListenAndServe(":4000", router)
	log.Print("I'm running")
}

func getFeed(w http.ResponseWriter, r *http.Request){
	
	resp, err := rss.Read(url, true)
	if err != nil {
		fmt.Println(err)
	}

	channel, err := rss.Regular(resp)
	if err != nil {
		fmt.Println(err)
	}

	data, _ := json.Marshal(channel.Item)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	fmt.Fprintf(w, string(data))
}