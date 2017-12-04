---
title: "Bitcoin yang Aku Ketahui"
---

> Bitcoin adalah uang digital.

Aku pikir kalimat itu sudah cukup untuk menerangkan *Bitcoin*. Namun dilihat semakin dalam, ada ada saja istilah asing yang menyangkut kuat dengan Bitcoin, seperti "sistem Bitcoin yang `trusless`" atau "`mining` Bitcoin pakek Hardware" atau "Sistem `blockchain`", dsb...

Bahkan banyak penjelasan sana-sini aku masih belum paham sistem Bitcoin secara keseluruhan, namun aku akan coba untuk penjelaskannya dibawah:

## Konsep Bitcoin

### Bitcoin adalah Uang Digital

Pokoknya Bitcoin itu mirip seperti Email, dan dikirim (transaksi) dengan uang digital (satuan BTC). Bitcoin dalam segi teknis hanyalah sebuah software yang berjalan di komputer yang butuh koneksi internet. 

### Tak ada yang mengatur

Tidak ada organisasi yang mengatur jalannya Bitcoin. Jadi misal anda transfer Bitcoin ke seseorang, maka Bitcoin bisa langsung dikirim utuh tanpa terpotong pajak. Bitcoin ada karena komunitas yang merata diseluruh dunia.

Bitcoin juga memproteksi identitas seseorang. Anda tak bisa melacak orang dari hanya melihat alamat Bitcoin.

Dari segi inilah Bitcoin merupakan ide bisnis pertama kali yang merevolusi Ekonomi. Semua orang bisa bertransaksi tanpa takut dipotong atau terhambar oleh aturan bank atau semacam fenomena krisis lainnya. 

### Bitcoin itu Konsisten dan bisa dipecah sekecil mungkin. 

Bitcoin disimpan dalam satuan BTC dan tak akan berubah meskipun ditransfer bolak-balik ke semua negara. Bitcoin juga dapat dipecah hingga satuan `Satoshi` yang satuannya setara dengan 0.00000001 (seperseratusjuta) BTC.

### Keamanan Bitcoin

Salah satu komponen mengapa Bitcoin bisa hidup adalah karena komunitas dari jutaan `miners` yang aktif merata diseluruh dunia. Para miner tersebut memvalidasi setiap request transaksi (`block`) yang nanti dikumpulkan pada (`blockchain`) untuk diverifikasi kebenaran transaksi dari dua pihak (sistem yang `trustless`) sebelum benar-benar melakukan transaksi.

Itulah gambaran *sangat sederhana* dari proses transaksi Bitcoin, dan setiap transaksi wajib melalui proses yang sama.

### Suplai Bitcoin yang Terbatas

Konsep inilah yang harus digarisbawahi. Suplai Bitcoin memang didesain *terbatas*, dan tak ada yang bisa merubahnya. Pekerjaan miner tidak hanya menyelesaikan transaksi, namun juga mencoba untuk *memanen* Bitcoin *hash* yang baru. Saat ini (Des 2017) ada sekitar `16.2 juta` Bitcoin beredar diseluruh dunia, dan terus bertambah sampai `21 juta` bitcoin yang diprediksi terjadi pada abad 22.

### Berinvestasi di Bitcoin

Nilai konversi antara rupiah dan Bitcoin dapat berubah setiap hari bergantung dari suplai Bitcoin dan Rupiah yang beredar di ekosistem Bitcoin. Inilah mengapa Bitcoin masih sangat jarang ditemui dibeberapa shopping online atau marketplace yang menjual barang-barang yang dipatok secara rupiah.

Fluktuasi di Bitcoin terkadang cukup curam. Misal pada 2015 harga US$ dapat naik 2 kali lipat dalam sehari, lalu kembali turun. Pada konversi 1 BTC = 160 Juta Rupiah (Des 2017), fluktuasi mungkin sering berubah pada level 100 ribuan. Meskipun demikian, statistik mengatakan nilai Bitcoin terus bertambah, dan diprediksi demikian.

Ya, beberapa orang memanfaatkan situasi seperti ini. Contoh kasus ekstrim, seperti Matt yang baru beli Lamborgini seharga 115 US$ hasil invest di Bitcoin sejak tahun pertama Bitcoin diluncurkan.

## Sistem Bitcoin

Bitcoin relatif mudah dipahami pada level konsep diatas. Namun banyak pula misteri yang masih belum kupahami pada bagian sistem atau penerapannya. Aku mungkin tetap mengembangkan bagian ini selama masih mencari informasi. Jika kamu ingin langsung terjun ke Bitcoin, silahkan cek panduan Getting Started di bitcoin.org.

### Alamat dan Wallet Bitcoin

Hal pertama yang dilakukan adalah mendapat akses ke Bitcoin **Wallet** (dompet) dengan mendownload software bitcoin.

Setelah wallet dijalankan, anda bisa *generate* alamat bitcoin pertama. Contoh alamat bitcoin seperti berikut:

[bitcoin:3FkenCiXpSLqD8L79intRNXUgjRoH9sjXa](bitcoin:3FkenCiXpSLqD8L79intRNXUgjRoH9sjXa)

dimana `3FkenCiXpSLqD8L79intRNXUgjRoH9sjXa` adalah alamat Bitcoin yang dituju (berupa hash SHA-256). Perlu diketahui SHA-256 adalah hash unik dan dapat menampung sebanyak 2<sup>64</sup>-1 (18.4467.440.7370.9551.615) hash, sehingga orang bisa berganti alamat berkali-kali untuk menambah keamanan dan melindungi privasi setiap orang didunia.

### Block dan Blockchain

Transfer BTC di Bitcoin tidak bisa dilakukan secara langsung. Jika `A` ingin transfer Bitcoin ke `B`, maka kedua pihak akan berkomunikasi secara otomatis lalu mengirimkan data transaksi (`block`) mereka kedalam `block chain`. Block chain adalah tempat publik (seperti database) dimana seluruh transaksi di kedua pihak dicatat. Setiap `block` dienkripsi menggunakan algorithma dan password tersendiri sehingga tak ada pihak jahat yang bisa mengubahnya. 

Setelah `block` terkirim, miner mulai memvalidasi dan menggabungkan block ke dalam blockchain. Jika penggabungan berhasil ( cek jika transaksi merupakan duplikat, tidak masuk akal, dsb. ) maka miner meng-approve dan melaksanakan transaksi, dan menerima imbalan yang berupa sekian rupiah (dalam BTC) dari yang sudah ditentukan. 

Keseluruhan proses itu dilakukan otomatis tanpa turun tangan oleh kedua pihak. Dan biasanya akan selesai sendiri dalam waktu diatas 10 menit bergantung dari prioritas dan miner yang turun tangan.

### Miner dan Hardware

Siapapun dapat menjadi miner asal mempunyai hardware yang memadai. Untuk menjadi miner, harus join kedalam grup miner (`mining pool`). Saat block dikirim, otak dari mining pool menyebarkan block kesejumlah miner yang berpatisipasi. Ingat bahwa  setiap miner dalam satu pool adalah saling bekerja sama, jadi jika satu transaksi sudah selesai, miner tersebut mengupload hash lalu mengintruksi miner lain untuk menghapus transaksi dalam antrean mereka.

Setiap mining pool mempunyai peraturan tersendiri, termasuk kesempatan waktu mining, waktu online, dsb. Setiap mining pool juga mempunyai support dan forum, jadi jika anda ikut, anda juga otomatis jadi bagian dari komunitasnya.

### Jual dan Beli Bitcoin

Ini adalah istilah untuk mengonversi bitcoin ke rupiah dan sebaliknya.

Bitcoin berjalan berdasarkan komoditasnya sehingga konversi seperti demikian harus dilakukan dengan pihak ketiga. Beberapa individu menjual Bitcoin secara retail di toko online (jadi jangan heran, dan hati-hati).

## Yang Tidak Jelas

Meskipun aku sudah menulisnya panjang lebar. Beberapa pertanyaan dan asumsi berikut masih membuatku jauh dari ingin mempunyai Bitcoin:

+ Bitcoin masih muda (baru ada sejak 8 tahun!) jadi fluktuasi harga masih umum.
+ Memilih agen Jual dan Beli Bitcoin di Indonesia yang diakui kredibilitasnya.

## Penutup

I punya feeling Bitcoin kalau bitcoin bakal mainstream di masa mendatang. Tapi untuk saya, mungkin menunggu dulu hingga bitcoin banyak digunakan di Indonesia.

Untuk jadi miner, aku pun hanya tahu sebatas mata memandang. Mungkin mastah yang berpengalaman bisa curhat ke kolom komen pengalaman-pengalaman selama jadi miner di Indonesia.

```none
* SEKEDAR PERINGATAN *
Bitcoin bukan sekedar mainan, tapi boleh-boleh saja kalau ingin mengerti masalah bisnis online.
Artikel ini bukanlah panduan, tapi boleh-boleh saja untuk sekedar perkenalan tentang Bitcoin.
```

[1]: https://bitcoin.stackexchange.com/questions/114/what-is-a-satoshi