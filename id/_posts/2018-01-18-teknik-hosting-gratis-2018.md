---
title: Teknik Dapatkan Hosting Gratis 2018
---

Punya Hosting Sendiri secara gratis mungkin seperti mimpi. Tapi tidak sekarang. Buktinya aku punya website [https://wellosoft.net][] sudah lengkap termasuk semua fitur hosting kelas Enterpise (Subdomain + Email Forward + SSL + Blog Comment + Analytics + CDN + Node Server), semua secara **cuma-cuma** (aku hanya perlu bayar domain per tahun).

Kuncinya adalah nggak rely sama satu provider. Untuk wellosoft.net aku gunakan:

## Hosting

+ Static Hosting: [https://pages.github.com][]

Buat repo baru, tambahkan `index.html`, push, lalu set page source ke master. Selesai! Lebih gampang dan fleksibel daripada Wordpress dengan temanya yang super terbatas.

Kalau nggak percaya, coba lihat betapa minimalisnya [tampilan beranda saya][https://wellsoft.net] dan enaknya [cara ku membuat website dengan frameworks tanpa template][https://github.com/willnode/willnode.github.io].

Oh ya, Aku harusnya bertemi kasih pada GitHub karena aku diberikan bandwidth 100 GB secara cuma-cuma. Plus aku bisa memforward ke custom domain secara gratis!

+ Image Hosting: [https://cloudinary.com][]

GitHub membatasi max 1 GB per repo, jadi untuk file besar (seperti gambar) harusnya beralih ke CDN lainnya. Untungnya ada Cloudinary yang menggratiskan bandwidth gambar hingga 20 GB! Yu Hu!

+ Video Hosting: [https://youtube.com][]

Iya iyalah sekalian buat channel biar bisa populer ;)

+ Document / Zip Hosting: [https://drive.google.com][]

Untuk upload file besar atau PDF ini yang paling mudah. Free Storage 15 GB sudah sangat leluasa itu.

## Fitur Server Lain

+ Blog Comment: [https://disqus.com][]

Karena GitHub hanya bisa digunakan buat static hosting, aku butuh platform lain untuk blog. Untungnya lagi ada Disqus. Disqus memang sudah lama jadi platform comment yang paling populer!

+ Analytics: [https://analytics.google.com][]

Enaknya bisa control HTML sendiri adalah aku bisa mentag dengan Google Analytics, jadi aku bisa melihat berkembangan views dan traffic website, Gratis! Kelas Google!

+ SEO: [https://www.google.com/webmasters/][]

Haha, buat yang maniac SEO jangan khawatir yah google punya banyak webmaster tools jadi aku bisa pantau kualitas SEO langsung dari Google!

+ CDN + SSL: [https://cloudflare.com][]

Kata siapa SSL (HTTPS) tidak gratis? Pakai Cloudflare, Aku bisa ganti nameserver (proxy) domain sendiri melewati Cloudflare, yang sudah menyiapkan SSL + Cache Control + DDoS Protection gratis! Bayangin tuh!

+ Email Forwarding: [https://mailgun.com][]

Jika aku ingin menggunakan email atas nama domain, aku butuh Email Forwarder. Untungnya saja ada mailgun, yang menyedian 10.000 email forward gratis per bulan! Kebayang nggak sih bagaimana leganya punya `willnode@wellosoft.net` secara gratis!

+ Server Deployment: [https://zeit.co/now][]

Karena GitHub hanya menyediakan static hosting, Aku tidak bisa mengeksekusi (deploy) server sendiri seperi PHP atau Node.JS.

Tapi dengan now.sh aku bisa! Dengan bandwidth 1 GB, aku bisa deploy aplikasi Docker (Bahasa Go) atau Node.JS (Bahasa JS) dan mengeksekusinya sebagai server! (Untuk PHP sori ya tidak disediakan).

untuk wellosoft.net kugunakan now.sh untuk crawl [list repository GitHub](https://wellosoft.net/#repos) saya sendiri setiap hari sekali. Jadi ia selalu up to date tanpa aku turun tangan 😁.

BONUS: di beranda wellosoft.net coba intercept network traffic ke now.sh menggunakan Inspector dan coba lihat gimana aku menyiapkan servernya!

## Penutup

Hah kan! Semua nya gratis! Aku hanya perlu membayar domain yang untuk saya sekitar 150 ribu per tahun. Buat agan-agan yang ingin rekomendasi tempat beli domain aku sarankan [https://niagahoster.co.id][] karena aku langganan disitu, dan mereka melayani bayar domain secara cash melalui Indomaret 🎉.

Sekian!