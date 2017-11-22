---
title: "Jadikan Windows Hebat Lagi"
---

Saya baru saja mengupdate windows ke versi terbaru `1709`. Sayangnya beberapa hack untuk membuat komputer saya jadi hilang karena sudah hilang terganti.

Jadi, saya membuat batch file untuk itu.

Batch file akan melakukan hal ini:

+ Membuat restore point (opsional)
+ Menonaktifkan layanan tugas berat
+ Nonaktifkan startup umum oleh Task Scheduler
+ Menonaktifkan Cortana
+ Hapus data telemetri
+ Beberapa privasi dan speed tweak
+ Menghapus aplikasi windows store (termasuk edge)
+ Uninstall onedrive
+ Restart explorer

Saat Anda menjalankan skrip ini dengan akses admin, komputer Anda akan berubah menjadi sangat ringan dan cepat.

Ini mungkin menghapus beberapa bagian yang penting bagi Anda. Tapi bagi saya, ini bekerja dengan cukup baik. Dengan 2 GB RAM windows hanya mengisi memori 40% tanpa lonjakan CPU atau disk saat startup. Sangat efisien!

Download dan jalankan skrip di bawah ini dengan akses admin:

> Hati-hati dengan apa yang Anda lakukan, bahkan kalau file ini tidak akan menggigit :)

{% include gist.html gist="willnode/056defa0107a5998426f238699500d88" %}
