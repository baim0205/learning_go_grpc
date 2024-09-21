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
