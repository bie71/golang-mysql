package utils

/*
Folder ini berfungsi untuk menampung fungsi atau kode yang sering digunakan di berbagai bagian aplikasi, namun tidak termasuk dalam logika inti aplikasi seperti yang ada di folder controllers atau models. Folder utils/ biasanya berisi kode untuk membantu tugas-tugas tertentu yang dapat digunakan di seluruh aplikasi, seperti validasi, format data, helper function, pengelolaan error, atau utility lainnya.
*/

// ParseBody adalah fungsi untuk membaca dan mengurai isi dari body permintaan HTTP (HTTP request).
// Fungsi ini menerima dua parameter:
// - r: objek http.Request yang berisi informasi tentang permintaan HTTP.
// - x: interface kosong yang akan diisi dengan data yang diurai dari body permintaan.
//
// Implementasi:
// 1. Membaca seluruh isi dari body permintaan menggunakan ioutil.ReadAll.
//    Jika terjadi kesalahan saat membaca body, fungsi akan berhenti dan tidak melakukan apa-apa.
// 2. Jika pembacaan body berhasil, data yang dibaca akan diubah menjadi slice byte.
// 3. Mengurai data JSON dari slice byte tersebut dan mengisinya ke dalam parameter x menggunakan json.Unmarshal.
//    Jika terjadi kesalahan saat mengurai JSON, fungsi akan berhenti dan tidak melakukan apa-apa.

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("[%s] %s %s", r.Method, r.URL.Path, r.RemoteAddr)

		// Lanjutkan ke handler berikutnya
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("[%s] %s %s - %v", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	APP_PORT   string
}

func Init() Config {
	// Inisialisasi konfigurasi via ENV

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		APP_PORT:   os.Getenv("APP_PORT"),
	}
}
