---
title: How Dare you Disqus
---

I always aiming for lightweight size on my blog site. Every content on my site is already precompiled and minified, so it's pretty lightweight and don't require javascript at all.

## But...

I do need Javascript, for:

+ Site traffic (Google Analytics)
+ Comment Platform (Disqus)

So how's the difference between JavaScript turned on and off?

![JS vs no Js]({{site.img}}Js_vs_no_Js.jpg)

Son of a *****

It's worth noting that there's 83 KB image there so my site is only 3 KB. **It's a hundred times bigger thanks to those JavaScript**.

## Solution

+ Change Google Analytics to [ga-lite](https://github.com/jehna/ga-lite). 30 KB down to 8 KB.
+ Lazily load comments when user need to. 300 KB+ down to zero. [This how I achieve it](https://gist.github.com/willnode/052cc2f0b0115c7e4041a7792c6ba8d3).

From now on this site you have to click the button to show comments if you curious. Anyway it's a worth saving, right?