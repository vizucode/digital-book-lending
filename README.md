# 📌 Digital Lending Book API
The Digital Lending Book API is an application designed to provide seamless access to a digital book lending service. It enables users to browse a wide collection of books, borrow them online, and enjoy reading anytime and anywhere

---

## 🔗 API Endpoints

| Method | Endpoint      | Deskripsi                               |
|--------|---------------|-----------------------------------------|

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

- **`DATABASE_URL`** → URL koneksi ke database PostgreSQL.  
- **`ACCESS_SECRET_KEY`** → Secret key untuk JWT authentication.  
- **`APP_PORT`** → digunakan untuk PORT
- **`APP_PORT`** → digunakan untuk menentukan HOST

---

## 🗄️ Database Migration

Gunakan **golang-migrate** untuk mengatur migrasi database.

### Menjalankan migration (up)

    migrate -path apps/migrations -database "$DATABASE_URL" up

### Rollback migration (down)

    migrate -path apps/migrations -database "$DATABASE_URL" down

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