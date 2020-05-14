package main

import (
    "fmt"
    "net/http"
)

var port = os.Getenv("PORT") 

func helloWorld(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World")
}

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":" + port, nil)
}
