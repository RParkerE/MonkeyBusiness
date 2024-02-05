package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./pages"))))
    http.ListenAndServe(":3100", nil)
}