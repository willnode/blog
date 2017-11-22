---
title:  "Hello World!"
---

Hello!, this is the first post of the blog.

Meanwhile, this first post is gonna be my testing playground with Jekyll + Kramdown.

<!-- This is HTML comment, but works well anyway -->
<!-- See http://wellosoft.net/blog/en/2017/hello-world.html for preview -->

## The playground
#### Not yet
###### Yet ...

***
<br>
Now, you may see an `inline code` or _italic_ or **bold** or ~~striking~~ texts.

```cs
void PerhapsSomeCode() { /* coding stuff */ }
```

> And some nifty quote here
> Here as well

| Table | Here |
|:--|--:|
|Hello|World|
|Bye|Bye|

![Made with MSPAINT :D][image]{:title="Made with the legendary MSPAINT!"}

May need some citation here[^1]

* Silver
  + Made
  + With
    - Some
* Bullet
  1. For
  2. Now
     As well

<details><summary>Example</summary>
<code>of a spoiler</code></details>
<br>
Keep an eye because i'm [Linking external page][absolute] or [Linking this page][relative] like <http://github.com>

Try with --- (dashes) and ... (dots) and << arrows >>

[absolute]: https://github.com/
[relative]: {{ "/en/" | relative_url }}
[image]: {{ "/images/2017/madeinmspaint.png" | relative_url }}
[^1]: Hello world

## The Gist

I love Kramdown, it can do so many things without need for any plugin. Anyway here's the gist if you wonder:

{% include gist.html gist="willnode/8443d9b88960ae8c810cb79a2b89c9b7" %}