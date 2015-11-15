package main

import(
"fmt"
"log"
"net/http"
)

func indexHandler( w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "Hello world, I'm answering your request from backend ____")

}

func main(){
http.HandleFunc("/", indexHandler)
port := ":8080"
log.Println("EPS Conduit Hello listening on port", port)
http.ListenAndServe(port ,nil)
}
