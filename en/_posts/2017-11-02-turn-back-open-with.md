---
title: "Turn back \"Open with CMD\""
---

So how lovely command prompt compared to new PowerShell? I don’t know what the trend right now, but I still love `cmd` because it's fast and compact. I don't quite like PowerShell though Microsoft pushes it further now in newer version of windows 10.

Long story short I want to pop cmd easily when I right clicking any folder like I was. It *was* working until it goes replaced with powelshell. I've tried to google anywhere but didn’t found anything useful. So I build mine….

## The idea

Turns out that because I use `vscode`, and it shows `Open with Code` when I right clicking on any folder, I can use the same method to show up another version like  Open with CMD. I just need to banging my self with `regedit` and found that the trick works great.

{% include img.html src="openwithcmd.jpg" %}

I can tell you steps but it is much simpler if I just export for you. In the nutshell, just view raw, download and run:

{% include gist.html gist="willnode/157a5b95bf894e6556cf02455a8d8cba" %}