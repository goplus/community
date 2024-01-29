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
``