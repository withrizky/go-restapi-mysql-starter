# 🚀 Go REST API MySQL Starter

Starter template untuk membangun REST API menggunakan Go (Golang) dengan Fiber, GORM, dan MySQL.  
Struktur project ini dirancang agar scalable, clean, dan mudah dikembangkan untuk project backend profesional.

---

## 🧩 Fitur
- ⚙️ Clean Architecture (`cmd/`, `internal/`)
- 🚀 Fiber sebagai HTTP framework
- 🗄️ GORM ORM untuk MySQL
- 🔐 Load environment via `.env`
- 🧱 Auto-migration model (otomatis bikin tabel)
- ✅ Siap di-deploy ke server

---

## ⚙️ Cara Menjalankan
1. Clone repository: 
- git clone https://github.com/withrizky/go-restapi-mongodb-starter.git
- cd go-restapi-mysql-starter
2. Install dependencies:
- go mod tidy
3. Buat file .env berdasarkan .env.example
4. Jalankan server : go run cmd/server/main.go
5. Endpoint example : 
- POST http://localhost:8080/api/users
- body :
{
  "name": "Rizky",
  "email": "rizky@example.com"
}
- respon diharapkan :
{
  "id": 1,
  "name": "Rizky",
  "email": "rizky@example.com",
  "created_at": "2025-10-30T15:10:00Z",
  "updated_at": "2025-10-30T15:10:00Z"
}