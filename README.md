# Belajar Membuat Authentication & CRUD gRPC Service Serderhana

Project pembelajaran ini terdiri dari dua microservices: Layanan Autentikasi dan Layanan CRUD. Layanan Autentikasi menangani login pengguna dan validasi token, sedangkan Layanan CRUD melakukan operasi Create, Read, Update, dan Delete pada tabel database sederhana (items). Akses ke operasi CRUD memerlukan token JWT yang valid, yang diperoleh dari Layanan Autentikasi.

## Getting Started

### Prerequisites
- **Go**: Pastikan Go sudah terinstal (versi >= 1.16).
- **MySQL/MariaDB: Sebuah instance MySQL atau MariaDB yang sedang berjalan untuk menyimpan data Layanan CRUD.
- **Postman**: Versi terbaru dengan dukungan gRPC untuk menguji layanan gRPC.

### Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-repo/auth-crud-grpc.git
   cd auth-crud-grpc

2. **Setup Database**:
   ##### Buat database MySQL atau MariaDB :
   ```sql
   CREATE DATABASE crud_service;
   USE crud_service;
   CREATE TABLE items (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255) NOT NULL
   );

3. **Install Dependencies**:
   ##### Di dalam direktori `auth-service` dan `crud-service`, jalankan perintah berikut:
   ```bash
   go mod tidy


2. **service CRUD**:
   ```bash
   cd crud-service
   go run main.go
CRUDService akan berjalan di `localhost:50052`

### Running the Services

1. **service auth**:
   ```bash
   cd auth-service
   go run main.go
AuthService akan berjalan di `localhost:50051`

2. **service CRUD**:
   ```bash
   cd crud-service
   go run main.go
CRUDService akan berjalan di `localhost:50052`


### Pengujian dengan Postman (gRPC)
**Service Auth**:

1. Buka Postman dan pilih gRPC.
2. Masukkan URL: localhost:50051.
3. Pilih metode Login (atau metode autentikasi lainnya yang sudah diimplementasikan).
4. Kirim permintaan dengan parameter yang diperlukan, misalnya:

   ```json
    {
    "username": "your-username",
    "password": "your-password"
    }
5. Anda akan menerima JWT token jika login berhasil. Simpan token ini untuk digunakan di Layanan CRUD.

**Service CRUD**:

1. Buka permintaan gRPC baru di Postman.
2. Masukkan URL: localhost:50052.
3. Pilih salah satu metode CRUD (Create, Read, Update, Delete).
4. Kirim permintaan dengan token dan data yang relevan.

    Contoh metode Create dan Read : 

   ```json
    {
    "username": "your-username",
    "password": "your-password"
    }


    {
    "token": "your-jwt-token",
    "id": 1
    }

### Teknologi yang Digunakan
##### Go: Digunakan untuk membangun kedua layanan (Autentikasi dan CRUD).
##### gRPC: Untuk komunikasi antara layanan dan klien.
##### JWT: Untuk autentikasi dan validasi token.
##### MySQL/MariaDB: Database yang digunakan untuk operasi CRUD.
### Struktur Folder

```json
.
├── auth-service
│   ├── main.go
│   ├── server.go
│   └── proto
│       └── protoc-auth
│           └── auth.proto
├── crud-service
│   ├── main.go
│   ├── server.go
│   ├── database
│   │   └── database.go
│   └── proto
│       └── protoc
│           └── crud.proto
└── README.md
