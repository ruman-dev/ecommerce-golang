# 🛒 E-commerce Backend (Golang + Gin + PostgreSQL)

A scalable backend API for an eCommerce system built using **Golang**, **Gin**, and **PostgreSQL**.
This project follows a clean architecture approach with authentication, product management, and JWT-based authorization.

---

## 🚀 Features

* 🔐 User Registration & Login (JWT Authentication)
* 🛡️ Protected Routes with Middleware
* 📦 Product Management (CRUD)
* 🗄️ PostgreSQL Database Integration (sqlx)
* 🔑 Password Hashing using bcrypt
* ⚡ Clean project structure (Handler → Service → Repository ready)

---

## 🏗️ Tech Stack

* **Language:** Go (Golang)
* **Framework:** Gin
* **Database:** PostgreSQL
* **ORM/Driver:** sqlx
* **Authentication:** JWT
* **Environment Config:** godotenv

---

## ⚙️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/ruman-dev/ecommerce-golang.git
cd ecommerce-golang
```

---

### 2. Install dependencies

```bash
go mod tidy
```

---

### 3. Setup environment variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db
SSLMODE=disable

JWT_SECRET=your-secret-key
```

---

### 4. Setup Database

#### Create Users Table

```sql
CREATE TABLE users (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    name VARCHAR(100),
    age INT,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Create Products Table

```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC NOT NULL,
    quantity INT DEFAULT 0,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

### 5. Run the server

```bash
go run .
```

Server will start at:

```
http://localhost:8080
```

---

## 🔑 API Endpoints

### 🟢 Public Routes

| Method | Endpoint         | Description   |
| ------ | ---------------- | ------------- |
| POST   | /api/v1/register | Register user |
| POST   | /api/v1/login    | Login user    |

---

### 🔒 Protected Routes (JWT Required)

> Add header:

```
Authorization: Bearer <your_token>
```

| Method | Endpoint               | Description       |
| ------ | ---------------------- | ----------------- |
| POST   | /api/v1/create-product | Create product    |
| GET    | /api/v1/products       | Get all products  |
| GET    | /api/v1/product/:id    | Get product by ID |
| DELETE | /api/v1/product/:id    | Delete product    |

---
