---
layout: post
title: "Kembalikan \"Open with CMD\""
---

Jadi bendingan mana command prompt versus PowerShell? Entah bagaimana  trend masa kini, aku masih suka `cmd` karena ia cepat dan ringan. Aku tidak terlalu suka PowerShell meskipun Microsoft promosiin jauh jauh sampai sekarang di versi baru windows 10.

Pokoknya aku ingin bisa membuka cmd segampang saat aku klik kanan folder apapun yang aku suka. Dulu *pernah* bisa sampai sekarang baru diganti oleh powelshell. Ahh, aku sudah coba google dimana-mana tapi tak nemu juga. Jadi aku karang sendiri….

## Idenya

Kesempatan ku terbuka karena aku punya `vscode`, dan ia asa `Open with Code` setiap aku klik kanan folder apapun, Jadi aku bisa gunakan metode yang sama untuk menambahkan seperti  Open with CMD. Aku cuma perlu muter muter dengan `regedit` dan akhirnya ketemu trik yang manjur.

![image]({{ "/images/2017/openwithcmd.jpg" | relative_url }})

Aku bisa beritahu stepnya tapi lebih mudah sih kalau cuma export. Pokoknya tinggal view raw, download dan run:

{% include gist.html gist="willnode/157a5b95bf894e6556cf02455a8d8cba" %}