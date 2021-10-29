package main

import("net/http")

func main() {
 http.HandleFunc("/", HandleHome)
 http.ListenAndServe(":8081", nil)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Se vienen"))
}
