# e2e_simple

Paket `e2e_simple` menyediakan implementasi sederhana untuk enkripsi dan dekripsi menggunakan algoritma RSA dalam bahasa pemrograman Go.

## Penggunaan

### Instalasi

Untuk menggunakan paket ini, Anda perlu menambahkannya ke proyek Go Anda:

```bash
go get -u github.com/refaldyrk/e2e_simple
```

## Contoh Penggunaan
```go
package main

import (
	"fmt"
	"github.com/refaldyrk/e2e_simple"
)

func main() {
    // Enkripsi pesan
    message := "Hello, e2e_simple!"
    ciphertext, privateKey, err := e2e_simple.E2e_enc(message)
    if err != nil {
        fmt.Println("Gagal melakukan enkripsi:", err)
        return
    }

    // Dekripsi pesan
    decryptedMessage, err := e2e_simple.E2e_dec(privateKey, ciphertext)
    if err != nil {
        fmt.Println("Gagal melakukan dekripsi:", err)
        return
    }

    fmt.Println("Pesan Asli:", message)
    fmt.Println("Pesan Terdekripsi:", string(decryptedMessage))
}

```

## Catatan 
Pastikan untuk menangani kesalahan secara memadai untuk memastikan keandalan aplikasi.
Ukuran kunci RSA yang digunakan dalam contoh ini adalah 2048 bit. Sesuaikan sesuai dengan kebutuhan keamanan aplikasi Anda.
