---
title: Images Please?
---

Well, it seems that i'm not placing images in my blog lately.

## Why?

Using Git for placing images or binary files is a bad idea, because Git motto

> You can't change your history unless you want the mess

This means once I commit an image, I can't delete it. Even if I can, it's still in history, hence waste space.

The solution however very obvious:

## Host Images Externally

Yes! I have [Cloudinary](https://cloudinary.com).

![Funny cat]({{site.img}}Comical_Animals-0002.jpg)

Obviously the first image that gets uploaded via external service!

Of course the uploading process must be easy peasy, that's why I make my own uploader [here](https://wellosoft.net/cloudup). Now blogging just made easier!