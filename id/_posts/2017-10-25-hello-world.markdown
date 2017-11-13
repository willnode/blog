---
layout: post
title:  "Hello World!"
---

Halo!, artikel ini adalah posting pertama dari blog ini.

Sementara itu, artikel pertama ini  bakal jadi bahan percobaan pake' Jekyll + Kramdown.

<!-- Ini adalah HTML comment, tapi juga bisa kok -->
<!-- Lihat http://wellosoft.net/blog/en/2017/hello-world.html untuk preview -->

## Bahan percobaan
#### Belum lagi
###### lagi ...

***
<br>
Sekarang, kau  mungkin lihat sebuah `inline code` atau _italic_ atau **bold** atau ~~striking~~ texts.

```cs
void MungkinBeberapaKode() { /* coding stuff */ }
```

> Dan beberapa quote yang keren
> Disini juga

| Table | Disini |
|:--|--:|
|Hello|World|
|Bye|Bye|

![Dibuat dengan MSPAINT :D][image]{:title="Dibuat dengan editor legendaris MSPAINT!"}

Mungkin butuh berapa keterangan disini[^1]

* Silver
  + Dibuat
  + Oleh
    - Beberapa
* Bulet
  1. Untuk
  2. Sekarang
     Juga

<details><summary>Contoh</summary>
<code>dari sebuah spoiler</code></details>
<br>
Jaga jaga karena aku [Ngelik halaman eksternal][absolute] atau [Linking this page][relatif] seperti <http://github.com>

Coba dengan --- (dashes) dan ... (dots) dan << arrows >>

[absolute]: https://github.com/
[relative]: {{ "/id/" | relative_url }}
[image]: {{ "/images/2017/madeinmspaint.png" | relative_url }}
[^1]: Hello world

## Gist nya

Aku demen saya Kramdown, karena ia punya banyak fitur tanpa butuh plugin apaun. Okelah ini gist versi english nya bagi yang penasaran:

{% include gist.html gist="willnode/8443d9b88960ae8c810cb79a2b89c9b7" %}