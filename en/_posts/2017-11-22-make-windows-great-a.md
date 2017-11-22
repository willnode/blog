---
title: "Make Windows Great Again"
---

I just updated my windows to the latest version `1709`. Unfortunately some oldie hack to make my computer smooth is gone as most is gone replaced.

So then I made the all-in-one bat file for it.

The bat file will does this:

+ Making restore point (optional)
+ Disabling commonly heavy-task services
+ Disabling common startup task by schedulers
+ Disabling cortana
+ Remove telemetry data
+ Some privacy and speed tweak
+ Removing windows store apps (including edge)
+ Uninstall onedrive
+ Restart explorer

As you execute this script with elevated permission, your computer will (hopefully) turns to be very lightweight and speedy.

This may also removing some part that essential for you. But for me, it works pretty well. With 2 GB RAM windows only fill 40% memory without any CPU or disk spike at startup. Very efficient!

Download and run the script below with elevated access:

> Be caution with what you're doing, even it really won't bite :)

{% include gist.html gist="willnode/056defa0107a5998426f238699500d88" %}
