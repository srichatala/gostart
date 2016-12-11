package main

import "net/http"

func main() {
    http.HandleFunc("/",someFunc)
    http.HandleFunc("/about",aboutFunc)
    http.ListenAndServe(":8088",nil)
}

func someFunc(w http.ResponseWriter, req *http.Request){
    w.Write([]byte("Hello world, welcome to http"))
}

func aboutFunc(w http.ResponseWriter, req *http.Request){
    w.Write([]byte("About page"))
}
