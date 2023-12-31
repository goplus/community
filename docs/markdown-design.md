# Markdown Subsystem Design

## Technical Analysis

The Markdown parser converts MD text into HTML for display. It mainly analyzes the relevant markdown parser of front-end and back-end.

### js markdown parser

---

Considering that the article needs to be written in MD format, we have researched the following markdown parser related to the JS, all of our surveys are open and free:

|                  | [markdown-it](https://github.com/markdown-it/markdown-it?tab=readme-ov-file)（16.5k） | [cherry-markdown](https://github.com/Tencent/cherry-markdown)（3k） |    [marked](https://github.com/markedjs/marked)（31.1k）     | [showdown](https://github.com/showdownjs/showdown)（13.7k）  | [simpleMDE](https://github.com/sparksuite/simplemde-markdown-editor)（9.6k） | [remarkable](https://github.com/jonschlinkert/remarkable)（5.6k） | [CommonMark](https://github.com/commonmark/commonmark.js)（1.4k） |
| :--------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
|   **feature**    | 1. Each subParser is responsible for parsing a syntax, and the regular expressions are clearer<br />2. There are a lot of [plugins](https://link.zhihu.com/?target=https%3A//www.npmjs.com/search%3Fq%3Dkeywords%3Amarkdown-it-plugin) or other [packages](https://link.zhihu.com/?target=https%3A//www.npmjs.com/search%3Fq%3Dkeywords%3Amarkdown-it), following the [CommonMark ]( https://link.zhihu.com/?target=https%3A//spec.commonmark.org/0.30/%23insecure-characters) , while adding some syntactic sugar | 1. Preview area directly synchronizes modifications <br />2. Copying Html and pasting it into MD format <br />3. Multi-cursor editing, drag and drop the size of the image to modify<br />4. Multiple views are supported | 1. it is resolved by regular parsing, without componentization and modularization, which reduces some of the additional overhead caused by improving readability<br />2. asynchronous parsing | Showdown defines a series of subParsers that can parse different syntax, which is logically clear, and it is easy to register your own subParser or extension to extend the new syntax. | 1. SimpleMDE is one of the first editors to have both built-in auto-save and spell checker. <br />2. The syntax is parsed together using [codemirror](https://github.com/codemirror/codemirror) and relies on [font awesome](http://fontawesome.io/), and the preview is rendered by [marked](https://github.com/chjj/marked) using GFM. | 1. There are a number of "presets" to make it easy to quickly enable/disable active syntax rules and options for some common use cases.<br />2. Plugin functionality is provided | 1. A function is provided with parsing a CommonMark document to an AST and rendering the document as an HTML or XML representation<br />2. Multi-language implementation |
|  **extensions**  |                          by plugins                          | [cherry-extensions](https://github.com/Tencent/cherry-markdown/wiki/%E8%87%AA%E5%AE%9A%E4%B9%89%E8%AF%AD%E6%B3%95) | [Using Pro - Marked Documentation-extensions](https://marked.js.org/using_pro#extensions) |                             yes                              |                             yes                              |                             yes                              | [Proposed Extensions](https://github.com/commonmark/commonmark-spec/wiki/Proposed-Extensions) |
|  **highlight**   | [markdown-it-highlightjs](https://www.npmjs.com/package/markdown-it-highlightjs) |                             yes                              | [Using Advanced - Marked Documentation-highlight](https://marked.js.org/using_advanced#highlight) |                             yes                              |  [highlight.js](https://github.com/isagalaev/highlight.js)   |                             yes                              |                             yes                              |
| **performance**  |                            faster                            | It adopts the method of local update and rendering, which is relatively fast. |                           fastest                            | Due to the introduction of modularity and lifecycle, the efficiency of parsing is significantly reduced |                      Slower than marked                      |                            normal                            | Only compared with PHP, Java, etc. of the same product, JS has the best performance |
|     **vue**      | [markdown-it-vue](https://github.com/ravenq/markdown-it-vue) |                             yes                              |                             yes                              |       [Vue Showdown](https://vue-showdown.js.org/zh/)        |   [vue-simplemde](https://f-loat.github.io/vue-simplemde/)   | [vue-remarkable](https://github.com/katalonne/vue-remarkable/blob/master/README.md) |                                                              |
| **Disadvantage** | Chinese support is not very good.(But at the moment the community is going to be international, and that might not be important） | Documentation updates don't update as fast as code updates.  | 1. A lot of analysis relies on regularization, which is more troublesome to maintain; <br />2. There is no protection against XSS attacks, and additional filtering is required,like （[DOMPurify](https://github.com/cure53/DOMPurify)）<br />3. Not supported [GFM](https://docs.github.com/en/get-started/writing-on-github/working-with-advanced-formatting) | 1. If there is a problem with the implementation of one subParser, it may cause the resolution of other subParsers to be affected as well<br />2. Many of the more important issues were not resolved <br />3. Last updated two years ago | 1. It does not support image drag and drop upload, paste and upload, we may need to implement it ourself ([example](https://zhuanlan.zhihu.com/p/35679561)) <br />2. Last update 8 years ago |                  3. Last update 5 years ago                  |  This format is newer and doesn't support a lot of features  |

Since the final community article needs to extend the GO+ syntax highlighting and also support the video playback function, the focus is on whether the markdown parser supports extension and highlighting. Based on the above tabular analysis,the following conclusions have been drawn:
1. Because showdown, simpleMDE, and remarkable have not been updated for a long time, it is difficult to solve problems in time if they occur during use, so they are eliminated.
2. CommonMark is newer and doesn't support a lot of features,so exclude it.
3. We all know that GFM is very important, so we don't think about marked.
4. After discussing with the front-end classmates and considering that our community is international, we need to consider the translation of markdown,so we chose **cherry** which supports bidirectional parsing and the parsing speed and performance are also satisfactory.


### go markdown parser

---

In addition to researching the markdown parser developed by JS, we also learned about the GO-related parser. Because we researched some relevant blogs on the market, they basically use go markdown parsers. All of our surveys are open and free:

|                  |     [goldmark](https://github.com/yuin/goldmark)（3.1k）     |         [lute](https://github.com/88250/lute)（930）         | [Blackfriday](https://github.com/russross/blackfriday)（5.3k） | [gomarkdown](https://github.com/gomarkdown/markdown)（1.2k） |
| ---------------- | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: | :----------------------------------------------------------: |
| **feature**      | 1. The CommonMarkdown parser relies only on the standard library；<br />2. The Japanese, Korean language support is friendly | 1. Support go and js, which can be called through the front-end, back-end and http.<br />2. In order to implement the implementation specification and ensure that there is no ambiguity in Markdown rendering, the untidy md will be formatted into a unified format and handled by the compilation principle. | 1. Fast speed, on-demand rendering of most web applications.<br />2.  It supports universal extensions<br />3. Parsing directly from the C markdown parser [Sundown](https://github.com/vmg/sundown). | [markdown package](https://pkg.go.dev/github.com/gomarkdown/markdown) |
| **extensions**   |                             yes                              |                             yes                              |       no,  AST uses structures rather than interfaces        |                             yes                              |
| **highlight**    | [goldmark-highlighting](https://github.com/yuin/goldmark-highlighting) | Use the [chroma](https://github.com/alecthomas/chroma) engine for syntax highlighting |                             yes                              |                             yes                              |
| **performance**  |                            medium                            |                            medium                            | The speed is relatively fast, and the rendering is on demand. |                             slow                             |
| **Disadvantage** | Use other libraries (e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday)) for HTML security filtering. | 1. Use other libraries (e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday)) for HTML security filtering.<br />2. The implementation of js is to compile the go language into js using [GopherJS](https://github.com/gopherjs/gopherjs). | 1. Use other libraries (e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday)) for HTML security filtering.<br />2. Last updated two years ago.<br />3. There are two versions, and there are currently some incompatibilities (and different ways of using them)<br />4. GFM was not implemented. | 1. Use other libraries (e.g. [bluemonday](https://github.com/microcosm-cc/bluemonday)) for HTML security filtering.<br />2. Poor performance |

Based on the above tabular analysis,the following conclusions have been drawn:
1. The disadvantages of Blackfriday and gomarkdown are more obvious,it is difficult to meet the needs of community articles.
2. Compared with Goldmark, Lute has fewer users and starts late, so we chose **Goldmark**.

### scheme

#### Go+JS Markdown Parser

After researching some existing blog systems, including [Feishu Online Documents](https://eq2pyit41ih.feishu.cn/drive/home/), [CSDN](https://www.csdn.net/), [Chaindrop](https://ld246.com/), etc., which are all  return to the front-end static page. When editing, they use the front-end markdown editor (except for Chaindrop, which does not have a real-time preview function). GitHub's markdown is compiled by the back-end, so **we also plan to use the back-end markdown to parse**. There is a high probability that it will not be modified after the article is published, so the back-end analysis only needs to be parsed and stored, while the front-end markdown parser, everyone needs to parse it every time when it is opened.

Although the back-end parser is used when displaying the article, the real-time preview effect must be better when writing the article using front-end markdown parser, because the modification is particularly frequent at this time, and there is no need to send it to the back-end to parse and then return to the front-end, which is sure to be real-time is not good, so the **front-end parser is used when writing the article**.

Based on the above analysis, it has been decided to use **cherry-markdown(js markdown parser)** when writing the article for real-time compilation, while using **goldmark(go markdown parser)** to publish articles for one-time compilation.

#### JS Markdown Parser

Another scheme is using the front-end markdown editor to parse and then transfer the parsed HTML results to the back-end for storage. In this way, it is possible to use only one markdown parser to parse the article once when we publish it, and we don't need to parse it again when  viewing it. 

Based on the above analysis, **We can implement both real-time preview and one-time compilation by using a front-end parser,  named cherry-markdown.**

#### result

---

Although *Go+JS Markdown Parser* is more flexible, due to the different parsers of the two parsers, there may be a situation where the format of the article previewed and displayed in real time does not match, so we choose ***JS Markdown Parser Scheme（cherry-markdown）*** for markdown parsing. 

