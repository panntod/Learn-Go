# Belajar Dasar GO

## Deskripsi

Dalam repo ini, berisikan documentation belajar go pribadi saya untuk mempermudah belajar bagi pemula untuk mendalami bahasa pemrograman GO, materi ini didapat dari [Programmer Zaman Now - Belajar GO](https://www.youtube.com/playlist?list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw).

### Installation

Pastikan sudah menginstal GO msi atau download menggunakan link:

[Download Go for Windows](https://go.dev/dl/go1.21.5.windows-amd64.msi).

## Tipe Data
<hr>

### String

Seperti dibahasa pemrograman lain, GO juga memiliki tipe data String yang ditandai dengan "...", '...'
```go
nama:= "Pandhu"
asal:= "Bogor"
```

### Integer
Berbeda dengan bahasa lain, GO memiliki beberapa jenis integer dan batas minimum dan maksimum dari sebuah jenis Integer
- `int8`: Bilangan bulat bertanda 8-bit. Rentang nilai dari -128 hingga 127.
- `uint8` (atau byte): Bilangan bulat tak bertanda 8-bit. Rentang nilai dari 0 hingga 255.
- `int16`: Bilangan bulat bertanda 16-bit. Rentang nilai dari -32768 hingga 32767.
- `uint16`: Bilangan bulat tak bertanda 16-bit. Rentang nilai dari 0 hingga 65535.
- `int32` (atau rune): Bilangan bulat bertanda 32-bit. Rentang nilai dari -2147483648 hingga 2147483647.
- `uint32`: Bilangan bulat tak bertanda 32-bit. Rentang nilai dari 0 hingga 4294967295.
- `int64`: Bilangan bulat bertanda 64-bit. Rentang nilai dari -9223372036854775808 hingga 9223372036854775807.
- `uint64`: Bilangan bulat tak bertanda 64-bit. Rentang nilai dari 0 hingga 18446744073709551615.

### Floating-point
Sama seperti Integer, float di GO juga memiliki beberapa jenis yang memiliki batasan minimum dan maksimum
- `float32`: Tipe data floating-point 32-bit. Rentang nilai sekitar -3.4E+38 hingga +3.4E+38 dengan presisi sekitar 7 digit desimal.
- `float64`: Tipe data floating-point 64-bit. Rentang nilai sekitar -1.7E+308 hingga +1.7E+308 dengan presisi sekitar 15 digit desimal.

### Boolean 
Seperti dibahasa pemrograman lain, GO juga memiliki tipe data Boolean yang ditandai dengan `bool`, memiliki dua nilai yaitu `true` atau `false`

### Array
Seperti dibahasa pemrograman lain, GO juga memiliki tipe data array yang ditandai dengan []JenisTipeData.
```go
numbers := []int{1, 2, 4, 5, 6, 7, 8, 9, 12}
```

## Slice
Slice menyimpan referensi ke sebuah array dan memiliki panjang (length) serta kapasitas (capacity) yang dapat berubah secara dinamis. Slice digunakan untuk mengelola dan memanipulasi koleksi elemen dalam program Go.
```go
array := [5]int{1, 2, 3, 4, 5}
slice := array[1:4] // Mengambil potongan dari index 1 hingga sebelum index 4
```

### Map
Map adalah struktur data yang merepresentasikan kumpulan pasangan kunci-nilai (key-value pairs).
```go
studentAges := map[string]int{
    "John": 20,
    "Alice": 22,
    "Bob": 21,
}
```

### Struct
Struct adalah tipe data yang digunakan untuk menggabungkan beberapa tipe data lainnya dalam satu kesatuan. Struct memungkinkan pembuatan tipe data baru yang memiliki properti-properti atau fields yang berbeda-beda
```go
// Mendefinisikan struct untuk merepresentasikan seorang siswa
type Student struct {
    Nama    string
    Umur    int
    Kelas   string
}
```

### Pointer
Pointer adalah variabel yang menyimpan alamat memori suatu nilai. Pointer memungkinkan akses langsung ke alamat memori suatu variabel dan memungkinkan manipulasi nilai tersebut melalui alamatnya.
```go
nama := 10
identitas := &nama // Mendapatkan alamat memori dari variabel nama
```

### Function
Function adalah blok kode yang dapat dipanggil untuk melakukan tugas tertentu. Function dapat menerima parameter, melakukan tugas, dan mengembalikan nilai
```go
func add(a, b int8) int8 {
    return a + b
}

func main() {
    result := add(5, 3)
}
```

### Interface
Interface adalah kumpulan definisi metode yang digunakan untuk menggambarkan perilaku suatu objek. mendefinisikan perilaku yang harus dimiliki oleh tipe data lain.
```go
// Definisi interface
type Shape interface {
    area() float64
}
```

### Channel
Channel adalah struktur yang digunakan untuk komunikasi antar goroutine (unit eksekusi ringan dalam Go) yang berjalan secara konkuren. Channel memungkinkan pengiriman data antar goroutine secara aman.
```go
messages := make(chan string) // Membuat channel

// Menjalankan goroutine untuk mengirim pesan ke channel
go func() {
    messages <- "Halo, ini dari goroutine!"
}()
```
### Nil
nil adalah nilai default untuk tipe data referensi seperti pointer, slice, map, channel, dan interface. Ini menunjukkan bahwa variabel atau struktur data tidak merujuk ke alamat memori atau tidak memiliki nilai. Biasanya kita mengenal dengan istilah `Null` untuk bahasa pemrograman lain

### Error
error adalah tipe data bawaan di Go yang digunakan untuk mengindikasikan kegagalan atau kesalahan dalam suatu operasi. Biasanya digunakan untuk memberi tahu bahwa sebuah fungsi tidak dapat menyelesaikan tugasnya dengan sukses dan memberikan informasi terkait kesalahan tersebut.

## Alias Tipe Data
- byte: Alias untuk uint8. Digunakan khusus untuk merepresentasikan data berbasis byte.
- rune: Alias untuk int32. Digunakan khusus untuk merepresentasikan nilai Unicode code point.
- int: Alias untuk tipe data bilangan bulat sesuai dengan arsitektur sistem (32-bit atau 64-bit).
- uint: Alias untuk tipe data bilangan bulat tak bertanda sesuai dengan arsitektur sistem (32-bit atau 64-bit).
- uintptr: Tipe data untuk menyimpan nilai pointer. Ukurannya cukup untuk menyimpan semua pointer dalam sistem.
<hr>

## Variable dan Constanta
Sama seperti bahasa pemrograman lain, GO juga memiliki variable dan constanta dengan cara penamaan yang berbeda dari bahasa lain
```go
// Variable
var nama = "Pandhu"
#atau
nama:= "pandhu"

// Constanta yaitu nilai yang tidak bisa dirubah
const phi float64 = 3.14
```
adapun cara lain untuk mendeklarasikan beberapa variable
```go
var (
	name = "Pandhu"
	age  = 13
)
```

## Print
sama seperti bahasa pemrograman lain, GO juga bisa melakukan output ke terminal dengan beberapa cara
- Pastikan sudah mengimport `fmt` (format)
- Jenis jenis Print:
1. Println: Mencetak argumen ke layar atau ke buffer dengan menambahkan newline (\n) di akhir.
```go
fmt.Println("Halo Dunia")
``` 
2. Print: Mencetak argumen ke layar atau ke buffer tanpa menggunakan format.
```go
fmt.Print("Halo")
```
3. Printf: Digunakan untuk mencetak output dengan format tertentu ke layar atau ke buffer.
```go
fmt.Printf("Nama: %s, Usia: %d\n", name, age)
```
4. Sprintf: Mirip dengan Printf, namun mengembalikan string yang telah diformat sebagai hasilnya.
```go
formattedString := fmt.Sprintf("Nama: %s, Usia: %d", name, age)
```
5. Fprintf: Mencetak output dengan format tertentu ke writer yang ditentukan (file, buffer, dsb.).
```go
// Contoh penulisan ke file
file, err := os.Create("output.txt")
if err != nil {
    panic(err)
}
defer file.Close()
fmt.Fprintf(file, "Nama: %s, Usia: %d\n", name, age)
```
- Jenis jenis format:
    - %d: Format integer dalam representasi desimal.
    - %f: Format floating-point numbers.
    - %s: Format string.
    - %t: Format boolean (true or false).
    - %v: Format nilai apapun dalam representasi defaultnya.
    - %T: Format tipe data dari suatu nilai.
    - %p: Format pointer.

    contoh: 
    ```go
    name:= "Pandhu"
    kelas:= 11
    sekolah:= "SMK Telkom Malang"
    fmt.Printf("nama saya %8s, kelas %3d, sekolah %s ", name, kelas, sekolah)
    ```
    `%`: Tanda bahwa akan ada sebuah format.
    `8`: Menentukan lebar total output menjadi 16 karakter.
    `s`: Menunjukkan bahwa output yang diharapkan adalah string.
    *jadi setiap format yang dipanggil akan sesuai dengan tipe dan urutan nya

## Operasi Matematika
sama seperti bahasa pemrograman lain, GO juga memiliki operasi matematika yang biasa digunakan yaitu:
- `+`: digunakan untuk menjumlahkan
- `-`: digunakan untuk mengurangi
- `*`: digunakan untuk mengkalikan
- `/`: digunakan untuk membagi
- `%`: digunakan untuk modulo/ modulus (sisa pembagian)
- `=`: digunakan untuk melihat hasil
```go
// Operasi matematika
sum := a + b  // Penjumlahan
diff := a - b // Pengurangan
product := a * b // Perkalian
quotient := a / b // Pembagian
remainder := a % b // Modulo (sisa pembagian)
```

adapun cara menggunakan operator penungasan yaitu:
```go
// Operator penugasan
c := 10
c += 5 // c = c + 5 (Operator +=)
``` 
*semua operasi bisa digunakan untuk mengganti `+`

## Operasi perbandingan
sama seperti bahasa pemrograman lain, GO juga memiliki operasi perbandingan yang digunakan untuk membandingkan dua argument dan akan mengembalikan nilai `true` atau `false`
- `>`: digunakan untuk membandingkan apakah argument pertama lebih besar dari argument ke dua
- `<`: digunakan untuk membandingkan apakah argument pertama lebih kecil dari argument ke dua
- `=`: digunakan untuk membandingkan apakah argument pertama sama besar dengan argument ke dua

adapun cara menggunakan operator penungasan yaitu dengan menambahkan `=` contoh:
- `>=`: digunakan untuk membandingkan apakah argument pertama lebih besar atau sama dengan dari argument ke dua

## Percabangan
Di Go, if digunakan untuk eksekusi kondisional dan switch digunakan untuk memilih kondisi yang sesuai dari beberapa opsi yang mungkin.
- if else expression:
```go
x := 10

// Contoh penggunaan if expression
if x > 5 {
    fmt.Println("x lebih besar dari 5")
} else {
    fmt.Println("x tidak lebih besar dari 5")
}
#Atau
// if expression dengan inisialisasi variabel
if y := 20; y > 15 {
fmt.Println("y lebih besar dari 15")
} else {
    fmt.Println("y tidak lebih besar dari 15")
}
```
- switch expression:
```go
day := "Senin"

// Contoh penggunaan switch
switch day {
case "Senin":
    fmt.Println("Hari Senin")
case "Selasa":
    fmt.Println("Hari Selasa")
case "Rabu":
    fmt.Println("Hari Rabu")
default:
    fmt.Println("Hari lainnya")
}
#Atau
// switch tanpa kondisi (mirip if else)
switch {
case day == "Senin":
    fmt.Println("Ini hari Senin")
case day == "Selasa":
    fmt.Println("Ini hari Selasa")
default:
    fmt.Println("Ini bukan hari Senin atau Selasa")
}
```

## Perulangan
Di Go juga terdapat perulangan digunakan untuk melakukan repetisi dalam suatu blok code, supaya mempersingkat kode yang berulang ulang, hanya ada satu perulangan di Go yaitu `for loops`

contoh code perulangan:
```go
counter:= 0
for counter <= 10 {
    fmt.Println("Perulangan ke",counter)
    counter++
}
```
jadi terdapat pengecekan kondisi jika kondisi belum terpenuhi maka akan menjalankan blok code yang ada di dalam sampai kondisi terpenuhi
*pastikan bahwa kondisi akan terpenuhi di dalam perulangan

### Perulangan dengan Statement
Di Go mendukung kita untuk menggunakan statement pada perulangan, dimana terdapat dua statement yang bisa ditambahkan yaitu:
`init statement` merupakan statement sebelum `for` di eksekusi
`post statement` merupakan statement yang akan dieksekusi di akhir tiap perulangan
contoh code:
```go
for counter:= 0; counter <= 10; counter++ {
    fmt.Println("Perulangan ke",counter)
}
```
### Perulangan dengan Range
Di Go mendukung kita untuk melakukan perulangan untuk melakukan iterasi terhadap semua data collection seperti `Array`, `Slice`, dan `Map`
contoh code:
```go
names:= {"pandhu", "arya"}
for index, name:= range names {
    fmt.Println("index",index,"=",name)
}
```
jadi di kode tersebut menggunakan dua statement yaitu index untuk dijadikan key nya dan juga name untuk range dari array names

## Break dan Continue
Di Go ada kata kunci khusus yang digunakan untuk melakukan manipulasi dalam suatu perulangan yaitu:
`break`: digunakan untuk menghentikan perulangan, jadi ketika suatu perulangan menemui kata kunci break maka secara otomatis perulangan tersebut kan berhanti
contoh code:
```go
for i := 0; i < 10; i++ {
    if i == 5{
        break;
    }
    fmt.Println("Perulangan ke",i)
}
```
`continue`: digunakan untuk menghentikan perulangan yan sedang berjalan, dan langsung melanjutkan ke perulangan selanjutnya tanpa perlu mengeksekusi perulangan saat ini
contoh code:
```go
for i := 0; i < 10; i++ {
    if i == 5{
        continue;
    }
    fmt.Println("Perulangan ke",i)
}
```

## Function
Di GO juga memiliki function yang digunakan untuk memisah misah kan code supaya lebih enak dilihat, dan juga bisa digunakan untuk berbagai hal sesuai dengan kebutuhannya

Di GO function ditandai dengan kata kunci: `func` yang berarti function

function dibuat supaya bisa digunakan secara berulang ulang, setelah membuat function kita bisa menjalankan nya dengan cara memanggilnya menggunakan kata kunci nama function diikuti dengan tanda kurung buka dan kurung tutup
```go
func Main(){
    greeting()
}

func greetingO(){
    fmt.Println("Hello World!")
}
```