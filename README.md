# Katalog Musik

## Deskripsi Proyek
Proyek katalog musik adalah aplikasi backend sederhana yang dibuat dengan Golang untuk mengelola koleksi musik. Aplikasi ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data musik.

## Fitur
- Tambah lagu baru ke katalog
- Lihat daftar lagu
- Update informasi lagu
- Hapus lagu dari katalog
- Pencarian lagu berdasarkan kriteria (artis, genre, tahun)

## Prasyarat
- Golang 1.16 atau versi lebih baru
- Database (PostgreSQL/MySQL)
- Git

## Instalasi

### Kloning Repositori
```bash
git clone https://github.com/username/katalog-musik.git
cd katalog-musik
```

### Konfigurasi
1. Salin file `.env.example` menjadi `.env`
2. Atur konfigurasi database di file `.env`

### Install Dependensi
```bash
go mod download
```

### Migrasi Database
```bash
go run migrations/migrate.go
```

## Menjalankan Aplikasi
```bash
go run main.go
```

## Struktur Proyek
```
.
├── cmd/
│   └── server/
├── internal/
│   ├── handlers/
│   ├── models/
│   └── repository/
├── migrations/
├── configs/
├── .env
├── go.mod
└── README.md
```

## Endpoint API
- `POST /musik` - Tambah lagu baru
- `GET /musik` - Daftar semua lagu
- `GET /musik/{id}` - Detail lagu
- `PUT /musik/{id}` - Update lagu
- `DELETE /musik/{id}` - Hapus lagu

## Teknologi
- Golang
- Database (pilih salah satu)
- Framework routing
- ORM/Query builder

## Kontribusi
1. Fork repositori
2. Buat branch fitur (`git checkout -b fitur/AturanKontribusi`)
3. Commit perubahan (`git commit -m 'Tambah fitur baru'`)
4. Push ke branch (`git push origin fitur/AturanKontribusi`)
5. Buat Pull Request

## Lisensi
[Tentukan Lisensi Anda, misalnya MIT]

## Kontak
[Nama Anda]
[Email/Profil GitHub]
