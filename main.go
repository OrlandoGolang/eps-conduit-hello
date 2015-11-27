package main

import(
"fmt"
"log"
"flag"
"net/http"
)

var port = flag.String("port", "8080", "port to bind to")

func indexHandler( w http.ResponseWriter, r *http.Request){
fmt.Fprintf(w, "Hello world, I'm answering your request from backend port %s", *port)
}

func main(){
  flag.Parse()
  http.HandleFunc("/", indexHandler)
  log.Println("EPS Conduit Hello listening on port", *port)
  http.ListenAndServe(":"+*port ,nil)
}
