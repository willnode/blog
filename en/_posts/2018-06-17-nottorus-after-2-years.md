---
title: Nottorus After 2 Years
---

2 years ago there was a newcomer, saying that after a year of development, he released a [Visual Scripting plugin for Unity](https://www.assetstore.unity3d.com/en/#!/content/59656), this time with **native compilation** to C#.

*Interesting.*

> I worked very hard on this plugin and after more than one year of developing I'm ready to release it.
> -- [Stridemann](https://forum.unity.com/threads/395750/)

At first glance I'm ignoring this asset because it was initially sold at $350 *per seat*, which is darn big for newcomers. How come? I though this guy must be crazy -- it will only lasts for few weeks...

*But it didn't.*

Early reviews saying it's an impressive plugin, has full potential. Some even says this guy should be hired by Unity. I then read it again in the product description, it claims that *it can generate over 150k nodes by dissambling UnityEngine.dll* and even claiming futher that *the plugin itself can be developed inside the plugin*.

**Mind Blown.**

## I want to get grip with this Asset, however....

First He (was) put [DRM](https://forum.unity.com/threads/395750/page-19#post-3048152) into this asset. That's bit annoying, first ever of asset doing it. However I still **want** to have this asset, but he already set the price *too damn high*: 250 US$!

Seriously man, that is too expensive. People were [cornerning](https://forum.unity.com/threads/395750/#post-2598208) about support for upcoming years, and I can't trust anybody which don't link any kind of social media (because I remember Vizio, the publisher just mysteriously [gone](https://forum.unity.com/threads/59121/page-26#post-2220138) into abyss), and the worst... **He ~~don't~~ wasn't giving the source code!**

Look, I want this asset not because I can't code (btw I'm pro in C# for years already), but it's for education -- I want to introduce code for beginners (nodes is damn easy to read), and know how this damn thing works magic. Closing the source makes me throw away this kind of dream, for 2 years..

## Then Suddenly it's Happening

> this forum will be shut down in April 13. I moved all documentation from site to pdf file in unitypackage also added source code. -- [Stridemann](https://google.com/search?q=cache:nottorus.net/forum/viewthread.php?thread_id=383)

People get pissed and I feel sorry for those who get pissed because Stridemann announcing it big-front in Nottorus description that **support [ends](https://forum.unity.com/threads/395750/page-22#post-3395732) in Apr 2018**. I don't know why, but I wasn't care -- Until something happened.

In recent months, he drop Nottorus price from 250$ to 150$ and call it a day with no announcement at all, then drop it far lower to [100](https://forum.unity.com/threads/395750/page-22#post-3447509) US$ then [75](https://forum.unity.com/threads/395750/page-22#post-3456300) US$, before [shutting](https://forum.unity.com/threads/395750/page-22#post-3455919) down whole Nottorus website including its forum.

That.. Pisses.. Everyone. People care support not price. Some even saying that the latest update (v2.0) is broken. Honestly, I just know these news few weeks ago when checking forums and it's pretty dishearting. However, I noticing one thing that's important...

## The Latest Update contains Source Code

I say "what the \$\$\$\$". Yes, of course, at 75 US$ it's the lowest price of Nottorus that ever be offered. I immedially purchase it with revenues from TEXDraw sale last month (75$ is highest asset I bought, but it'll restore back when 2 people purchase TEXDraw probably in few weeks). I don't care with no support, atleast I got the source code.

## First Play

Before cracking the source code, I played with Nottorus for few hours... Just to get a simple Mouse Rotation working.

![image]({{site.img}}nottorus-camview.png)

Actually it was kinda hard to figure out which node doing the right job (biggest confusion is on where do I insert operators :/ ). But then I know by testing Nottorus parsing capability, and then I say it's really impressive. I had good hours spend with amazing experience of wonders.

## So How complex (in code) is Nottorus?

This is comparison by lines of code written in Nottorus vs. [TEXDraw](http://u3d.as/cco).

![image]({{site.img}}cdmtr-not-tex.png)

The result is shown that Nottorus is approximate 4 times heavier than TEXDraw (I sold TEXDraw at $50 so $200 seems fair price for Nottorus). However I also saw this:

[lots of compiler warning about unused variables]

I don't know why, but those looks like debug-only variables that's left over, and because he always compile it to DLL anyway (and VS don't complain about unused variables by default) so he ignores it. Somebody has to fix it of course, so I did it for my self.

Overall after shaving though overviews of Nottorus code, I really feel like I can bring it down and better. I spotted many commented out codes and sections that's marked as 'WIP'. Also those scripts below Editor with prefix are partials of big big `public sealed partial class Nottorus : EditorWindow` which is too huge to be true.

What makes me wonder is that the addons under folder 'CSP' (abbv. for CSharp Parser, maybe?), because it is written in entirely different style, and looks better and properly documented. That's account for 30% Nottorus code. I have feeling that Stridemann don't write this parser but who knows? I failed to find matching code online (though ISharpCode is closest).

If I'm really going to use Nottorus, I would 1.) Cleaning the code 2.) Fix Annoying UX Failures.

## Closing

Like everybody has said, this asset has full potential, however it is [not](https://forum.unity.com/threads/395750/page-23#post-3465395) on the right dev -- either he burnout, not enough having sales or just choose to fall into abyss. *(I'll pull my words when he came back)*

As much I hate and love parts in Nottorus, I have full respect to its creator. If he didn't show sign of anything about coming back, we know that somebody has to take over the project or the world will [miss](https://forum.unity.com/threads/395750/page-23#post-3502475) this unique masterpiece of software ever after.

## EDIT: Post Mortem

*last edit is on June 21*

Over the last three days I *did* made modifications to Nottorus... well, maybe too much.

![img]({{site.img}}nottorus-gitpatches.png)

Upon refactoring the code I noticed that Nottorus has a lot of strings attached (read: couple of issues). Here I list problems that I spotted and wanted to fix before I can use it for game production:

Problem | Solution | Fixed? | Sidenote |
---|---|---|---|
A lot of obsolete warning popped up upon extracting the source code | Attempt to update code to use Newer API manually | ✔ | I still fixing code style inconsistencies and commented code, +4000 lines of code is stripped as of today |
Lack of keyboard shortcuts (insert, delete, move, attach nodes) | Add those shortcut | (WIP) | Added shortcut: Ctrl+D (duplicate), PageUp + PageDown (Zoom). |
Nottorus sometimes require Unity Restart | "Nottorus -> Flush Cache" | ✔ | This is a temporary solution, but obviously much faster than restarting whole Unity |
Slow Load, Slow Compilation, Slow Saving, Slow Undo Snapshoting | See if I can make it faster | ✔ | Almost done. Loading went fast from 1 minute to just under 5 sec. Slow Compilation have been solved by stripping Precompiler and using StringBuilder. Slow saving and undo snapshot will be solved by (see 👇) |
Big `.nts` file | Find faster serializer ([Odin](https://github.com/TeamSirenix/odin-serializer)?) and propose `.ntsx` if breaking changes introduced | (Yet) | This is my HUGE issue. `.nts` are formatted as XML object, which is far inefficient (I'm proposing `.ntsx` saved as raw JSON or compressed binary file) |
Unseamless `.cs` parsing | Add Right-click .CS -> "Edit with Nottorus", have `.nts` file saved outside Assets (optional option), PRESERVE unmodified part of code | (Yet) | Still looking for best way to do this |
Hardcoded Theme | Move everything to a `GUISkin` File | (Yet) | Still looking for best way to do this |

Nottorus is a big plugin. There's many part from it that I can learn. However refactoring things to solve above problems, is challenging. I feel like a janitor. Tomorrow I'm moving to another project, and probably going back here again to finish what left.

*(PS: I'm working on a new Game, so I'm also trying to make it using Nottorus)*
