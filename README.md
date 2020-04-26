# Exercise #4: HTML Link Parser

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/link)

## Exercise details

In this exercise your goal is create a package that makes it easy to parse an HTML file and extract all of the links (`<a href="">...</a>` tags). For each extracted link you should return a data structure that includes both the `href`, as well as the text inside the link. Any HTML inside of the link can be stripped out, along with any extra whitespace including newlines, back-to-back spaces, etc.

Links will be nested in different HTML elements, and it is very possible that you will have to deal with HTML similar to code below.

```html
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
```

In situations like these we want to get output that looks roughly like:

```go
Link{
  Href: "/dog",
  Text: "Something in a span Text not in a span Bold text!",
}
```

Once you have a working program, try to write some tests for it to practice using the testing package in go.


### Notes

**1. Use the x/net/html package**

I recommend checking out the [x/net/html](https://godoc.org/golang.org/x/net/html) package for this task, which you will need to `go get`. It is provided by the Go team, but isn't included in the standard library. This makes it a little easier to parse HTML files.


**2. Ignore nested links**

You can ignore any links nested inside of another link. Eg with following HTML:

```html
<a href="#">
  Something here <a href="/dog">nested dog link</a>
</a>
```

It is okay if your code returns only the outside link.

**3. Get something working before focusing on edge-cases**

Don't worry about having perfect code. Chances are there will be a lot of edge cases here that will be kinda tricky to handle. Just try to cover the most basic use cases first and then improve on that.

**4. A few HTML examples have been provided**

I created a few simpler HTML files and included them in this repo to help with testing. They won't cover all potential use cases, but should help you start testing out your code.


**5. The fourth example will help you remove comments from your link text**

Chances are your first version will include the text from comments inside a link tag. Mine did. Use [ex4.html](ex4.html) to test that case out and fix the bug.

*Hint: See [NodeType](https://godoc.org/golang.org/x/net/html#NodeType) constants and look for the types that you can ignore.*


## External Resources

In the solution for this exercise I end up using a DFS, which is a graph theory algorithm. If you want to learn a little more about that, I have discussed it on YouTube here - <https://www.youtube.com/watch?v=zboCGDMnU3I>

There is a complete series on algorithms and graph theory, though at this time it is somewhat incomplete. I never have enough time in the day 🙁. Hopefully one day *Let's Learn Algorithms* will be its own series like *Gophercises*.

## Bonus

The only bonuses here are to improve your tests and edge-case coverage.
