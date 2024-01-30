# Markdown module

## Module purpose

The markdown module uses cherry-markdown to implement a markdown online editor, with an editing area on the left and a preview area on the right. Added highlighting for go+ languages.

## Module scope

- Module outputs user-edited markdown text, and inputs existing user markdown text.
- Module provides go+highlight operation and file upload.

## Module structure

None.

## Module Interface



## Functions

### gop rendering

Example:

```gop
// go+ code
println ""
```

### submit_markdown

- Function: Submits a user-edited markdown document.
- input: none
- Returns: None
- Error: axios request exception

Example:

```js
async submit_markdown() {
    let data = {
        title: 'Default Title',
        content: this.getCherryContent(),
        html_content: this.getCherryHtml()
    }
    axios({
        method: 'post',
        url: '/commit',
        data: data,
        headers: {
        }
    })
    .then(response => {
        console.log('内容发送成功');
        console.log(response.data);
        
    })
    .catch(error => {
        console.error('内容发送失败');
        console.error(error);
    });
},
```