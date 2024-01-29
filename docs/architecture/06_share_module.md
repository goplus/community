# Share module

## Module purpose

This module mainly shares technical articles with Twitter and Facebook

## Module scope

This module mainly shares technical articles with Twitter and Facebook which only need current page url and these third-party platforms will automatically parse the meta tags in the URL page

If the Meta tags are parsed incorrectly or passed to the wrong parameter, the parsing will fail and only share url address without detail information. 

## Module structure
None

## Module Interface
None


## Functions

### share

Example:

```js
function shareOnTwitter() {
var url 
var intentUrl = "https://twitter.com/intent/tweet?"
        + "&url="      + encodeURIComponent(url);
window.open(intentUrl, "_blank", "width=550,height=420"); 
}


function shareOnFaceBook() {
var url 
var intentUrl = "http://www.facebook.com/sharer/sharer.php?"
        + "&u=" + encodeURIComponent(url) ;
window.open(intentUrl, "_blank", "width=550,height=420"); 
}

```