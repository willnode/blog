---
title: Portal Belajar Web untuk Pemula
---

Membuat website adalah suatu hal yang sangat asik dilakukan. Menurutku, setiap orang yang ingin belajar programming, dia harusnya mencicipi untuk membuat web terlebih dahulu. Namun tahukah kamu? Membangun sebuah website tidak hanya cukup dengan belajar bahasanya saja, namun perlu pengetahuan tentang menyiapkan sebuah server, sebuah database, yang tentu membuat bingung pemula yang ingin belajar murni tentang Web.

Jadi bagaimana caranya? Perkenalkan [Codepen.io](https://codepen.io/).

Codepen.io ialah situs untuk menulis dan berbagi front-end web code. Ia bisa jadi sarana latihan, bisa jadi sebuah koleksi, atau bisa jadi untuk demonstrasi web. Cara penggunaannya itu sederhana, anda tinggal [mendaftar](https://codepen.io/accounts/signup/user/free) lalu membuat sebuah [Pen baru](https://codepen.io/pen). Setiap pen anda akan disimpan didalam [dashboard](https://codepen.io/dashboard/) anda dan muncul dalam profil anda.

Anda bisa melihat fitur-fitur Codepen [disini](https://codepen.io/hello/)

![](https://blog.codepen.io/wp-content/uploads/2014/04/autoupdate.jpg)

Dari situ, anda disuguhkan tiga panel, yakni HTML, CSS, JS.

<a style="font-size:2rem" href="//codepen.io/pen">-- Coba Codepen Sekarang --</a>

### HTML

`HTML` adalah kerangka web. HTML dapat ditulis layaknya tulisan biasa yang diapit oleh [HTML Tags](https://www.w3schools.com/tags/).

![](https://www.computerhope.com/jargon/h/html-tag.gif)

Contoh beberapa elemen HTML:

<iframe width="100%" height="400" src="//jsfiddle.net/willnode/wxgz5v6r/embedded/html,result/" frameborder="0"></iframe>

Resources:

+ [W3School HTML Tutorial](https://www.w3schools.com/html/)
+ [W3School HTML Tags](https://www.w3schools.com/tags/)
+ [HTML Cheatsheet (Interaktif)](https://htmlcheatsheet.com/)
+ [HTML Cheatsheet (PNG/PDF)](https://websitesetup.org/html5-cheat-sheet/)
+ [HEAD Tag Cheatsheet](https://github.com/joshbuchea/HEAD)

Konsep dan Istilah dasar:

+ [Struktur HTML](https://bethsoderberg.com/blog/html-basics-elements-tags-and-document-structure/)
+ [Inline vs Block](https://codeburst.io/block-level-and-inline-elements-the-difference-between-div-and-span-2f8502c1f95b)
+ [HTML Form](https://www.w3schools.com/html/html_forms.asp)
+ [HTML5 + Semantic Tags](https://webflow.com/blog/html5-semantic-elements-and-webflow-the-essential-guide)

### CSS

`CSS` adalah dekorasi web. CSS ditulis untuk menambah [atribut dekorasi](https://www.w3schools.com/cssref/) setiap elemen yang dituju. Bisa juga menggunakan HTML [class](https://www.w3schools.com/html/html_classes.asp) atau [id](https://www.w3schools.com/html/html_id.asp) untuk menunjuk elemen yang lebih spesifik.

![](https://www.proteusthemes.com/wp-content/uploads/2017/12/Basic-Anatomy-of-a-CSS-Rule.png)

Contoh beberapa elemen dan atribut dekorasi:

<iframe width="100%" height="400" src="//jsfiddle.net/willnode/c35pv4bx/embedded/html,css,result/" allowfullscreen="allowfullscreen" allowpaymentrequest frameborder="0"></iframe>

Resources:

+ [W3School CSS Tutorial](https://www.w3schools.com/css/)
+ [W3School CSS References](https://www.w3schools.com/cssref/)
+ [CSS Cheatsheet (Interaktif)](https://cssreference.io/)
+ [CSS Selector Demo](https://www.w3schools.com/cssref/trysel.asp)

Konsep dan Istilah dasar:

+ [Box Model](https://www.w3schools.com/css/css_boxmodel.asp)
+ [Box Sizing](https://developer.mozilla.org/en-US/docs/Web/CSS/box-sizing)
+ [CSS Flex Guide](https://css-tricks.com/snippets/css/a-guide-to-flexbox/)
+ [CSS Grid Guide](https://css-tricks.com/snippets/css/complete-guide-grid/)
+ [HTML + CSS Table Guide](https://css-tricks.com/complete-guide-table-element/)
+ [Responsive CSS (@media rule)](https://www.w3schools.com/cssref/css3_pr_mediaquery.asp)

### JS

`JS` adalah interaksi web. JS adalah bahasa programming ditulis khusus untuk menambah interaksi didalam web.

<iframe width="100%" height="400" src="//jsfiddle.net/willnode/q7n6a4d2/embedded/" allowfullscreen="allowfullscreen" allowpaymentrequest frameborder="0"></iframe>

+ [W3School JS Reference](https://www.w3schools.com/js/default.asp)
+ [W3School JS Tutorial](https://www.w3schools.com/js/exercise_js.asp?)

### Menjadikannya satu HTML

Tiga file tersebut bisa dijadikan satu HTML. Bagaimana formatnya?

```html
<!DOCTYPE html> <!-- HTML Doctype -->
<html lang="en">  <!-- HTML -->

<head>  <!-- HEAD -->
  <meta charset="UTF-8"> <!-- Tipe Encoding (UTF-8 adalah Standard sekarang) -->
  <title>Judul Dokumen</title> <!-- Judul -->

  <!-- Link ke External CSS -->
  <link rel="stylesheet" href="https://link.to/external/file.css">
  <style>
	/* CSS ditaruh disini */
  </style>
</head>

<body> <!-- BODY -->

  Konten HTML Berada disini <!-- Konten HTML diisi disini -->

  <!-- HTML merekomendasikan Javascript untuk ditaruh dibawah Body -->

  <!-- Link ke External JS -->
  <script src="https://link.to/external/file.jss"></script>
  <script>
    /* JS ditaruh disini */
  </script>

</body>
</html>
```

Codepen mempunyai definisi sendiri untuk menyatukan HTML/CSS/JS: [Lihat disini](https://blog.codepen.io/documentation/features/preview-template/).

## Format Tambahan

Tahukah kamu bagaimana Web developers membangun sebuah website? Jika kau telusuri baik-baik, 3 bahasa itu saja tidak akan cukup. Termasuk [blog ini](https://github.com/willnode/blog), anda membutuhkan bahasa lebih dari sekedar HTML/CSS/JS.

### Markdown

`Markdown` adalah HTML Prepocessor yang paling populer. Ia adalah cara populer untuk menulis artikel didalam HTML. Sering digunakan jika anda menulis di GitHub, Wordpress, Jekyll, dan beberapa site generator lainnya.

<iframe height="600" style="width: 100%;" scrolling="no" title="Markdown Previewer" src="//codepen.io/2alin/embed/ZmoyzK/?height=600&theme-id=0&default-tab=resultundefined" frameborder="no" allowtransparency="true" allowfullscreen="true">
</iframe>

Ingat, bahwa browser tidak paham bahasa Markdown, sehingga perlu untuk di-transpile menjadi HTML. Anda bisa mencoba Markdown dengan [Codepen](https://blog.codepen.io/documentation/editor/using-html-preprocessors/).

Referensi:

+ [Github Markdown Guide](https://guides.github.com/features/mastering-markdown/)
+ [Kramdown (Digunakan untuk Jekyll)](https://kramdown.gettalong.org)

### Sass/SCSS

`Sass` (dan `SCSS`) adalah CSS Processor yang paling populer. Ia memfasilitasi CSS dengan fitur-fitur yang tidak dipunyai CSS mulai dari nesting, mixins, partial, dan masih banyak lainnya. Jika kamu menjumpai CSS yang panjang dan ribet, SCSS bisa menjadi solusi alternatif.

Ingat sekali lagi, bahwa browser tidak paham bahasa SCSS, sehingga perlu untuk di-transpile menjadi CSS. Anda bisa mencoba SCSS dengan [Codepen](https://blog.codepen.io/documentation/editor/using-html-preprocessors/).

<iframe height="400" style="width: 100%;" scrolling="no" title="CSS fox" src="//codepen.io/mcvmel/embed/wRLLWW/?height=400&theme-id=0&default-tab=css,resultundefined" frameborder="no" allowtransparency="true" allowfullscreen="true">
</iframe>

Referensi:

+ [Intro to Sass](https://codepen.io/sasstantrum/post/intro-to-sass)
+ [Sass Guide](https://sass-lang.com/guide)

### PHP

Didalam membangun sebuah web, ada dua jenis scripting. `Javascript` adalah *Client-side scripting*, sedangkan `PHP` adalah *Server-side scripting*. Bedanya ialah eksekusi PHP terletak pada Server, sedangkan JavaScript terletak pada browser (client). Karena `PHP` terletak pada server, anda perlu menyiapkan sebuah hosting yang bisa memproses PHP, anda tak bisa menjalankan PHP pada hosting yang menyediakan web statis (seperti GitHub Pages atau Codepen).

Meskipun demikian, anda tetap bisa menjalankan PHP dalam localhost, menggunakan XAMPP atau aplikasi lain sejenisnya.

Referensi:

+ [Install XAMPP](https://www.apachefriends.org/download.html)
+ [W3School PHP](https://www.w3schools.com/php/)

### Penutup

Bagaimana? Sudah tahu gambaran apa saja komponen web?

Jika anda sudah siap, anda bisa lanjut untuk menyiapkan [website pertama anda](/teknik-hosting-web-gratis.html).