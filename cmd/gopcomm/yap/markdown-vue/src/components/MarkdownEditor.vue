<template>
    <div id="markdown-container"></div>
</template>

<script>
    import Cherry from 'cherry-markdown'
    import axios from 'axios'

    import Plyr from 'plyr';
    import 'plyr/dist/plyr.css'

    // axios.defaults.baseURL = 'http://localhost:8080/';

    var cherrInstance = null

    var vtt_src = ""

    let fileType = "video/mp4"

    let vtt_id = []

    function getFileTypeByExtension(fileName) {
        var extension = fileName.split('.').pop().toLowerCase();
        var imageExtensions = ['jpg', 'jpeg', 'png', 'gif'];
        var videoExtensions = ['mp4', 'avi', 'mov', 'mkv', 'wmv', 'flv', 'webm', 'avchd'];
        var audioExtensions = ['mp3', 'wav', 'ogg'];
        var documentExtensions = ['doc', 'docx', 'pdf', 'txt'];

        if (imageExtensions.includes(extension)) {
            return 'image/' + extension;
        } else if (videoExtensions.includes(extension)) {
            return 'video/' + extension;
        } else if (audioExtensions.includes(extension)) {
            return 'audio/' + extension;
        } else if (documentExtensions.includes(extension)) {
            return 'document/' + extension;
        } else {
            return '';
        }
    }

    function fileUpload(file) {
        let type = getFileTypeByExtension(file.name)
        console.log("upload", type)

        // if(type.includes("video") && checkVideo()){
        //     console.log("only can upload a video.")
        //     return 
        // }

        var formData = new FormData()
        formData.append('file', file)
        // axios post 请求                
        axios({
            method: 'post',
            url: '/api/media',
            data: formData,
            headers: {
                'Content-Type': 'multipart/form-data'
            }
        })
            .then(response => {
                console.log(response.data); 
                if(type.includes("video")) {
                    axios({
                        method: 'get',
                        url: '/api/video/'+response.data,
                    }).then(res => {
                        console.log("get video")
                        let url = res.data.url.fileKey
                        let stitle = res.data.url.subtitle
                        let status = res.data.url.status
                        let newType = res.data.url.type
                        if(status !== "1"){
                            vtt_id.push(response.data)
                        }
                        vtt_src = stitle
                        
                        console.log("URL", url)
                        onloadCallback(newType, url)                       
                    })
                } else {
                    axios({
                        method: 'get',
                        url: '/api/media/'+response.data+'/url',
                    }).then(res => {
                        let url = res.data.url
                        console.log(url)
                        onloadCallback(type, url)                        
                    })
                }
            
            })
            .catch(error => {
                console.error('文件上传失败！');
                console.error(error);
            });
    }

    function checkVideo(){
        const md = getCherryContent()
        const regex = /!video\[[^\]]*\]\([^)]*\)(\([^)]*\))?/;

        if (regex.test(md)) {
            console.log("have video")
            return true
        }
        return false
    }

    function onloadCallback(type, url) {
        console.log("call back")
        let imgMdStr = ""
        type = type.toLowerCase()
        fileType = type
        if (/mp4|avi|rmvb|mov|wmv|flv|avi|webm/i.test(type.toLowerCase())) {
            // replace <video controls src=""> <video> into !video
            imgMdStr = `!video[${type}](${url})(${vtt_src})\n`;
        } else if (/mp3/i.test(type)) {
            imgMdStr = `!audio[${type}](${url})\n`;
        } else if (/bmp|gif|jpg|jpeg|png/i.test(type)) {
            imgMdStr = `![${type}](${url})\n`
        } else {
            imgMdStr = `[${type}](${url})\n`
        } 
        cherrInstance.insert(imgMdStr) 
    }

    function urlProcessor(url, srcType) {
        return url;
    }
            
    function afterChange(text, html) {
        this.content = text
        let video = document.querySelectorAll('video');
        for (var i = 0; i < video.length; i++) {
            const player = new Plyr(video[i], {captions: {active: true, update: true, language: 'en'}});
        }
    }

    function onCopyCode(event, code) {
        return code;
    }

    function changeString2Pinyin(string) {
        var pinyin = require("./pinyin/pinyin.js");
        return pinyin.pinyin(string, " ");
    }

    function setMarkdown(content) {
        console.log("setMarkdown", content)
        if (!cherrInstance) { 
            initCherryMD(content)
            return
        }
        // this.cherrInstance.setMarkdown(content, 0)
        cherrInstance.insert(content, false, false, true)

    }

    function getCherryContent() {
        if (!cherrInstance) { 
            initCherryMD("")
            // return
        }
        var result = cherrInstance.getMarkdown()
        return result
    }

    function getCherryHtml() {
        var result = cherrInstance.getHtml()
        return result
    }

    function initCherryMD(value, config) {
        var defaultValue = value || ""

        var myBlockHook = Cherry.createSyntaxHook('myBlock', Cherry.constants.HOOKS_TYPE_LIST.PAR, {
            makeHtml(str) {
                return str.replace(this.RULE.reg, function(whole, m1) {
                    // []
                    const regexType = /\[(.*?)\]/g; 
                    const matchType = whole.match(regexType); // 使用 match 函数进行匹配
                    if (matchType){
                        fileType = matchType[0].replace("[","").replace("]","")
                    }
                    // ()
                    console.log("all video ", whole)
                    const regex = /\((.*?)\)/g; // 定义正则表达式，其中 \() 表示左括号，(.*?) 表示非贪婪模式匹配任意字符，\)) 表示右括号
                    const matchResult = whole.match(regex); // 使用 match 函数进行匹配
                    if (matchResult) {
                        let video_src = matchResult[0].replace("(","").replace(")","")
                        console.log("video src", video_src)
                        let poster = video_src + "?vframe/jpg/offset/7"
                        if(matchResult[1]){
                            vtt_src = matchResult[1].replace("(","").replace(")","")
                        }
                        const p = new Plyr('video', {captions: {active: true}});
                        return `<div><video controls="" crossorigin="" playsinline="" data-poster=${poster}><source src=${video_src} type=${fileType} size="576"/><track kind="captions" label="English" srclang="en" src=${vtt_src} default/><a href=${video_src} download>Download</a></video></div>`
                        
                    } else {
                        console.log("can't match ()");
                        let r = "show video failed"
                        return `<div style="border: 1px solid;border-radius: 15px;background: gold;">${r}</div>`;
                    }
                });
            },
            
            rule(str) {
                return {
                    reg: /!video\[.*?\]\(.*?\)\(.*?\)/g
                }
            },
        });

        // import 'https://unpkg.com/vue@2.6.12/dist/vue.min.js'
        var CustomHookA = Cherry.createSyntaxHook('important', Cherry.constants.HOOKS_TYPE_LIST.SEN, {
            makeHtml(str) {
                console.log("custom hook", str)
                return str.replace(this.RULE.reg, function(whole, m1, m2) {
                    return `<span style="color: green;"><strong>${m2}</strong></span>`;
                });
            },
            rule(str) {
                console.log("rule", str)
                return { reg: /(\*\*\*)([^\*]+)\1/g };
            },
        })
        
        cherrInstance = new Cherry({
            id: 'markdown-container',                    
            value: defaultValue,
            customSyntax: {
                importHook: {
                    syntaxClass: CustomHookA, // 将自定义语法对象挂载到 importHook.syntaxClass上
                    force: true, // true： 当cherry自带的语法中也有一个“importHook”时，用自定义的语法覆盖默认语法； false：不覆盖
                    before: 'fontEmphasis', // 定义该自定义语法的执行顺序，当前例子表明在加粗/斜体语法前执行该自定义语法
                },
                myBlock: {
                    syntaxClass: myBlockHook,
                    force: true,
                    before: 'blockquote',
                },
            },
            fileUpload: fileUpload,
            // 第三方包
            externals: { // externals
            },
            // 解析引擎配置
            engine: {
                // 全局配置
                global: {
                    // 是否启用经典换行逻辑
                    // true：一个换行会被忽略，两个以上连续换行会分割成段落，
                    // false： 一个换行会转成<br>，两个连续换行会分割成段落，三个以上连续换行会转成<br>并分割段落
                    classicBr: false,
                    /**
                     * 全局的URL处理器
                     * @param {string} url 来源url
                     * @param {'image'|'audio'|'video'|'autolink'|'link'} srcType 来源类型
                     * @returns
                     */
                    urlProcessor: urlProcessor,
                    htmlWhiteList: ''
                },
                // 内置语法配置
                syntax: {
                    // 语法开关
                    // 'hookName': false,
                    // 语法配置
                    // 'hookName': {
                    //
                    // }
                    autoLink: {
                        /** 是否开启短链接 */
                        enableShortLink: true,
        
                        /** 短链接长度 */
                        shortLinkLength: 20
                    },
                    list: {
                        listNested: false,
                        // 同级列表类型转换后变为子级
                        indentSpace: 2 // 默认2个空格缩进
                    },
                    table: {
                        enableChart: false // chartRenderEngine: EChartsTableEngine,
                        // externals: ['echarts'],

                    },
                    inlineCode: {
                        theme: 'red'
                    },
                    codeBlock: {
                        theme: 'dark',
                        // 默认为深色主题
                        wrap: true,
                        // 超出长度是否换行，false则显示滚动条
                        lineNumber: true,
                        // 默认显示行号
                        copyCode: true,
                        // 是否显示“复制”按钮
                        customRenderer: { // 自定义语法渲染器
                            // 创建自定义渲染函数
                            gop: {
                                render: (src, sign, cherryEnding)=> {
                                    console.log("custom render", src)
                                    // return `<p style="color: red">${src}</p>`;
                                    return `<goplus-code half-code language="gop" style="width: 85vw">${src}</goplus-code>`;
                                }
                            }
                        },
                        mermaid: {
                            svg2img: false, // 是否将mermaid生成的画图变成img格式
                        },
                        indentedCodeBlock: true
                    },
                    emoji: {
                        useUnicode: false, // 是否使用unicode进行渲染
                        customResourceURL: 'https://github.githubassets.com/images/icons/emoji/unicode/${code}.png?v8',
                        upperCase: true,
                    },
                    fontEmphasis: {
                        /**
                         * 是否允许首尾空格
                         * 首尾、前后的定义： 语法前**语法首+内容+语法尾**语法后
                         * 例：
                         *    true:
                         *           __ hello __  ====>   <strong> hello </strong>
                         *           __hello__    ====>   <strong>hello</strong>
                         *    false:
                         *           __ hello __  ====>   <em>_ hello _</em>
                         *           __hello__    ====>   <strong>hello</strong>
                         */
                        allowWhitespace: false
                    },
                    strikethrough: {
                        /**
                         * 是否必须有前后空格
                         * 首尾、前后的定义： 语法前**语法首+内容+语法尾**语法后
                         * 例：
                         *    true:
                         *            hello wor~~l~~d     ====>   hello wor~~l~~d
                         *            hello wor ~~l~~ d   ====>   hello wor <del>l</del> d
                         *    false:
                         *            hello wor~~l~~d     ====>   hello wor<del>l</del>d
                         *            hello wor ~~l~~ d     ====>   hello wor <del>l</del> d
                         */
                        needWhitespace: false
                    },
                    mathBlock: {
                        engine: 'MathJax',
                        // katex或MathJax
                        src: 'https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-svg.js', // 如果使用MathJax plugins，则需要使用该url通过script标签引入
                        plugins: true // 默认加载插件
                    },
                    toc: {
                        /** 默认只渲染一个目录 */
                        allowMultiToc: false,
                    },
                    header: {
                        /**
                         * 标题的样式：
                         *  - default       默认样式，标题前面有锚点
                         *  - autonumber    标题前面有自增序号锚点
                         *  - none          标题没有锚点
                         */
                        anchorStyle: 'default',
                    },
                    inlineMath: {
                        engine: 'MathJax',
                        // katex或MathJax
                        src: 'https://cdn.jsdelivr.net/npm/mathjax@3/es5/tex-svg.js' // 如果使用MathJax plugins，则需要使用该url通过script标签引入
                    },
                    header: {
                        /**
                         * 标题的样式：
                         *  - default       默认样式，标题前面有锚点
                         *  - autonumber    标题前面有自增序号锚点
                         *  - none          标题没有锚点
                         */
                        anchorStyle: 'default'
                    }
                },
                customSyntax: {
                    myBlock: {
                        syntaxClass: myBlockHook,
                        before: 'blockquote',
                    },
                },
            },
            editor: {
                id: 'code',
                // textarea 的id属性值
                name: 'code',
                // textarea 的name属性值
                autoSave2Textarea: false,
                // 是否自动将编辑区的内容回写到textarea里
                theme: 'default',
                // depend on codemirror theme name: https://codemirror.net/demo/theme.htm
                // 编辑器的高度，默认100%，如果挂载点存在内联设置的height则以内联样式为主
                height: '100%',
                // defaultModel 编辑器初始化后的默认模式，一共有三种模式：1、双栏编辑预览模式；2、纯编辑模式；3、预览模式
                // edit&preview: 双栏编辑预览模式
                // editOnly: 纯编辑模式（没有预览，可通过toolbar切换成双栏或预览模式）
                // previewOnly: 预览模式（没有编辑框，toolbar只显示“返回编辑”按钮，可通过toolbar切换成编辑模式）
                defaultModel: 'edit&preview',
                // 粘贴时是否自动将html转成markdown
                convertWhenPaste: true,
                codemirror: {
                    // 是否自动focus 默认为true
                    autofocus: true
                }
            },
            toolbars: {
                theme: 'light',
                // light or dark
                showToolbar: true,
                // false：不展示顶部工具栏； true：展示工具栏; toolbars.showToolbar=false 与 toolbars.toolbar=false 等效
                toolbar: ['bold', 'italic',
                    {
                        strikethrough: ['strikethrough', 'underline', 'sub', 'sup', 'ruby', 'customMenuAName'],
                    },
                    'size','|', 'color', 'header', '|','drawIo', '|', 'list',
                    'panel', 
                    // 'justify', // 对齐方式，默认不推荐这么“复杂”的样式要求
                    'detail', '|', {
                        insert: ['image', 'audio', 'video', 'link', 'hr', 'br', 'code', 'formula', 'toc', 'table',
                            'line-table', 'bar-table', 'pdf', 'word'
                        ]
                    }, 'graph', 'export', 'codeTheme', 'switchModel', 'togglePreview',
                    // {
                    //   customMenuBName: ['ruby', 'audio', 'video', 'customMenuAName'], //实验室
                    // },
                    // 'settings', 'customMenuCName', 'customMenuFName', 'theme', 

                    'settings', 'theme', 
                ],
                // toolbarRight: ['fullScreen', '|', 'customMenuFName'],
                toolbarRight: ['fullScreen'],
                bubble: ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', 'ruby', '|', 'size',
                    'color'
                ],
                // array or false
                float: ['h1', 'h2', 'h3', '|', 'checklist', 'quote', 'quickTable', 'code'], // array or false
                toc: {
                updateLocationHash: false, // 要不要更新URL的hash
                defaultModel: 'full', // pure: 精简模式/缩略模式，只有一排小点； full: 完整模式，会展示所有标题
                },
                sidebar: ['mobilePreview', 'theme', 'copy'], // 'copy',
                customMenu: {
                    // customMenuAName: customMenuA,
                    // customMenuBName: customMenuB,
                    // customMenuCName: customMenuC,
                    // customMenuFName: customMenuF
                },
            },
            // 打开draw.io编辑页的url，如果为空则drawio按钮失效
            drawioIframeUrl: window.location.origin + '/CherryMarkdown/drawio_demo.html',
            /**
             * 上传文件的时候用来指定文件类型
             */
            fileTypeLimitMap: {
                video: 'video/*',
                audio: 'audio/*',
                image: 'image/*',
                word: '.doc,.docx',
                pdf: '.pdf',
                file: '*',
            },
            callback: {
                afterChange: afterChange,
                // afterInit: afterInit,
                // beforeImageMounted: beforeImageMounted,
                // // 预览区域点击事件，previewer.enablePreviewerBubble = true 时生效
                // onClickPreview: onClickPreview,
                // 复制代码块代码时的回调
                onCopyCode: onCopyCode,
                // 把中文变成拼音的回调，当然也可以把中文变成英文、英文变成中文
                changeString2Pinyin: changeString2Pinyin
            },
            // 预览页面不需要绑定事件
            isPreviewOnly: false,
            // 预览区域跟随编辑器光标自动滚动
            autoScrollByCursor: true,
            // 外层容器不存在时，是否强制输出到body上
            forceAppend: true,
            // The locale Cherry is going to use. Locales live in /src/locales/
            locale: "en_US",
        })
        console.log(cherrInstance.engine)
    }

    import 'cherry-markdown/dist/cherry-markdown.min.css'

    export default {
        name: "GoplusMarkdown",
        mounted() {
            initCherryMD()
        },
        methods: {    
            getMarkdown() {
                return getCherryContent()
            },      
            getCherryHtml() {
                return getCherryHtml()
            },   

            insertMarkdown(content) {
                setMarkdown(content)
            },  
            setMarkdown(content) {
                // setMarkdown(content)
                cherrInstance.setMarkdown(content)
                cherrInstance.editor.editor.setCursor(cherrInstance.getMarkdown().length)
            },
            getVttId() {
                return vtt_id
            },
            checkVideo(){
                return checkVideo()
            }
        }
    }
</script>