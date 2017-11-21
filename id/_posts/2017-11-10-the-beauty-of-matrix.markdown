---
title: "Keindahan dari Matriks untuk Rotasi"
---

Matrik bisa disimbolkan sebagai rotasi sebuah objek. Jika kamu kesini karena kesulitan memahaminya saat memprogram game, maka lanjutkan. Aku punya sakit yang sama sampai aku kesini untuk berbagi dengan anda hari ini.

## Sebab Sebab

Aku saat ini sedang mengembangkan [Engine4][engine4], game engine 4D. Tantangan tersulit saat membangun engine tersebut itu adalah bagaimana aku harus memanipulasi rotasi, yang tentu aku simpan dengan rotasi matrik. Anda mungkin melihat artikel ini dikhususkan untuk area 3D atau 2D, namun sebenarnya bisa diterapkan di area dimensi apapun.

### Kegunaan

Graphic library seperti [OpenGL][opengl] dan [DirectX][directx] menggunakan matrik untuk memanipulasi rotasi. Jadi sebaiknya ikuti apa yang mereka gunakan. Selain itu untuk 3D atau dimensi lebih tinggi, Aku mending memakai matrix daripada rotasi euler untuk menjauhi masalah dengan Gimbal Lock.

Anda mungkin pernah mendengar tentang [Quaternion][quaternion] juga, ia super efisien namun sulit untuk dipahami dan hanya ada di area 3D. Jadi aku tak membahasnya disini. Yang penting, mari kita lihat  lihat yang mana yang ingin kamu gunakan setelah membaca artikel ini.

### Tujuan

Artikel ini ditulis sehingga kamu tidak kebingungan dengan matrik, sehingga kamu bisa menggunakan sifat analitisnya dengan indah. Aku anggap kamu sudah tahu apa dan bagaimana penggunaan dasar dari rotasi matrik dan relasinya dengan transformasi objek termasuk istilah seperti area [*lokal* dan *global*][localglobal].

Cukuplah untuk intro, sekarang mari kita bahas...

## Sifat dari Rotasi Matrik

Apa yang membuat sebuah matrik bisa dianggap sebagai rotasi matrik?

### Jumlah Ordo yang sama

Dasarnya matrix harus berbentuk persegi jadi ia tidak akan mengganti jumlah dimensi sebuah vektor.

```
2D matrix:      3D matrix:        4D matrix:
| 1 0 |         | 1 0 0 |         | 1 0 0 0 |
| 0 1 |         | 0 1 0 |         | 0 1 0 0 |
                | 0 0 1 |         | 0 0 1 0 |
                                  | 0 0 0 1 |
```

### Determinan = 1

Sebuah rotasi matrik harus selalu punya nilai determinan 1, jadi ia tidak akan merubah penjang vector apapun. Jenis matrik seperti ini juga disebut sebagai [matrik ortogonal][orthomatrix], dan itu pertanda baik, yang akan kita diskusikan nanti.

Saat menulis sebuah program yang menggunakan rotasi matrik, adalah ide yang bagus untuk menulis beberapa tes determinan untuk memastikan operasi matrik and berjalan dengan baik.

> **Penting**: Di [wikipedia][orthodeter] matrik orthogonal juga bisa mempunyai determinant -1, namun aku tak tahu jika itu juga relevan untuk rotasi matrik pula.

### Tak bisa ditambah

Matrik dengan ordo yang sama bisa ditambah atau dikurangkan, namun konsep ini tidak lagi relevan untuk matrik ortogonal karena operasi itu tentu akan merubah determinannya meskipun kedua matrik sama sama ortogonal awalnya.

Yang jelas kamu seharusnya tidak perlu menulis pertambahan atau pengurangan matrix jika hanya dimaksudkan untuk rotasi untuk menghindari skrip nakal merusak software anda.

## Hal yang Menarik

Jadi apa yang menarik dari rotasi matrik?

### Menggabung Rotasi

Ini hal yang dasar. Jika anda sudah mendengan kuaternion maka perkalian rotasi matrik itu sama sifatnya seperti perkalian dari Kuaternion. Anda bisa menggabungkan dua rotasi dari mengkalikannya.

```
C = A X B
```

Anda juga bisa [meng-invers][minverse] salah satu matrik. Contohnya, membuat matrik yang membuat sebuah poin yang relative ke objek B menjadi relatif ke A.

```
C = inverse(A) * B
```

Pelu diingat kalau perkalian matrix tidak komutatif. Kita akan diskusikan nanti.

### Transpos = Invers

Ini favoritku. Jika kamu bisa pastikan bahwa skripmu tidak gagal saat mengetes determinan matrik, maka kamu bisa mengganti semua skrip invers dengan fungsi [transpos][transpose] yang jauh lebih enteng karena keduanya adalah operasi yang identik khusus matrik yang orthogonal.

```
Matrik Searah jam:                    Matrix tidak searah jam
|  cos x   sin x | <- Ditranspos -> |  cos x  -sin x |
| -sin x   cos x | <- Ditranspos -> |  sin x   cos x |
```

Jika anda belum pernah mendengar yang satu itu, maka sebaiknya terapkan konsep ini karena ini termasuk peluang besar untuk mengoptimasi skrip.

### Kolom untuk Vektor Arah

Adalah hal yang wajar jika kita ingin mendapatkan vektor orientasi (normal) dari sebuah objek. Biasanya anda bisa mendapatkannya dengan mengkalikannya dengan rotasi objek tersebut.

```
Hasil:  Rotasi:     Vektor:
| 0 |   | 1  0 0 |   | 0 |
| 1 | = | 0  0 1 | X | 0 |
| 0 |   | 0 -1 0 |   | 1 |
  ^                    ^
  |                    +--- Area Local
  +------------------------ Area Global
```

Operasi seperti diatas bisa dioptimalkan. Anda hanya butuh mengambil **kolom ke-n** (bukan baris) dari matrik rotasi tergantung dari vektor. Lihat kembali kolom ketiga dari contoh diatas.

```
Result: Rotation:    Vector:
| 0 |   | 1  0 0 |   | 0 |
| 1 | = | 0  0 1 | X | 0 |
| 0 |   | 0 -1 0 |   | 1 |
  ^            ^
  +------------+----------- Area global (sama persis)
```

Ingat jika anda justru mengambil dari baris ke-n, maka itu sama dengan mengkalikan vektor dengan matrik yang ditranpos.

### Gabung Rotasi dengan Posisi

Ini adalah fitur yang hanya ada jika menggunakan rotasi matrik. Kita semua tahu bahwa sebuah objek tidak hanya mempunyai sebuah rotasi, namun juga posisi mereka sendiri. Game designer punya solusi lebih baik dengan cara menyimpan data lokasi kedalam matrik dibagian kolom berikutnya:

```
|  1  0  0 10  |
|  0  0  1  1  |
|  0 -1  0  2  |
   <--+-->  ^
      |     +--- Posisi
      +--------- Rotasi
```

Lalu kita bisa mengkalinya dengan menambahkan ekstra satu baris kedalam perkalian. Cara ini bisa digunakan untuk `matrix * vector` dan `matrix * matrix`.

```
| 10 |   |  1  0  0 10 |   | 1 |
|  1 | = |  0  0  1  1 | X | 2 |
|  0 |   |  0 -1  0  2 |   | 0 |
|  1 |   |  0  0  0  1 |   | 1 |  <---  Ekstra baris
```

## Bagian yang Menantang

Adakah faktor yang mungkin membuatmu tidak bisa menggunakan konsep ini? Kamu tidak sendirian, Aku sudah menangkap semua yang ganjal untukmu ...

### Penyimpanan (Major Ordo)

Ya, ini adalah masalah #1 untuk kebanyakan developer. Ini terjadi karena memori (RAM) tidak disimpan sebagai array 2D. Ia harus direlokasi sebagai array 1D terlebih dahulu.

Ini mungkin terdengar sederhana, namun sebenarnya tidak karena ada dua cara untuk melakukannya:

```
Row major:  Column major:

| 0 1 2 |   | 0 3 6 |
| 3 4 5 |   | 1 4 7 |
| 6 7 8 |   | 2 5 8 |

*Numbers denote array index
```

Lebih buruk lagi, OpenGL merelokasi matrix sebagai major kolom **sedangkan DirectX adalah [kebalikannya][matrixlayout]**.

Mengapa ini penting? Karena memilih baris/kolom secara tak tentu bisa memicu *skrip nakal*, lebih parah lagi, perkalian matrik banyak bergantung dari itu:

```c#
// matrix * vector paling sederhana
vector operator *(matrix, vector){
    return vector {
        x = dot(matrix.row[0], vector)
        y = dot(matrix.row[1], vector)
        z = dot(matrix.row[2], vector)
    }
}
```

```c#
// matrix * matrix paling sederhana
matrix operator *(matrix1, matrix2){
    return matrix {
        column[0] = matrix1 * matrix2.column[0]
        column[1] = matrix1 * matrix2.column[1]
        column[2] = matrix1 * matrix2.column[2]
    }
}
```

Saat membaca skrip seseorang, pastikan untuk tahu apakah kode matrik mereka didesain secara baris berbaris atau kolom berkolom. Ini penting karena kebanyakan orang menulis skrip matrik mereka kedalam veriabel vektor seperti `matrix.ex` sedangkan anda tidak tahu apakah `ex` adalah baris atau kolom ke-n dari `matrix`. Ini juga berlaku saat matrik diakses menggunakan array berangkap seperti `matrix[column][row]` atau `matrix[row][column]`.

Saat kamu menggunakan skrip pihak ketiga dan tahu bahwa ada major matrik yang konflik, atau ingin menterjemah OpenGL ke DirectX dan sebaliknya, pastikan untuk `transpose()` matrik itu dan semua akan OK.

> Saat menulis kode matrik sendiri, Saran saya adalah gunakan *major berbaris* karena sudah naluri kita untuk membaca kata baris berbaris, dan konsisten dengan notasi matematika saat membaca array berangkap.

### Konversi Euler

Ini biasanya tidak jadi masalah besar. Kecuali jika kamu menggunakan kode pihak ketiga dan kamu tahu bahwa pasti disitu ada konversi euler yang terlibat.

Singkatnya, inilah caranya rotasi euler 3D dikonversi menjadi matrik:

```
| 1 1      0     |   | cos y 0 -sin y |   | cos z -sin z 0 |
| 0 cos x -sin x | X | 0     1  0     | X | sin z  cos z 0 |
| 0 sin x  cos x |   | sin y 0  cos y |   | 0      0     1 |
```

Lihat bagaimana mudahnya urutan perkaliannya bisa dirubah. Dengan susunan diatas kita bisa menyebutnya Konversi Euler ZYX (axis Z lalu Y lalu X). Ini hanya salah satu dari [banyaknya susunan][eulconvent] yang ada. Lagipula, konversi euler ke matrix hanyalah permulaan, mengkonversi kembali itu yang susah (karena butuh semacam strategi menebak). Oh ya, apakah aku sudah mengatakan CW vs. CCW juga?

Tidak ada bakuan untuk konversi euler (mungkin kecuali [di beberapa profersi][eulformal]). Tiap orang bebas memilih urutan mana yang terbaik, selama mereka konsisten.

> Untuk mencegah masalah seperti ini dengan pihak ketiga, cobalah untuk hanya menggunakan konversi euler yang disediakan oleh mereka (cocok seperti game engine). Atau lebih baik, hindari mempassing euler bolak balik, langsung gunakan matriknya saja (dengan memperhatikan major order).

### Urutan perkalian

Memang mudah untuk membuat (namun sulit untuk melihat) kesalahan saat  mengalikan matrik karena sifat ketidak komutatifnya.

Sebagai awalan, semua orang tahu bagaimana caranya mengkonversi arah vektor dari area lokal ke global (dan sebaliknya). Karena `matrix * vector` itu diperbolehkan secara matematika, namun tidak untuk `vector * matrix`.

```js
global = rotation * local;            // lokal ke global
local = transpose(rotation) * global; // global ke lokal
localInB = transpose(rotationB) * rotationA * localInA; // lokal dari A menjadi lokal dari B
```

Namun masalah muncul ketika *memanipulasi matrik*, khusunya saat mengintegrasi delta rotasi atau kecepatan angular.

Di dunia matrix, anda menukar keceparan angular menjadi delta rotasi dengan cara ini:

```c#
delta = fromeuler(angularvelocity); // Sip. tinggal konversi ke euler
```

Saat mengintegrasi delta ke rotasi, biasanya untuk orang tak peduli dengan urutannya akan menulis seperti ini:

```c#
rotation *= delta; // jalan pintas untuk "rotation = rotation * delta"
```

Itu **tidak** sepenuhnya benar. Walaupun contoh diatas ok ok saja untuk rotasi sederhana (contoh putar meriam, karakter kesamping), namun ia akan rusak untuk situasi yang kompleks (seperti physics simulation) karena orientasi dari delta tersebut akan *terganti* oleh matrix sebelumnya.

Pilihan anda adalah membalikkan perkaliannya.

```js
rotation = delta * rotation; // Ini urutan yang benar. Dapat dikata-kata menjadi "ganti rotasi berdasarkan delta" dan bukan sebaliknaya
```

OK. Tapi dalam beberapa situasi situasi, saat kamu ingin delta *yang orientasinya relatif* terhadap sebuah objek, yang harus kamu lakukan adalah membelit delta tersebut dengan matrik dari objek. Beberapa orang menyebut ini sebagai *operasi sandwich*

```js
rotation = transpose(object) * delta * object * rotation; // Operasi ajaib dimana "Aku ingin mengadop delta berdasarkan objek tapi tak ingin membuatnya merubah rotasi"
```

## Semua Tercover

Itu adalah semua yang aku tahu tentang matrik sebagai orientasi objek. Ada pertanyaan letakkan saja di komen bawah.

[engine4]: http://wellosoft.net/engine4-doc/
[directx]: https://en.wikipedia.org/wiki/DirectX
[opengl]: https://en.wikipedia.org/wiki/OpenGL
[quaternion]: https://en.wikipedia.org/wiki/Quaternion
[localglobal]: https://teamtreehouse.com/community/global-space-vs-local-space
[orthomatrix]: https://en.wikipedia.org/wiki/Orthogonal_matrix
[orthodeter]: https://en.wikipedia.org/wiki/Orthogonal_matrix#Matrix_properties
[minverse]: https://www.mathsisfun.com/algebra/matrix-inverse.html
[transpose]: https://en.wikipedia.org/wiki/Transpose
[matrixlayout]: http://www.mindcontrol.org/~hplus/graphics/matrix-layout.html
[eulconvent]: https://ntrs.nasa.gov/search.jsp?R=19770019231
[eulformal]: https://en.wikipedia.org/wiki/Rotation_formalisms_in_three_dimensions#Euler_rotations