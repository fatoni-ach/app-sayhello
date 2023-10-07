# **Simple say hello appication using go**
### How to run it 
First download modules from `go.mod`
```shell
go get
```
after that run `main.go`
```shell
go run main.go
```
<br>

# **Note**
Catatan ini dibuat selama belajar bahasa pemrograman bahasa Go.
## Unit Test
Unit test merupakan hal yang wajib dibuat ketika membuat aplikasi. Hal ini untuk memastikan aplikasi yang kita buat memiliki kualitas yang bagus, selain itu unit test berguna untuk memberitahu kita ketika ada fungsi yang tersenggol selama develop fitur baru.<br>
Berikut tipe implementasi unit test yang bisa digunakan : <br>
* unit test : merupakan unit test pada umumnya
* sub unit test : merupakan unit test yang memiliki beberapa sub. hal ini bertujuan agar fungsi yang dibuat untuk unit test tidak terlalu banyak sehingga bisa disederhakan menggunakan sub unit test.
* table unit test : merupakan unit test yang menggunakan struct untuk meng inisial parameter dan juga expectednya. Hal ini bertujuan setiap kali membuat scenario baru kita hanya perlu menambahkannya di struct yang sudah dibuat.

Contohnya bisa dilihat di [file ini](helpers/sum_test.go)
<br><br>

## Benchmark
Berbeda dengan unit test yang bertujuan untuk mengetes setiap skenario yang terjadi disebuah fungsi, benchmark ini bertujuan untuk meng benchmark atau menguji performa dari fungsi yang telah dibuat.<br>
Berikut tipe implementasi benchmark yang bisa digunakan : <br>
* benchmark : merupakan benchmark pada umumnya
* sub benchmark : merupakan benchmark yang memiliki beberapa sub. hal ini bertujuan agar fungsi yang dibuat tidak terlalu banyak sehingga bisa disederhakan menggunakan sub benchmark.
* table benchmark : merupakan benchmark yang menggunakan struct untuk meng inisialisasi parameter. Hal ini bertujuan untuk memudahkan dalam membuat skenario baru dimana setiap kali ada penambahan skenario baru kita hanya perlu menambahkannya di struct yang sudah dibuat.


Contohnya bisa dilihat di [file ini](helpers/sum_test.go)
<br><br>