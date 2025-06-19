## Deploying go on vercel

Repositori ini dirancang sebagai starter pack untuk kamu yang ingin mempelajari cara mendeploy aplikasi Go di Vercel, khususnya bagi yang sedang belajar membuat REST API dengan Go menggunakan Gin. Cocok banget buat kamu yang ingin mencoba hosting secara gratis sebagai bagian dari proses belajar.

demo:

- https://go-starter.vercel.app/api/categories
- https://go-starter.vercel.app/api/users

## Struktur folder

```
.
├── api
│   ├── index.go
│   └── _pkg
│       ├── handlers
│       │   ├── category.handler.go
│       │   └── user.handler.go
│       └── models
│           ├── category.go
│           └── user.go
├── cmd
│   └── main.go
├── go.mod
├── go.sum
└── vercel.json
```

## Penjelasan

- **`api/`**  
  Folder ini digunakan oleh Vercel untuk fungsi serverless. Setiap file `.go` di dalam folder ini akan diproses sebagai endpoint secara otomatis saat dideploy.

  - **`index.go`**  
    File endpoint utama. Bisa berfungsi sebagai root path atau sebagai router awal yang nantinya akan dibaca divercel.

  - **`_pkg/`**  
    Folder internal untuk menyimpan kode modular. Tidak akan diproses sebagai endpoint oleh Vercel.

- **`cmd/`**  
  Direktori untuk menjalankan aplikasi secara lokal (tanpa Vercel). Berguna untuk testing dan development.

  - **`main.go`**  
    Entry point utama saat aplikasi dijalankan secara lokal menggunakan `go run`.

- **`vercel.json`**  
  Konfigurasi deployment untuk Vercel. Mengatur runtime, routes, dan pengaturan deployment lainnya.

Contoh konfigurasi `vercel.json` pada project ini:

```json
{
  "trailingSlash": false,
  "rewrites": [
    {
      "source": "/api/:path*",
      "destination": "/api/index.go"
    }
  ],
  "build": {
    "env": {
      "GO_BUILD_FLAGS": "-ldflags '-s -w'"
    }
  }
}
```

## Instalasi

1. Clone repositori ini

```bash
git clone https://github.com/yuefii/go-vercel.git
cd go-vercel
```

2. Install dependencies

```bash
go mod tidy
```

3. Menjalankannya secara lokal

```bash
go run cmd/main.go
```

**Note:** Pastikan file `main.go` konfigurasi routenya sama dengan `index.go` agar tidak terjadi error saat melakukan deploy.

Terima kasih semoga bermanfaat.
