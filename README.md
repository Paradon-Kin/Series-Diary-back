# 🍿 Series Diary - Backend (Go)

The robust and scalable RESTful API for **Series Diary**, a personal series tracking application. Built with **Go (Golang)** and designed with a focus on clean architecture, security, and efficient data management.

## 🚀 Key Features
* **Secure Authentication**: Implemented **JWT (JSON Web Token)** for stateless authentication and **Bcrypt** for secure password hashing.
* **Relational Database Management**: Powered by **MySQL** and **GORM (ORM)**, featuring automatic schema migration for Users, Series, and Watch History.
* **Data Isolation (Multi-tenant Logic)**: Every record is strictly tied to a `user_id` using GORM queries, ensuring users can only access their own private data.
* **Intelligent Progress Tracking**: Custom logic to track and update "Current Episodes" and "Total Episodes," enabling progress calculation and categorization.
* **Soft Deletion**: Utilized `gorm.Model` to handle deletions safely, allowing for potential data recovery and audit logs.
* **CORS Configuration**: Pre-configured Cross-Origin Resource Sharing to securely communicate with the React frontend.
* **Env Management**: Centralized configuration for sensitive data (DB Credentials, JWT Secrets) using `godotenv`.

## 🛠 Tech Stack
* **Language**: Go (Golang)
* **Web Framework**: Gin Gonic
* **ORM**: GORM (MySQL Driver)
* **Database**: MySQL
* **Security**: JWT (golang-jwt), Bcrypt
* **Configuration**: Godotenv

## 🏗 Project Structure
The project follows a modular and clean structure:
- `/config` - Database connection and Environment setup
- `/controllers` - API Logic (Authentication, Series CRUD, History/Ratings)
- `/middlewares` - Security layer (JWT Auth Middleware)
- `/models` - GORM database schemas & relations
- `/routes` - API endpoint definitions and CORS setup
- `main.go` - Application entry point & server initialization
- `.env` - Environment variables (Secrets)

## 🔧 Installation & Setup

**1. Clone the repository**:
```bash
git clone [https://github.com/Paradon-Kin/Series-Diary-back.git](https://github.com/Paradon-Kin/Series-Diary-back.git)
cd Series-Diary-back
```

**2. Environment Setup**:
Create a `.env` file in the root directory and configure your credentials:
```env
DB_DSN="username:password@tcp(127.0.0.1:3306)/series_diary?charset=utf8mb4&parseTime=True&loc=Local"
JWT_SECRET="your_secure_secret_key"
PORT="8080"
```

**3. Install Dependencies**:
```bash
go mod tidy
```

**4. Run the Server**:
```bash
go run main.go
```

## 🛣 API Documentation

### 🔓 Public Endpoints
| Method | Endpoint    | Description                     |
| :----- | :---------- | :------------------------------ |
| GET    | `/ping`     | Health check for API status     |
| POST   | `/register` | User registration (Bcrypt hash) |
| POST   | `/login`    | Authentication & JWT Issuance   |

### 🔐 Protected Endpoints (Requires `Authorization: Bearer <Token>`)
| Method | Endpoint           | Description                        |
| :----- | :----------------- | :--------------------------------- |
| GET    | `/api/series`      | Fetch all series for the user      |
| POST   | `/api/series`      | Add a new series to the diary      |
| DELETE | `/api/series/:id`  | Remove a series (Soft Delete)      |
| GET    | `/api/history`     | Get all watch history records      |
| POST   | `/api/history`     | Update current episode progress    |
| POST   | `/api/rate`        | Update series rating (1-10 scale)  |
