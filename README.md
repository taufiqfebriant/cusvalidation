# Custom Validation di Fiber dengan `go-playground/validator`  

Implementasi custom validation di Fiber menggunakan `go-playground/validator` dengan response yang lebih rapi dan mudah dibaca.  

## Fitur  
- Validasi request body dengan `go-playground/validator`  
- Terjemahan error otomatis menggunakan `universal-translator` dan `locales`  
- Mapping error menggunakan tag `json` dari struct

## Cara Menjalankan  
1. Clone repo ini  
2. Install dependensi:  
   ```sh
   go mod tidy
   ```
3. Jalankan aplikasi dengan `go run main.go`:  
   ```sh
   go run main.go
   ```
   Atau, untuk development dengan live-reloading, Anda bisa menggunakan `air`:  
   ```sh
   air
   ```

4. (Opsional) Terdapat [Bruno](https://github.com/usebruno/bruno) (alternatif dari Postman) collection di folder `bruno` untuk testing endpoint-endpoint yang tersedia.

## Contoh Response  
Jika validasi gagal, response akan berbentuk seperti ini:  
```json
{
  "errors": {
    "email": "Email is a required field",
    "name": "Name is a required field",
    "password": "Password is a required field"
  }
}
```

## Teknologi yang Digunakan  
- [Fiber](https://gofiber.io/)  
- [go-playground/validator](https://github.com/go-playground/validator)  
- [go-playground/locales](https://github.com/go-playground/locales)  
- [go-playground/universal-translator](https://github.com/go-playground/universal-translator)