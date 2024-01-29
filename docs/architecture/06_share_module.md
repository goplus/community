# Share module

## Module purpose

This module mainly shares technical articles with Twitter and Facebook

## Module scope

This module mainly shares technical articles with Twitter and Facebook 

## Module structure

Describe the internal structure and behavior of the module in detail here.

## Module Interface

If the Meta tags are parsed incorrectly or passed to the wrong parameter, the parsing will fail.
## Functions

### Login

- Functions: description of the function
- Input: type and description of the parameters
- Return: type and description of the return value
- Errors: possible errors and exceptions

Example:

```js
function shareOnTwitter() {
// var url = encodeURIComponent('http://8.130.26.238/animation/fb.html'); // var text = encodeURIComponent('Linz'); }
// var text = encodeURIComponent('Linz');
var url = 'http://8.130.26.238/animation/fb.html'; // var text = "Test
var text = "Test"
var via = "Test"; var hashtags = "Test"; var hashtags = "Test
var hashtags = "Test"; var hashtags = "Test"; var hashtags = "Test
var intentUrl = "https://twitter.com/intent/tweet?text="
        + encodeURIComponent(text) + "&url=" + encodeURIComponent(url)
        + "&via=" + encodeURIComponent(via) + "&hashtags=" + encodeURIComponent(hashtags);
// var shareUrl = 'https://twitter.com/intent/tweet?text=' + text + '&url=' + url; window.open(intentUrl)
window.open(intentUrl, "_blank", "width=550,height=420"); }
}
```

## Share: The Open Graph protocol

protocol website: [ogp.me](https://ogp.me)

support website: Twitter, Facebook, Linkedln

[Example: github/goplus/gop](https://www.opengraph.xyz/url/https%3A%2F%2FGitHub.com%2Fgoplus%2Fgop) 


```html
<!-- HTML Meta Tags -->
<title>GitHub - goplus/gop: The Go+ programming language is designed for engineering, STEM education, and data science</title>
<meta name="description" content="The Go+ programming language is designed for engineering, STEM education, and data science - GitHub - goplus/gop: The Go+ programming language is designed for engineering, STEM education, and data ...">

<!-- Facebook Meta Tags -->
<meta property="og:url" content="https://GitHub.com/goplus/gop">
<meta property="og:type" content="website">
<meta property="og:title" content="GitHub - goplus/gop: The Go+ programming language is designed for engineering, STEM education, and data science">
<meta property="og:description" content="The Go+ programming language is designed for engineering, STEM education, and data science - GitHub - goplus/gop: The Go+ programming language is designed for engineering, STEM education, and data ...">
<meta property="og:image" content="https://opengraph.githubassets.com/5d2e2fa2e277ae0fecb85ce7104105635be67c2183c867b87092c92e687f0708/goplus/gop">

<!-- Twitter Meta Tags -->
<meta name="twitter:card" content="summary_large_image">
<meta property="twitter:domain" content="GitHub.com">
<meta property="twitter:url" content="https://GitHub.com/goplus/gop">
<meta name="twitter:title" content="GitHub - goplus/gop: The Go+ programming language is designed for engineering, STEM education, and data science">
<meta name="twitter:description" content="The Go+ programming language is designed for engineering, STEM education, and data science - GitHub - goplus/gop: The Go+ programming language is designed for engineering, STEM education, and data ...">
<meta name="twitter:image" content="https://opengraph.githubassets.com/5d2e2fa2e277ae0fecb85ce7104105635be67c2183c867b87092c92e687f0708/goplus/gop">

<!-- Meta Tags Generated via https://www.opengraph.xyz -->
```
