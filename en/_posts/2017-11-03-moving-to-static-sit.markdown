---
layout: post
title: "Moving to Static Site (and why)"
---

So I have released my self from the tiranny of `wordpress`. In the past year, I used to post everything there. It was nice, and then problem comes...

```
Note that this post is never be intended to critize wordpress.
```

First thing is that I use the free `wordpress.com` simply because I can't afford `4$` a month. However this means I have to get used with limited style customization, ads on the homepage and terrible posts system. Note the posts system I mention here, because it forced me to create two different blog `wellosoft.wordpress.com` and `willandgottaloveideas.wordpress.com` just for supporting multi-language, and It's ain't good.

## GitHub to the Rescue

Early this year, I started to become active by the time [GitHub Desktop][gitdesktop] is born (why? pushing to github via CLI is really complicated). It started to change my workflow so much that I always `git init` before starting any new project.

## Problem with Pages

The ability of GitHub to create personal static website catch my interest immediately, but it *was* hard:

+ Everything need to be done in CLI (there wasn't GitHub Desktop).
+ You have to put the page in either `master` or orphan `gh-pages` branch which is kinda complicated.
+ **I failed to setup Jekyll**. So bad that installation is really complicated for windows.

With these challenges I still managed to create my first gh-pages for my project [ComCalc][comcalc]. It was hard, so I don't look at it anymore for few months.

## Here Comes the Light

I can't remember when I first used `GitHub Desktop` but in the previous month, I want to create a documentation website for [Engine4][engine4]. It comes to surprise that [DocFX][docfx] is exactly what I needed. After few days I managed to make it online and I was really happy. I love how easy to install `DocFX` using [Chocolatey][choco] compared with [Jekyll][jekyll].

> And then everything has changed

I can't believe how things changed rapidly. By the time, GitHub support hosting pages via folder `/docs` and it's motivate me to create another website to host list of my projects, again using DocFX (yeah, really). I put the generated content under `/docs` and then put it online [here][expertise]. 

So I have two new website in under a month, and it's dramatically gain my understanding in web development.

## The Movement

> I want my website look professional. No matter how it designed my website would still looks like a joke unless I own a domain.

There the problem happen. Wordpress won't allow me to add domain unless I subscribe, and wordpress hosting still kinda expensive. However, things leads me into freedom when I heard that github allow pages to be put under custom domain, for free.

This leads me to make another chance to try Jekyll again. But to may surprise (again), Jekyll also evolve. Now I have a choice to install it in windows via `bash` or `choco`. It takes few hours until I managed to successfully install and execute jekyll into my computer (I choose `choco` and still kinda hard but okay anyway). 

Then after few days, I managed to create the homepage and `/blog/` using Jekyll. It's awesome that both platform (Jekyll and DocFX) are powered by [Liquid][liquid] so I don't have to worry about laboring and manually copy stuff.

Seems like everything is backed by static website, right? This is also the first time I **purchased a domain** `wellosoft.net` and put the`CNAME` to `willnode.github.io`. Lo behold! My professional [website][mysite] is now born :)

## Conclusion

I feel very gratitude to see these things is possible today. All I need is a domain for around `10$` a year which is a win-win for my budget, and then everything else (the hosting) is backed by GitHub.

It's truly powerful. GitHub solves my dilemma in this online era. It's like the Facebook for programmers today. I truly appreciate that the site and its service is available for everyone, for free.

I do fear that this site will get banned when it reaching the [100 GB traffic limit][usagelimit] (which around 400 visits in a hour). But as amateur as I am, I think I don't need to worry after all =)

[gitdesktop]: https://desktop.github.com
[jekyll]: https://jekyllrb.com
[choco]: https://chocolatey.org
[docfx]: https://dotnet.github.io/docfx/
[comcalc]: willnode.github.io/ComCalc/index.html
[expertise]: willnode.github.io/expertise/
[engine4]: willnode.github.io/engine4-doc/
[mysite]: http://wellosoft.net
[usagelimit]: https://help.github.com/articles/what-is-github-pages/#usage-limits
[liquid]: https://shopify.github.io/liquid/