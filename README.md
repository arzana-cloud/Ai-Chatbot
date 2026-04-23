# Gemini AI Chatbot - Dockerized Go & JS

Proyek ini adalah aplikasi chatbot AI sederhana yang menggunakan **Google Gemini AI** sebagai otaknya. Dibangun dengan arsitektur **Server-Client** menggunakan **Golang** di sisi Backend dan **Vanilla JavaScript** di sisi Frontend, semuanya diorkestrasi menggunakan **Docker**.

---

## 🏗️ Arsitektur Sistem
- **Backend**: Golang (Gin Framework) + Google Generative AI SDK.
- **Frontend**: HTML5, CSS3, & JavaScript (Nginx).
- **Orchestration**: Docker Compose.

---

## 👥 Pembagian Tugas (Roadmap)

### 1. Backend (Raihan) - Folder `/backend`
- [ ] Inisialisasi Go Modules dan Setup Gin Gonic.
- [ ] Integrasi SDK Gemini AI.
- [ ] Pembuatan endpoint `POST /chat`.
- [ ] Konfigurasi CORS agar bisa diakses Frontend.
- [ ] Pembuatan `Dockerfile` untuk Backend.

### 2. Frontend (Collaborator) - Folder `/frontend`
- [ ] Desain UI Chat (HTML/CSS).
- [ ] Integrasi Fetch API untuk komunikasi ke Backend.
- [ ] Penanganan state (loading, error, display response).
- [ ] Pembuatan `Dockerfile` menggunakan Nginx.

---

## 🚀 Cara Menjalankan (Local Development)

### Prasyarat
- Docker dan Docker Compose terinstal di mesin Anda.
- API Key Gemini (Dapatkan di [Google AI Studio](https://aistudio.google.com/)).

### Langkah-langkah
1. **Clone Repositori**
   ```bash
   git clone [https://github.com/username/repo-name.git](https://github.com/username/repo-name.git)
   cd repo-name
   ```

2. **Setup Environment**
   Buat file `.env` di direktori utama dan tambahkan API Key Anda:
   ```env
   GEMINI_API_KEY=your_api_key_here
   ```

3. **Jalankan dengan Docker Compose**
   ```bash
   docker-compose up --build
   ```

4. **Akses Aplikasi**
   - **Frontend**: Buka `http://localhost:80` di browser Anda.
   - **Backend API**: Berjalan di `http://localhost:8080`.

---

## 📂 Struktur Direktori
```text
.
├── backend/
│   ├── main.go            # Logika API & Gemini Integration
│   ├── Dockerfile         # Docker Image untuk Go
│   └── go.mod
├── frontend/
│   ├── index.html         # UI Chatbot
│   ├── script.js          # Logika Fetch API
│   └── Dockerfile         # Docker Image untuk Nginx
├── docker-compose.yml     # Orkestrasi container
└── .env                   # API Keys (Jangan di-push ke git!)
```

---

## 🛠️ Tech Stack
- **Language**: [Go](https://go.dev/), [JavaScript](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
- **Framework API**: [Gin Gonic](https://gin-gonic.com/)
- **AI Engine**: [Google Gemini AI](https://ai.google.dev/)
- **Containerization**: [Docker](https://www.docker.com/)
- **Web Server**: [Nginx](https://www.nginx.com/)

---

## 📝 Catatan untuk Kontributor
- Pastikan setiap perubahan besar dibuat dalam *branch* baru.
- Update file `.env.example` jika ada penambahan variabel environment baru.
- Jalankan `go mod tidy` di folder backend setelah menambah dependensi baru.
```

### Tips Tambahan untuk Kamu:
1. **File `.gitignore`**: Pastikan kamu membuat file `.gitignore` dan masukkan `.env` di dalamnya agar API Key kamu tidak tersebar secara publik.
2. **File `.env.example`**: Buat file ini dengan isi `GEMINI_API_KEY=` (kosongkan nilainya) agar kawanmu tahu dia harus mengisi variabel tersebut di file `.env` miliknya sendiri.

Ada bagian dari struktur ini yang ingin kamu sesuaikan lagi?