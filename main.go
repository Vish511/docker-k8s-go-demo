package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	}).Methods("GET")
	serverPort := os.Getenv("SERVER_PORT")

	fmt.Println("Server started at ", serverPort)
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), router)
}
