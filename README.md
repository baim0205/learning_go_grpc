# Authentication & CRUD gRPC Service

This project consists of two microservices: an Authentication Service and a CRUD Service. The Authentication Service handles user login and token validation, while the CRUD Service performs Create, Read, Update, and Delete operations on a simple database table (`items`). Access to the CRUD operations requires a valid JWT token, which is obtained from the Authentication Service.

## Getting Started

### Prerequisites
- **Go**: Make sure Go is installed (version >= 1.16).
- **MySQL/MariaDB**: A running MySQL or MariaDB instance for the CRUD Service to store data.
- **Postman**: Latest version with gRPC support to test the gRPC services.

### Installation
1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-repo/auth-crud-grpc.git
   cd auth-crud-grpc
