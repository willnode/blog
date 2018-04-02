---
title: NodeJS and microservices
---

So I just heard about Node.JS, and it is awesooooooome!

## Javascript is Awesome

NodeJS is an Javascript Intepreter, built because making a HTTP server is suck at the time. But because Javascript was suck too, it had a clever solution to make it modular:

```js
var foo = require('bar');
```

```js
module.exports = 'baz';
```

See? no new syntax is necessary. Simple enough to make it so modular. No external project config needed (like C#, which is also a reason that I start to look into Go for small tools).

And look at those new syntax sugars that Javascript had evolved in the last few years. NodeJS simply had a good investment on that.

## Microservices

Microservices is like you have a server but instead of writing everything on it, you're pulling specific server-side operation off from the main server, which is very cool especially for small sites like mine.

If you're looking for headstart, have a look to [micro](https://github.com/zeit/micro) and start from this code:

```js
module.exports = function(request, response) => {
	response.end('hello world!')
}
```

then run it with `$ micro`. Simply awesome, right?!

I have two microservices running on my site, [gh-latest-repos](https://github.com/sindresorhus/gh-latest-repos) and [uas-gitlab-auth](https://github.com/willnode/uas-gitlab-auth). I deploy those microservices with Now, which is free (enough for small quantity) and surprisingly simple to setup (just click the [button](https://deploy.now.sh/)).

Most microservices today are written in NodeJS, and yep, I learn a lot from learning how they work in code basis, especially after writing that uas-gitlab-auth and [bandit-proxy](https://github.com/willnode/bandit-proxy).

## In conclusion I'm....

So excited to create another microservice with Node.JS! ✨
