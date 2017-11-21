---
title: "Berhijrah ke Situs Statis (dan alasannya)"
---

Jadi aku sudah membebaskan diriku sendiri dari siksa `wordpress`. Tahun-tahun lalu, Aku terbiasa untuk posting semua hal disana. Dulu awalnya OK, lalu masalah muncul...

```
Artikel ini tidak pernah dimaksudkan untuk mengeritik wordpress.
```

Alwalnya aku menggunakan `wordpress.com` yang gratis karena aku tak mampu membayar `4$` per bulan. Namun ini berarti aku harus membiasakan dengan terbatasnya tema, iklan di beranda dan sistem artikel yang buruk. Perhatikan kata sistem artikel yang aku tulis disini, karena itu membuatku menciptakan dua blog yang berbeda `wellosoft.wordpress.com` dan `willandgottaloveideas.wordpress.com` hanya untuk bisa mempunyai blog multi-bahasa, dan Itu nggak bagus.

## Diselamatkan oleh GitHub

Awal tahun ini, Aku mulai aktif di GitHub saat [GitHub Desktop][gitdesktop] sudah muncul (mengapa? karena github push via CLI memang rumit). Hal itu merubah kebiasaanku sangat sampai sampai aku selalu `git init` sebelum memulai project baru.

## Masalah dengan GitHub Pages

Kemampuan GitHub untuk membuat website statis sendiri menarik perhatianku seketika itu juga, namun *dulu* memang susah:

+ Semua wajib dilakukan dengan CLI (GitHub Desktop belum muncul).
+ Harus meletakkan page di branch `master` atau orphan `gh-pages` yang emang agak rumit.
+ **Aku gagal menginstal Jekyll**. Sayang instalasi Jekyll memang rumit untuk windows.

Dengan tantangan tersebut aku masih bisa untuk membuat gh-pages pertama untuk project saya [ComCalc][comcalc]. Dulu memang greget, jadi aku tidak ingin mencoba lagi untuk beberapa bulan.

## Lalu Datanglah Cahaya

Aku tak ingat kapan aku pertama kali menggunakan `GitHub Desktop` namun bulan lalu, Aku ingin membuat website dokumentasi untuk [Engine4][engine4]. Yang membuatku kagum adalah [DocFX][docfx] adalah software yang tepat aku gunakan. Setelah beberapa hari akhirnya aku bisa mempostingnya online dan itu membuatku sangat senang. Aku suka bagaimana mudahnya menginstall `DocFX` menggunakan [Chocolatey][choco] daripada with [Jekyll][jekyll].

> Lalu semua telah berubah

Aku tidak percaya bagaimana cepat semuanya berubah. Pada waktu itu, GitHub baru support untuk hosting pages via folder `/docs`. Itu memotivasi saya untuk membuat website lagi untuk hosting daftar project saya, lagi menggunakan DocFX. Aku tempatkan hasil HTML kedalam `/docs` lalu membuatnya online [disini][expertise].

Jadi aku sudah punya dua website dalam satu bulan, dan itu secara dramatis sudah menambah pemahamanku dalam desain web.

## Hijrah

> Aku ingin websiteku terlihat professional. Bagaimana pun aku mendesainnya website ku selalu terlihat main-main sampai aku miliki sebuah domain.

Masalah muncul disini. Wordpress tidak memperbolehkanku untuk menambah domain sampai aku berlangganan, dan wordpress hosting masih agak mahal. Namun, kebebasanku dimulai saat aku mendengar GitHub pages bisa diletakkan ke domain sendiri, dengan cuma-cuma.

Ini menuntunku untuk mencoba Jekyll lagi. Namun sekali lagi, Jekyll pun sudah berubah. Sekarang aku bisa menginstall dengan `bash` atau `choco`. Butuh beberapa jam sampai selesai menginstal jekyll di laptop (Aku install via `choco` meski masih agak kesulitan tapi ok sajalah).

Setelah beberapa hari, Aku buat website untuk halaman beranda dan `/blog/` dua-duanya pakai Jekyll. Memang indah bahwa that dua software web generator (Jekyll dan DocFX) menggunakan [Liquid][liquid] jadi aku tidak perlu khawatir soal manual copas dan bisa fleksibel dengan tema.

Sepertinya semua website ku sekarang pakai website statis, kan? Oleh karena itu ini pertama kali juga aku **membeli domain** `wellosoft.net` dan meletakkan `CNAME` ke `willnode.github.io`. Dan lahirlah Website professional saya [website][mysite] sendiri :)

## Kesimpulan

Aku merasa bersyukur melihat skenario sedemikian bisa dilakukan di masa kini. Semua yang kubutuhkan hanyalah sebuah domain seharga `10$` atau `Rp 150 rb` pertahun yang tentu pas banget buat budget kecil seperti saya, untuk yang lain (hosting) biar diselesaikan dengan GitHub.

Emang super. GitHub menyelesaikan dilema saya di era serba online. Ia seperti Facebook-nya programmer masakini. Aku benar-benar menghargai bahwa situs dan pelayanan mereka itu ada untuk setiap orang, secara gratis.

Aku sebenarnya khawatir kalau situs ini di banned kalau sudah melewati [100 GB bandwidth limit][usagelimit] (dirata-rata sekitar 400 penguntung setiap jam). Namun karena saya emang amatir, jadi gak perlu takut yang aneh aneh =)

[gitdesktop]: https://desktop.github.com
[jekyll]: https://jekyllrb.com
[choco]: https://chocolatey.org
[docfx]: https://dotnet.github.io/docfx/
[comcalc]: https://willnode.github.io/ComCalc/index.html
[expertise]: https://willnode.github.io/expertise/
[engine4]: http://wellosoft.net/engine4-doc/
[mysite]: http://wellosoft.net
[usagelimit]: https://help.github.com/articles/what-is-github-pages/#usage-limits
[liquid]: https://shopify.github.io/liquid/