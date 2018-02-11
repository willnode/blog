---
title: Test HTML OK
---
# Website Dasar dengan HTML

> HTML: Seperti Microsoft Word, pakek Notepad.

## Persiapan Pertama

Untuk kelengkapan, anda bisa install VS Code biar lebih gampang.

1. Buat folder kosong, nama terserah.
2. Buat file kosong `index.html`.
3. Coba dibuka file pakai Google Chrome.
4. Kosong kan? Mari kita isi.

## Menulis HTML

HTML paling minimal harus diisi:

```html
<html>
    <head>
        <!-- Tulis head disini -->
    </head>
    <body>
        <!-- Tulis body disini -->
    </body>
</html>
```

HTML tak peduli akan identasi atau kelebihan spasi disetiap tag `<>`.

## Head

Head itu berisi metadata, alias identitas halaman. Contoh paling popular itu judul dan ikon.

Judul isi dengan tag `title`:

```html
<title>Judul disini</title>
```

Favicon (ikon) diisi dengan `meta icon`:

```html
<link rel="icon" type="image/png" href="favicon.png" />
```

Lalu letakkan gambar bernama `favicon.png` disesama folder.

Simpan dan buka lagi `index.html` di Chrome. Lihat perubahan di Judul Tab Browser?

## Body

Body inilah dimana kamu mengisi konten internet.

Cara kerjanya sama seperti notepad. Tinggal ketik langsung isinya. Namun jika ingin ditambahkan gaya, perlu ada tag khusus. Contoh:

```html
Mawar itu <i>merah</i>, Laut itu <b>Biru</b>.
```

Menjadi: Mawar itu <i>merah</i>, Laut itu <b>Biru</b>.

Daftar tag styling:

| Tag | Arti + Contoh |
|---|---|
| `<b>` | <b>Tebal</b> |
| `<i>` | <i>Miring</i> |
| `<s>` | <s>Garis</s> |

Untuk judul atau paragraf:

| Tag | Arti + Contoh |
|---|---|
| `<p>` | <p>Paragraf</p> |
| `<pre>` | <pre>Paragraf (kode)</pre> |
| `<h1>` | <h1>Judul</h1> |
| `<h2>` | <h2>Subjudul</h2> |
| `<h3>` | <h3>Subsubjudul</h3> |

Untuk pembatas vertikal:

| `<br>` | Baris Baru (Tidak usah ditutup) |
| `<hr>` | Baris Pembatas (Tidak usah ditutup) |

Tambahan yang wajib diketahui:

| `<ul>` | Baris 1.2.3.4 ... Didalam `<ul>` diisi dengan `<li>` yang berisi poin |
| `<ol>` | Baris bulet-bulat ... Didalam `<ol>` diisi dengan `<li>` yang berisi poin |
| `<img src="...">` | Gambar (Tidak usah ditutup) (`...` diisi dengan alamat gambar) |

## Siap?

PR Anda: [Buat HTML seperti ini (klik disini)](html-test-example.html). Lebih persis Lebih baik. Copas teks boleh.

Selesai langsung dikumpulkan. Pertanyaan tinggal kirim ke `willnode@wellosoft.net` (saya sendiri).

Enjoy 😸!