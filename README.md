# 📌 Digital Lending Book API
The Digital Lending Book API is an application designed to provide seamless access to a digital book lending service. It enables users to browse a wide collection of books, borrow them online, and enjoy reading anytime and anywhere

---

## 🔗 API Endpoints

| Method     | Endpoint       | Deskripsi                                                                   |
| ---------- | -------------- | --------------------------------------------------------------------------- |
| **POST**   | `/auth/signin` | Login user dengan email & password. Return JWT token.                       |
| **POST**   | `/auth/signup` | Registrasi user baru. Return token jika berhasil.                           |
| **GET**    | `/books`       | Ambil daftar buku dengan pagination (`page`, `limit`).                      |
| **POST**   | `/books`       | Tambah buku baru (admin only, butuh bearer token).                          |
| **PUT**    | `/books/{id}`  | Update detail buku (admin only).                                            |
| **DELETE** | `/books/{id}`  | Hapus buku (admin only).                                                    |
| **POST**   | `/borrow`      | Pinjam buku berdasarkan `book_id`. Return record id, borrow_date, due_date. |
| **POST**   | `/return`      | Kembalikan buku berdasarkan `record_id`. Return record id, return_date.     |


---

## ⚙️ Prerequisites

Sebelum menjalankan project, pastikan sudah menginstall software berikut:

### 1. Golang
- [Download Go](https://go.dev/dl/)  
- Verifikasi instalasi:

    go version

### 2. golang-migrate
- [Install Guide](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)  
- Install via Homebrew (Mac/Linux):

    brew install golang-migrate

- Verifikasi:

    migrate -version

### 3. Docker & Docker Compose
- [Install Docker](https://docs.docker.com/get-docker/)  
- Verifikasi instalasi:

    docker --version  
    docker-compose --version

### 4. MySQL (mysql client)
- Install MySQL sesuai OS.  
- Verifikasi:

    mysql --version

---

## 🔑 Environment Variables

Buat file `.env` di root project. Berikut contoh template:

    DB_CONNECTION=mysql://user:password@localhost:3306/taskdb?parseTime=True&loc=Local
    ACCESS_SECRET_KEY=my-secret-key
    APP_PORT=8080
    APP_HOST=localhost
    DB_ROOT_PASSWORD=${DB_ROOT_PASSWORD:-verysecret}
    DB_DATABASE=${DB_DATABASE:-db_book_lending}
    DB_USERNAME=${DB_USERNAME:-root}

- **`DATABASE_URL`** → URL koneksi ke database PostgreSQL.  
- **`ACCESS_SECRET_KEY`** → Secret key untuk JWT authentication.  
- **`APP_PORT`** → digunakan untuk PORT
- **`APP_PORT`** → digunakan untuk menentukan HOST
- **`DB_ROOT_PASSWORD`** → Root password for MySQL database
- **`DB_DATABASE`** → Database name for the application
- **`DB_USERNAME`** → Username for database connection

---

## 🗄️ Database Migration

Gunakan **golang-migrate** untuk mengatur migrasi database.

### Menjalankan migration (up)

    migrate -path apps/migrations -database "mysql://user:password@tcp(host:port)/dbname?query" up

### Rollback migration (down)

    migrate -path apps/migrations -database "mysql://user:password@tcp(host:port)/dbname?query" down

### Membuat migration baru

    migrate create -ext sql -dir apps/migrations -seq create_tasks_table

---

## 🚀 Menjalankan Aplikasi

### 1. Menjalankan secara lokal

    go run main.go

### 2. Menjalankan dengan Docker Compose

    docker-compose up --build

Aplikasi akan tersedia di:  
👉 http://localhost:8080

---

## 🧪 Menjalankan Unit Tests

Jalankan semua unit test dengan:

---

## 📦 Dependencies

Daftar dependency utama dalam project ini:

### 1. [Fiber](https://github.com/gofiber/fiber)
Framework web **Go** yang cepat, ringan, dan mirip dengan Express.js pada Node.js. Digunakan untuk routing, middleware, dan request handling.

### 2. [Golang JWT](https://github.com/golang-jwt/jwt)
Library untuk membuat dan memverifikasi **JSON Web Token (JWT)**. Dipakai untuk otentikasi dan otorisasi user.

### 3. [GORM](https://gorm.io/index.html)
Library untuk melakukan koneksi dan query ke database dengan mudah.

### 4. [Gokit](https://github.com/vizucode/gokit)
Paket library yang ready-to-use tanpa perlu setting dll.

---

## 📝 Catatan Tambahan
- Pastikan `.env` sudah dikonfigurasi dengan benar.  
- Jika menjalankan dengan Docker Compose, database akan otomatis berjalan dalam container `db`.  
- Gunakan `migrate create` untuk menambah migrasi baru sesuai kebutuhan.