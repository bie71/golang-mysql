package main

import (
	"log"
	"net/http"

	"golang-mysql/pkg/routers"
	"golang-mysql/pkg/utils"

	"github.com/gorilla/mux"
)

func main() {

	cfg := utils.Init()

	r := mux.NewRouter()
	routers.RegisterBookStoreRoutes(r)

	http.Handle("/", r)

	log.Println("Server is running on port ", cfg.APP_PORT)

	log.Fatal(http.ListenAndServe(":"+cfg.APP_PORT, r))
}
