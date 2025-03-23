package models

/*
Folder ini digunakan untuk mendefinisikan struktur data atau model yang  berhubungan dengan database. Misalnya, di sini bisa terdapat file yang mendefinisikan struktur tabel dalam database atau objek yang digunakan untuk interaksi dengan data (misalnya, dengan menggunakan ORM seperti GORM).
*/

import (
	"golang-mysql/pkg/config"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Connect() // Menjalankan fungsi Connect() dari package config untuk menghubungkan ke database

	db = config.GetDB() // Mengambil instance koneksi database yang sudah disetting di config dan menyimpannya di db

	db.AutoMigrate(&Book{}) // Melakukan migrasi otomatis untuk tabel "books", jika belum ada akan dibuatkan tabel
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // Membuat record baru (tapi belum disimpan)

	db.Create(&b) // Menyimpan data buku yang ada pada objek b ke dalam database

	return b // Mengembalikan objek buku yang baru saja dibuat
}

func GetALlBooks() []Book {
	var books []Book // Mendeklarasikan slice yang akan menampung semua buku

	db.Find(&books) // Mengambil semua buku yang ada di database dan menyimpannya di slice Books

	return books // Mengembalikan slice Books yang berisi semua data buku
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book // Membuat variabel getBook untuk menyimpan buku yang ditemukan

	db := db.Where("ID=?", id).Find(&getBook) // Menemukan buku berdasarkan ID dan menyimpannya di getBook

	return &getBook, db // Mengembalikan pointer ke objek getBook dan objek db (untuk error handling)
}

func DeleteBook(id int64) Book {
	var book Book // Membuat variabel book untuk menyimpan buku yang akan dihapus

	db.Where("ID=?", id).Delete(&book) // Menghapus buku berdasarkan ID

	return book // Mengembalikan objek buku yang telah dihapus
}

func UpdateBook(id int64) Book {
	var book Book // Membuat variabel book untuk menyimpan buku yang akan diupdate

	db.Where("ID=?", id).Find(&book) // Menemukan buku berdasarkan ID dan menyimpannya di book

	db.Save(&book) // Menyimpan perubahan pada buku

	return book // Mengembalikan objek buku yang telah diupdate
}
