---
title: Teknik Hosting Web Gratis
---

Dimanakah para blogger ingin mencurahkan karyanya di internet? Tentu saja, jawabannya ialah kalau tidak di blogspot.com, atau wordpress.com. Apalagi mereka yang berbudget nol, karena kedua platform menyediakan layanan gratis (tidak berbayar), meski fiturnya terbatas dan tidak bebas. Mungkin bagi orang awam, ini adalah hal lumrah. Tapi saat anda mulai ingin mampu mengedit hingga ke-akar-akarnya (entah CSS atau HTML-nya, atau ingin mengubah elemen tertentu dalam tema), disinilah kebebasan anda dirampas. Anda perlu merogoh dompet dalam-dalam sebelum mulai merubahnya.

Tapi jangan khawatir, setelah sekian lama, akhirnya aku menemukan solusinya, perkenalkan [GitHub Pages](pages.github.com).

## GitHub Pages -- Hosting Web Statis

GitHub Pages adalah sebagian fitur yang difasilitasi oleh GitHub. Jadi sebenarnya, GitHub bukanlah layanan Hosting Web, namun ia adalah Git Hosting Service, atau istilahnya, layanan hosting untuk project coding (Git Repository). GitHub ini ialah pilihan Git #1 untuk semua developer (bahkan, aku membackup hampir [semua project ku disana](//github.com/willnode)). GitHub menambah fitur yang mampu meng-host suatu repository Git menjadi sebuah file web yang nantinya bisa diakses disebuah domain seperti wellosoft.net, dan lain sebagainya.

Apa keunggulannya? GitHub memfasilitasi 1 GB storage *per repo*, anda bebas mengedit file HTML/CSS/JS sekreatif yang anda inginkan, dan tak kalah hebat lagi, anda bisa menaruhnya ke domain yang anda telah beli secara gratis. Dari sini, tak ada yang bisa menyaingi kemurahan GitHub Pages. Bahkan, domain ini <https://wellosoft.net> dan semua laman di subdomainnya 100% di hosting oleh GitHub Pages.

Jadi, apa yang kamu tunggu? Waktunya kebebasan kreasimu di kancah web segera dibuka!

## Memulai Repo GitHub Pages pertama anda

> **SEBELUM ANDA MEMULAI:**
> Ingat bahwa GitHub Pages ialah Hosting untuk static files, jadi file server seperti _PHP_ atau _CGI_ tak akan berfungsi. Juga anda tak bisa menyisipkan database, fitur upload atau semacamnya melalui GitHub Pages, anda harus menggunakan layanan pihak ketiga jika anda tetap ingin menggunakannya.

Hal pertama yang anda lakukan ialah mendaftarkan akun GitHub. Buka laman GitHub dan segera mendaftar jika anda belum punya akun. Baca disini jika anda ingin tahu bagaimana caranya.

Setelah itu, anda membuat repo dengan format "*usernameanda*.github.io", misal username saya ialah `willnode`, maka nama reponya harusnya `willnode.github.io`. Mulai dari sini anda telah membuat sebuah repo yang men-hosting `https://willnode.github.io`. Ingat bahwa awalnya harus berupa username anda, atau ia tak akan berhasil.

> Tip tambahan: Barusan anda baru membuat *user pages*. Jika nama repo tersebut selain `github.io`, maka itu adalah *project pages*, dan anda harus mengaktifkan hosting melalui pengaturan pada lalu bagian hampir paling bawah, pilih source ke `master` atau `gh-pages`. Project Pages nantinya dapat diakses melalu subfolder, misal <https://willnode.github.io/blog>

Langkah selajutnya ialah mengisi konten website tersebut.

## Aplikasi yang anda perlukan: IDE, Jekyll dan GitHub Desktop.

Sebenarnya anda bisa langsung mengedit dan mengupload konten website langsung dari Interface GitHub. Namun akan lebih mudah jika anda tahu cara mengedit nya langsung dari komputer anda terlebih dahulu (jika internet anda kurang mendukung, anda bisa meloncati bagian ini). Ada 3 aplikasi yang harus anda install:

#### Web IDE: Visual Studio Code

Hal pertama yang dibutuhkan adalah sebuah Editor yang memadai. Anda boleh memakai IDE yang ada (Sublime, Notepad++) namun aku sarankan anda memasang [Visual Studio Code](https://code.visualstudio.net), karena ia tidak terlalu berat namun mempunyai fitur yang sangat memadai untuk memodifikasi Web.

#### Git Client: GitHub Desktop

Anda butuh Git Client untuk memodifikasi (mengunduh dan mengunggah) konten repositori yang ada. Cara yang paling mudah ialah anda menggunakan [GitHub Desktop](https://desktop.github.com). Anda bisa melihat disini tutorial cepat menggunakan GitHub desktop untuk mengupload konten website [disini](https://pages.github.com/#setup-in-desktop).

#### Webserver: Jekyll

GitHub Pages tidak sekedar hanya menghosting static files (HTML/CSS/JS). GitHub Pages menggunakan Jekyll sebagai *Site Generator*, sehingga anda tak perlu pusing harus mengedit secara manual (HTML/CSS/JS). Anda bisa mengisi konten blog dari bahasa yang lebih sederhana, Markdown (pengganti HTML) dan SCSS (pengganti CSS). Jekyll juga dilengkapi dengan sistem template (Liquid) sehingga anda tak perlu membuat semua komponen dari awal.

Webserver diperlukan sehingga anda bisa langsung melihat pratinjau website tanpa harus menungunggah (push) konten ke GitHub setiap kali anda melakukan perubahan.

Untuk menginstall Jekyll, silahkan anda memasang [Ruby + Devkit](https://rubyinstaller.org/downloads/) (jangan lupa untuk mengeksekusi `ridk install`), lalu menginstall gem yang diperlukan: `gem install bundler github-pages` melalui CMD (pastikan internet terhubung selama instalasi berjalan).

## Mengisi Konten Blog

Setelah semua aplikasi diinstall, anda harus mengunduh (clone) repository anda terlebih dahulu, lalu anda bisa menyisipkan file melalui Explorer `Menu Bar -> Repository -> Show in Explorer` atau langsung mengeditnya menggunakan VS Code  `Menu Bar -> Repository -> Open in Visual Studio Code`. Jangan lupa untuk menyiapkan server Jekyll `localhost` dengan menjalankan CMD  `Menu Bar -> Repository -> Open in Command Prompt` lalu `bundle exec jekyll serve` (atau cukup `jekyll serve` kalau tidak ada `Gemfile`).

Ada 3 cara untuk membangun website melalui GitHub-Pages. Dimulai dari cara termudah hingga cara terumit.

### 1. Cara Manual

Cara ini adalah paling mudah, asalkan anda sudah tahu basicnya tentang HTML/CSS/JS. Jika tidak, saya anjurkan untuk baca artikel sebelumnya.

Sebagai langkah awal, melalui VS Code atau Explorer, anda isi halaman beranda dengan mengisi konten `index.html` dengan template berikut:

```html
<!DOCTYPE html>
<html>
<head>
	<title>Contoh laman</title>
</head>
<body>
	<h1>Halo dunia!</h1>
	<p>Laman beranda ini dihosting oleh <i>GitHub Pages</i>.</p>
</body>
</html>
```

Simpan, lalu coba nyalakan server jekyll dan buka <http://localhost:4000>. Munculkah laman beranda tersebut?

Selain HTML, anda bisa menambah file lain, CSS atau JavaScript, bahkan gambar maupun dokumen kedalam folder repo tersebut. Semua bergantung pada kreasi anda.

#### Menambahkan Markdown

Apa yang lebih baik dari HTML? Markdown tentunya. Markdown cocok digunakan untuk konten blog dan artikel. Contoh seperti `about.md`:

```md
---
title: Tentang Saya # <~ masukkan judul disini
---

Saya sedang belajar *web* dan **programming**.
```

Simpan, lalu buka <http://localhost:4000/about.html>. Munculkah laman about tersebut?

#### Menambahkan SCSS

Bagaimana dengan SCSS? Seperti Markdown, SCSS akan ditanggapi oleh Jekyll jika dilengkapi oleh front matter. Contoh seperti `main.scss`:

```scss
---
---

@import "minima"; // <~ Import SCSS lain dari _sass/minima.scss
```

### 2. Melalui Template Jekyll

Jika anda ingin membuat website dengan sistem blog, anda bisa mulai dari sebuah template. Caranya dengan membuka CMD pada folder repo lalu `jekyll new .` dan tunggu hingga proses instalasi selesai.

Setelah itu anda akan disuguhi dengan file-file berikut

```bash
├─ _posts   # <~ Artikel-artikel masuk sini
│  └─ 20xx_xx_xx-welcome-to-jekyll.markdown # <~ Format artikel (thn-bln-hri-url.md)
├─ _config.yml # <~ Konfigurasi Jekyll
├─ 404.html # <~ Halaman 404 (saat nyasar)
├─ about.md # <~ Halaman '/about'
├─ Gemfile # <~ Gemfile (sejenis daftar depedensi Ruby)
└─ index.md # <~ Halaman '/' (Beranda)
```

Anda bisa menambah artikel dengan menambah file markdown pada `_posts`. Ada pula folder2 spesial lain `_drafts`, `_includes`, `_layouts`, `_posts`, `_data`, `_sass`, `_site` dan mereka sudah dijelaskan di dokumentasi Jekyll [disini](https://jekyllrb.com/docs/structure/).

#### Mengganti Tema

Didalam `_config.yml` terdapat `theme: minima`. Itulah tema yang sedang digunakan. Beberapa daftar tema [ada disini](https://pages.github.com/themes/), anda bisa mengganti `minima` tersebut dengan tema-tema lain sesuai daftar tersebut. Jangan lupa untuk `bundle add <namatema>` sehingga Gemfile bisa terupdate dengan tema tersebut.

Bagaimana jika anda tidak puas? Anda bisa mengganti beberapa file dari tema. Lihat sumber code dengan `bundle show minima` di CMD lalu liat foldernya. Jika perlu copas semua filenya lalu taruh kedalam repo web sehingga anda bisa menghapus `theme: minima` dari  `_config.yml`.

#### Memodifikasi Tema

Ada 2 folder penting disini: `_layouts` dan `_includes`. `_layouts` berisi daftar template sedangkan `_includes` berisi daftar penggalan HTML.

Saat anda menulis HTML/Markdown, anda bisa memilih memakai template mana menggunakan tag `layout: <namalayout>` pada front matter.

Contoh sebuah template sederhana:

```bash
├─ _includes
│  ├─ head.html # <~ Penggalan untuk header
│  └─ foot.html # <~ Penggalan untuk footer
├─ _layouts
│  ├─ default.html # <~ Template untuk interface dasar
│  └─ post.html # <~ Template untuk posting/artikel
├─ _config.yml # <~ Konfigurasi Jekyll
└─ about.md # <~ Halaman '/about'
```

{::options parse_block_html="true" /}
<details><summary markdown='span'>Isi konten dari contoh tersebut:</summary>


Pada `about.md`:

```md
---
layout: post
title: Tentang Saya
---

Halo dunia!
```

Lalu `_layouts/post.html`:

{% raw %}

```html
---
layout: default
---

<div class="main-content">
	<h1>{{page.title}}</h1>
	<span>By: {{site.author}}</span>
	{{content}}
</div>
```

Lalu `_layouts/default.html`:

```html
---
---

<html>
<head>
	<title>{{page.title}}</title>
</head>
<body>
	{% include head.html %}
	{{content}}
	{% include foot.html %}
</body>
</html>
```

{% endraw %}

Lalu `_includes/head.html`:

```html
<header>
	<a href="/">Beranda</a>
</header>
```

Lalu `_includes/foot.html`:


```html
<footer>
	<a href="/about.html">Tentang</a>
</footer>
```

Dan `_config.yml`:

```yml
author: Anonymous
```

Maka hasil jadi nya dari `about.html` ialah:

```html
<html>
<head>
	<title>Tentang Saya</title>
</head>
<body>
	<header>
		<a href="/">Beranda</a>
	</header>
	<div class="main-content">
		<h1>Tentang Saya</h1>
		<span>By: Anonymous</span>
		<p>Halo dunia!</p>
	</div>
	<footer>
		<a href="/about.html">Tentang</a>
	</footer>
</body>
</html>
```

</details>

{::options parse_block_html="false" /}

<div class="sandbox">
<header>
	<a href="/">Beranda</a>
</header>
<div class="main-content">
	<h1>Tentang Saya</h1>
	<span>By: Anonymous</span>
	<p>Halo dunia!</p>
</div>
<footer>
	<a href="/about.html">Tentang</a>
</footer>
</div>


Syntax Liquid tidak cuman berhenti disitu. Terdapat beberapa variabel selain `site` dan `page` yang anda bisa lihat [disini](https://jekyllrb.com/docs/variables/). Selain itu anda perlu belajar tentang [dasar-dasar syntax liquid](https://shopify.github.io/liquid/), beserta Jekyll [Filter](https://jekyllrb.com/docs/liquid/filters/) dan [Tags](https://jekyllrb.com/docs/liquid/tags/) yang tersedia agar bisa membantu desain template lebih lengkap.

### 3. Menggunakan CI (Continous Integration)

Ini adalah metode terakhir yang cukup kompleks. Misalkan, Bagaimana jika aku ingin menggunakan Content Site Generator selain Jekyll? Disinilah CI (Continous Integration) dibutuhkan.

CI digunakan oleh banyak OSS Developers untuk mengotomatiskan pekerjaan mulai dari linting, transpiling, compiling hingga distributing releases. Didalam GitHub Pages, kita dapat menggunakan CI untuk memproses halaman yang ada, dari file-file mentah (Markdown/SCSS/Liquid) menjadi file web (HTML/CSS/JS)

Contoh penggunaan CI adalah laman web utama <https://wellosoft.net>. Laman tersebut menggunakan Framework Vue.JS sehingga aku harus melakukan step-step berikut setiap kali merubah website tersebut:

+ `npm run build` (dari sini folder `/dist` akan berisi file web yang siap hihosting)
+ Commit dan push project ke branch `source`
+ memindahkan folder `/dist` keluar dari git repo
+ `git checkout master` untuk pindah ke branch master
+ mengembalikan folder `/dist` tadi kedalam git
+ Commit dan push project ke branch `master`
+ `git checkout source` untuk kembali ke pindah branch source lagi

Ribet kan? Iya. Makanya CI diperlukan untuk mengotomatiskan semua steps diatas. CI Rekomendasi saya ia <https://travis-ci.org>. Travis CI itu 100% Gratis dan tidak batasi untuk repo publik di GitHub. Untuk mengaturnya aku hanya perlu menyalakan CI pada repo tersebut dan mengisi beberapa config file, dan tada! Setiap push ke source akan mengeksekusi langkah2 diatas dalam waktu beberapa menit saja.

## Kesimpulan


