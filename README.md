# 🍿 Series Diary - Backend (Go)

The robust and scalable backend for **Series Diary**, a personal series tracking application. Built with **Go** and designed with a focus on security, data integrity, and efficient API performance.

## 🚀 Key Features
* **Secure Authentication**: Implemented **JWT (JSON Web Token)** for stateless authentication and **Bcrypt** for secure password hashing.
* **RESTful API Design**: Clean and modular endpoints for managing series, watch history, and user ratings.
* **Relational Database**: Powered by **MySQL** with **GORM (ORM)**, featuring automatic schema migration and relational mapping.
* **Data Isolation**: Every record is strictly tied to a `user_id`, ensuring a private and secure experience for each user (Multi-tenant logic).
* **Soft Deletion**: Utilizes GORM's `gorm.Model` to handle record deletions safely, preventing permanent data loss.
* **Environment Configuration**: Managed sensitive data (Database DSN, JWT Secret) via `.env` for production-grade security.

## 🛠 Tech Stack
* **Language**: Go (Golang)
* **Web Framework**: Gin Gonic
* **ORM**: GORM
* **Database**: MySQL
* **Security**: JWT, Bcrypt, Middleware-based Authorization
* **Config Management**: godotenv

## 🏗 Project Structure
```text
/config      - Database connection and environment setup
/controllers - Request handling logic (Auth, Series, Ratings)
/middlewares - Security layer (JWT verification)
/models      - GORM schema definitions
/routes      - API endpoint mapping and CORS setup
main.go      - Application entry point

##🔧 Installation & Setup