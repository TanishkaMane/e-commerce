# 🛒 E-Commerce Backend API (Golang + Gin + MongoDB)

This is a backend service for an e-commerce platform, built entirely in **Golang** using the **Gin** web framework. It includes authentication, role-based access control, MongoDB integration, and well-structured RESTful API design — showcasing real-world backend architecture and best practices.

---

## 🚀 Features

- 🔐 **JWT-based Authentication** with **Refresh Tokens**
- 🧩 **Middleware for Protected Routes**
- 🛠️ **Role-based Access Control** (`/admin/*` vs `/users/*`)
- 🗃️ **MongoDB Integration** with full **CRUD support**
- 🧪 Cleanly separated controller logic
- ⚙️ Built using **Gin** — a lightweight, high-performance web framework
- 📦 Modular design for **scalability and clarity**

---

## 🧱 Routes Overview

### 👤 User Routes
- `POST /users/signup` — User Registration  
- `POST /users/login` — User Login with JWT issuance

### 🛒 Admin Routes
- `POST /admin/addProduct` — Add a new product  
- `GET /admin/productView` — View all products  
- `GET /admin/search` — Search products by query

---

## 🛡️ Security Highlights

- ✅ JWT Authentication with access + refresh token flow  
- ✅ Middleware that verifies tokens before accessing protected routes  
- ✅ Admin-only routes to secure product management  
- ✅ Token expiration and refresh logic to maintain session security

---

## 🛢️ Tech Stack

| Component | Technology      |
|----------|------------------|
| Language  | Go (Golang)      |
| Framework | Gin              |
| Database  | MongoDB          |
| Auth      | JWT + Refresh Token |
| Middleware | Custom auth middleware |
| Data Model | MongoDB Collections for Users and Products |

---

## 📚 Purpose

This project was built to deeply understand backend systems using **Go** — including REST API design, middleware, authentication, and database interaction using **MongoDB**.

It simulates real-world features like:
- Token-based authentication flow
- Admin/user route separation
- Product management system

---

## 🧪 Sample Workflow

1. **User signs up** → token issued  
2. **User logs in** → access + refresh token returned  
3. **Access token expires** → refresh token used to get a new one  
4. **Admin adds/searches products** using protected endpoints

---

## 🛠 Setup Instructions

```bash
# Clone the repository
git clone https://github.com/your-username/e-commerce.git

# Navigate into the project
cd e-commerce

# Run the Go server (assuming you have main.go)
go run main.go
