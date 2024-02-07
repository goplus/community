# Markdown module

## Module purpose

The markdown module uses cherry-markdown to implement a markdown online editor, with an editing area on the left and a preview area on the right. Added highlighting for go+ languages.

We export the markdown module with the UMD format, which can be used as a UI component in the browser.

## Module scope

Module outputs user-edited markdown text, and inputs existing user markdown text.

## Module structure

None.

## Module Interface

Provides a markdown editor and a markdown viewer. Use the markdown editor to edit markdown text, and use the markdown viewer to preview the markdown text as HTML.

```js
import {MarkdownEditor, MarkdownViewer} from './GoplusMarkdown.js';

const { createApp } = Vue

createApp({
  data() {
    return {
      message: 'Hello Goplus Markdown!'
    }
  },
  components: {
    MarkdownEditor: MarkdownEditor,
    MarkdownViewer: MarkdownViewer
  }
}).mount('#app')
```

## Functions

### Markdown editor

- Function: Submits a user-edited markdown document.
- input: None
- Returns: None
- Error: None

Example:

```html
<markdown-editor></markdown-editor>

<script type="module">
import {MarkdownEditor, MarkdownViewer} from './GoplusMarkdown.js';

createApp({
  data() {
    return {
      message: 'Hello Goplus markdown!'
    }
  },
  components: {
    MarkdownEditor: MarkdownEditor,
  }
}).mount('#app')

</script>
```

### Markdown viewer

- Function: Returns the HTML text of the markdown text.
- input: markdown text
- Returns: HTML text
- Error: None

Example:

```html
<markdown-viewer md="```gop ```"></markdown-viewer>

<script type="module">
import {MarkdownEditor, MarkdownViewer} from './GoplusMarkdown.js';

createApp({
  data() {
    return {
      message: 'Hello Goplus markdown!'
    }
  },
  components: {
    MarkdownEditor: MarkdownEditor,
  }
}).mount('#app')

</script>
```