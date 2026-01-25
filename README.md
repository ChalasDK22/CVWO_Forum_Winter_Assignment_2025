# CVWO-WinterAssignment
Name: Vu Dang Khoa

https://chalasdk-cvwo-winter-assignment-web-app.vercel.app/

Frontend Repo: https://github.com/ChalasDK22/CVWO_WinterAssignment_WebApp.git  
Backend Repo: https://github.com/ChalasDK22/CVWO_Forum_Winter_Assignment_2025.git

## Functionalities
- Register / Login (username)
- CRUD Topics, Posts, Comments
- Pagination (topics/posts/comments)

## Local Setup

1) Clone
```bash
git clone https://github.com/ChalasDK22/CVWO_WinterAssignment_WebApp.git
git clone https://github.com/ChalasDK22/CVWO_Forum_Winter_Assignment_2025.git
```

2) Start MySQL + phpMyAdmin
```bash
cd CVWO_Forum_Winter_Assignment_2025
docker compose up -d
```
phpMyAdmin: http://localhost:8081

3) Backend env  
   Create `CVWO_Forum_Winter_Assignment_2025/.env`
```env
ENV=local
FORUM_PORT=8080
FORUM_JWT=your_secret_key
DATABASE_URL=adminuser:12345678@tcp(localhost:3306)/defaultdb?parseTime=true
```
4) Import database schema/data (SQL file).
```bash
mysql -h localhost -P 3306 -u adminuser -p defaultdb < DB/Scripts/forum-db-3.sql
```
5) Run backend
```bash
cd CVWO_Forum_Winter_Assignment_2025
go mod tidy
go run ./API/cmd/main.go
```
Backend: http://localhost:8080

6) Update frontend baseUrl (Login API component)  
   Change:
```ts
export const baseUrl = "https://cvwo-forum-winter-assignment-2025.onrender.com/";
```
To:
```ts
export const baseUrl = "http://localhost:8080/";
```

7) Run frontend
```bash
cd CVWO_WinterAssignment_WebApp
npm install
npm install @mui/material @emotion/react @emotion/styled
npm run dev
```
Frontend: http://localhost:5173

