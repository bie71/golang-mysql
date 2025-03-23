package routers

/*
Folder ini digunakan untuk mendefinisikan semua route atau endpoint dalam aplikasi. Biasanya, di dalam folder ini, Anda akan membuat file yang berisi fungsi untuk mengatur URL path dan menghubungkannya dengan controller atau handler yang sesuai. Di Go, routing sering dilakukan dengan menggunakan package seperti gorilla/mux atau http.ServeMux, dan folder ini berfungsi untuk menyusun semua route yang digunakan dalam aplikasi.
*/

import (
	"golang-mysql/pkg/controllers"
	"golang-mysql/pkg/utils"

	"github.com/gorilla/mux"
)

func RegisterBookStoreRoutes(r *mux.Router) {
	r.HandleFunc("/api/books", controllers.GetBook).Methods("GET")
	r.HandleFunc("/api/books/{bookId}", controllers.GetBookById).Methods("GET")
	r.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{bookId}", controllers.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{bookId}", controllers.DeleteBook).Methods("DELETE")
	r.Use(utils.LoggingMiddleware)
}
