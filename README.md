# golang-mysql
# CRUD Book API

CRUD Book API adalah aplikasi sederhana yang dibangun menggunakan Golang, GORM, Mux, dan MySQL. Aplikasi ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data buku.

## Teknologi yang Digunakan

- **Golang**: Bahasa pemrograman yang digunakan untuk membangun aplikasi.
- **GORM**: ORM (Object Relational Mapping) untuk Golang yang memudahkan interaksi dengan database.
- **Mux**: Router HTTP untuk menangani permintaan dan rute.
- **MySQL**: Database yang digunakan untuk menyimpan data buku.
- **Go Dotenv**: Paket untuk memuat variabel lingkungan dari file `.env`.

## Fitur

- Menambahkan buku baru
- Mengambil daftar semua buku
- Mengambil detail buku berdasarkan ID
- Memperbarui informasi buku
- Menghapus buku

## Prerequisites

Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/)
- [MySQL](https://www.mysql.com/downloads/)
- [Go Dotenv](https://github.com/joho/godotenv)

## Instalasi

1. **Clone repositori ini**:
   ```bash
   git clone https://github.com/bie71/golang-mysql.git
   cd repo-name
