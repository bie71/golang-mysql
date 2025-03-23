package controllers

/*
Folder ini menyimpan kode untuk menangani logika bisnis dan interaksi dengan model. Controllers biasanya menangani request dari pengguna dan mengembalikan response yang sesuai. Dalam konteks aplikasi web atau API, controllers bisa berisi fungsi untuk mengelola berbagai route dan endpoint API.
*/

import (
	"encoding/json"           // Mengimpor package untuk menangani encoding dan decoding JSON
	"fmt"                     // Mengimpor package untuk print error atau output lainnya ke konsol
	"golang-mysql/pkg/models" // Mengimpor package models yang berisi logika untuk operasi database
	"golang-mysql/pkg/utils"  // Mengimpor package utils yang berisi helper functions
	"net/http"                // Mengimpor package untuk menangani request dan response HTTP
	"strconv"                 // Mengimpor package untuk mengonversi string ke tipe data lain seperti int

	"github.com/gorilla/mux" // Mengimpor Gorilla Mux untuk routing HTTP
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetALlBooks() // Memanggil fungsi GetAllBooks() dari models untuk mendapatkan semua buku

	w.Header().Set("Content-Type", "application/json") // Menetapkan tipe konten menjadi JSON

	w.WriteHeader(http.StatusOK) // Menetapkan status code HTTP 200 OK

	if err := json.NewEncoder(w).Encode(newBooks); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // Mengirimkan response dengan data buku dalam format JSON
	}
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Mendapatkan parameter ID dari URL

	bookId := vars["bookId"] // Mendapatkan nilai ID dari parameter

	ID, err := strconv.ParseInt(bookId, 0, 0) // Mengonversi ID menjadi int64

	if err != nil {
		fmt.Println("Error while parsing")
	}

	bookDetails, _ := models.GetBookById(ID) // Memanggil fungsi GetBookById() dari models untuk mendapatkan buku berdasarkan ID

	w.Header().Set("Content-Type", "application/json") // Menetapkan tipe konten menjadi JSON

	w.WriteHeader(http.StatusOK) // Menetapkan status code HTTP 200 OK

	if err := json.NewEncoder(w).Encode(bookDetails); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // Mengirimkan response dengan data buku dalam format JSON
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{} // Membuat objek buku baru

	utils.ParseBody(r, CreateBook) // Mengurai body permintaan HTTP dan mengisinya ke objek buku baru

	b := CreateBook.CreateBook() // Membuat buku baru menggunakan metode CreateBook() dari objek buku

	w.Header().Set("Content-Type", "application/json") // Menetapkan tipe konten menjadi JSON

	w.WriteHeader(http.StatusOK) // Menetapkan status code HTTP 200 OK

	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // Mengirimkan response dengan data buku baru dalam format JSON
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Mendapatkan parameter ID dari URL

	bookId := vars["bookId"] // Mendapatkan nilai ID dari parameter

	ID, err := strconv.ParseInt(bookId, 0, 0) // Mengonversi ID menjadi int64

	if err != nil {
		fmt.Println("Error while parsing")
	}

	deletedBook := models.DeleteBook(ID) // Menghapus buku berdasarkan ID menggunakan metode DeleteBook() dari models

	w.Header().Set("Content-Type", "application/json") // Menetapkan tipe konten menjadi JSON

	w.WriteHeader(http.StatusOK) // Menetapkan status code HTTP 200 OK

	if err := json.NewEncoder(w).Encode(deletedBook); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // Mengirimkan response dengan data buku yang telah dihapus dalam format JSON
	}
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{} // Membuat objek buku baru untuk diupdate

	utils.ParseBody(r, updateBook) // Menggunakan fungsi ParseBody untuk memparsing request body menjadi objek updateBook

	vars := mux.Vars(r) // Mengambil variabel path (ID) dari request URL

	bookId := vars["bookId"] // Mendapatkan ID buku dari URL

	ID, err := strconv.ParseInt(bookId, 0, 0) // Mengonversi string ID menjadi tipe int

	if err != nil {
		fmt.Println("error while parsing") // Menangani error jika ID tidak valid
	}

	bookDetails, db := models.GetBookById(ID) // Memanggil fungsi GetBookById dari models untuk mendapatkan buku berdasarkan ID

	if updateBook.Name != "" { // Jika ada perubahan pada nama buku
		bookDetails.Name = updateBook.Name // Update nama buku
	}

	if updateBook.Author != "" { // Jika ada perubahan pada penulis buku
		bookDetails.Author = updateBook.Author // Update penulis buku
	}

	if updateBook.Publication != "" { // Jika ada perubahan pada penerbit buku
		bookDetails.Publication = updateBook.Publication // Update penerbit buku
	}
	db.Save(&bookDetails) // Menyimpan perubahan ke dalam database

	w.Header().Set("Content-Type", "pkglication/json") // Menetapkan tipe konten menjadi JSON
	w.WriteHeader(http.StatusOK)                       // Menetapkan status code HTTP 200 OK

	if err := json.NewEncoder(w).Encode(bookDetails); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError) // Mengirimkan response dengan data buku yang diperbarui dalam format JSON
	}
}
