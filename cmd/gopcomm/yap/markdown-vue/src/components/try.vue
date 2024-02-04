<script setup>

</script>
<template>

<div id="markdown-container"></div>
</template>


<script>
import Cherry from 'cherry-markdown'
import axios from 'axios'


axios.defaults.baseURL = 'http://localhost:8080/';

import 'cherry-markdown/dist/cherry-markdown.min.css'
export default {
    name: "communityEditor",
    props: {

    },
    data() {
        return {
            content: null,
            cherrInstance: null,
            child_con: 'Hellll',
            content_id: -1,
            isAdd: false,
            input_title: "",
            req_suc: false,
            open: false,
            formState: {
                resource: '',
                tags: '',
                abs: '',

            },
            resources: [
                "original",
                "reposted",
                "translated from the other ariticle"
            ],
            labelCol: { span: 10 },
            wrapperCol: { span: 14 },       
            // 自定义 
            CustomHookA: Cherry.createSyntaxHook('important', Cherry.constants.HOOKS_TYPE_LIST.SEN, {
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
        }
    },
    mounted() {
        this.initCherryMD()
        this.init_req()

        let t = "https://goplus.org/_next/static/widgets/code.85827e18ab6a0fa63bdc.js"
        let e = "https://goplus.org/_next/static/widgets/code.f81abac15122c88e65d7.css"
        var o=document.createElement("script");
        o.src=t,null!=e&&o.setAttribute("data-style-url",e),document.body.appendChild(o)
        
    },
    methods: {
        // 点击事件
        submit_article() {
            // this.open = true
            // this.submit_markdown()
        },
        // 初始化请求
        init_req() {
            // this.setMarkdown("hello")
            let url_len = window.location.href.split("/")
            if (url_len.length > 4) {
                // ID 不为空 编辑页面
                let id = url_len[4]
                axios({
                    method: 'get',
                    url: '/getArticle/'+id,
                    params: {
                        // id: id
                    }
                    }).then(response => {
                        let article = response.data.data
                        console.log(article)
                        if(article) {
                            // 是编辑 要设定文章 id 与 内容     
                            this.setMarkdown(article.content)
                            this.content_id = id
                            this.isAdd = false
                        }                      
                })
            } else {
                // 为空 新建页面
                this.isAdd = true
            }
            
        },
        
        submit_markdown() {
            let formData = new FormData()
            let title = this.getToc()
            if (title.length > 0) {
                title = title[0].text
            } else {
                title = "Default Title"
            }

            if(this.input_title != "") {
                title = this.input_title
            }

            formData.append("title",title)
            formData.append("content",this.getCherryContent())
            formData.append("html",this.getCherryHtml())


            let trans = ""
            
            if (this.content_id > 0) {
                trans = "translate"
                formData.append("id",this.content_id)
            }
            // 添加内容 标题修改 ----hxy
            
            let r = this.resources[this.formState.resource]
            console.log("resource", r)
            // 结束
            formData.append("trans", r)
            // 上传了一张图片
            formData.append("cover", "")
            formData.append("tags",this.formState.tags)
            formData.append("abstract",this.formState.abs)


            // console.log("submit", this.getCherryHtml())
            // commit接口
            // let data = {
            //     title: 'Default Title',
            //     content: this.getCherryContent(),
            //     html_content: this.getCherryHtml(),
            //     trans: trans,
            //     cover: "",
            //     tags: "",
            //     id: this.content_id
            // }
            axios({
                method: 'post',
                // url: '/api/commit', // 后端接口地址，需要根据实际情况修改
                url: '/commit', // 后端接口地址，需要根据实际情况修改
                data: formData,
                headers: {
                }
            })
                .then(response => {
                    console.log('内容发送成功');
                    // console.log(response.data); // 输出服务器返回的数据（可选） 获取 返回的ID

                    // let id = response.data.article.id
                    // let md = response.data.article.content
                    // alert("文章提交成功")
                    this.req_suc = true
                    setTimeout(() => {
                        this.req_suc = false
                    }, 2000)
                    window.location.href="http://localhost:8080/p/" + response.data.data; 
                })
                .catch(error => {
                    console.error('内容发送失败');
                    console.error(error);
                });

        },
        tranlate_md() {
            let formData = new FormData()
            formData.append("content",this.getCherryContent())
            formData.append("html",this.getCherryHtml())

            if (this.content_id > 0) {
                formData.append("id",this.content_id) 
            }
            axios({
                method: 'post',
                url: '/translate', // 后端接口地址，需要根据实际情况修改
                // url: '/api/translate', // 后端接口地址，需要根据实际情况修改
                data: formData,
                headers: {
                }
            })
                .then(response => {
                    console.log('内容发送成功');
                    console.log(response.data.id); // 输出服务器返回的数据（可选） 获取 返回的ID
                    let trans_content = response.data.data
                    console.log(trans_content)
                    this.setMarkdown(trans_content.content)
                    this.content_id = trans_content.id
                })
                .catch(error => {
                    console.error('内容发送失败');
                    console.error(error);
                });
        },  

        //获取cookie
        getcookie(sname){
            var acookie = document.cookie.split("; ");
            for (var i = 0; i < acookie.length; i++) {
                var arr = acookie[i].split("=");
                if (sname == arr[0]) {
                    if (arr.length > 1)
                        return unescape(arr[1]);
                    else
                        return "";
                }
            }
            return "";
        },
        // 初始化编辑器
        initCherryMD(value, config) {
            var {
                afterChange,
                afterInit,
                beforeImageMounted,
                fileUpload,
                // mdId
            } = this
            // var defaultValue = value || this.value
            var defaultValue = value || ""

            // var defaultValue = `Hello`
            // 自定义工具栏 - 提交按钮
            // var customMenuC = Cherry.createMenuHook('Submit', {
            //     noIcon: true,
            //     // iconName: 'preview',
            //     name: 'Submit',
            //     onClick: (selection) => {
            //         console.log("点击")
            //         // 调用 提交函数
            //         this.submit_markdown()

            //     },
            // }); 
            // 自定义工具栏 - 翻译按钮
            var customMenuF = Cherry.createMenuHook('Translate', {
                noIcon: true,
                iconName: true,
                name: 'Translate',
                onClick: (selection) => {
                    console.log("点击翻译")
                    // 调用 提交函数
                    this.tranlate_md()

                },
            }); 
            //  测试数据
            // !video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)

            var myBlockHook = Cherry.createSyntaxHook('myBlock', Cherry.constants.HOOKS_TYPE_LIST.PAR, {
                makeHtml(str) {
                    return str.replace(this.RULE.reg, function(whole, m1) {
                        console.log("html", whole, m1)
                        // 匹配whole括号中的内容
                        const regex = /\((.*?)\)/; // 定义正则表达式，其中 \() 表示左括号，(.*?) 表示非贪婪模式匹配任意字符，\)) 表示右括号
                        const matchResult = whole.match(regex); // 使用 match 函数进行匹配
                        let video_src = ""
                        if (matchResult && matchResult[1]) {
                            console.log("提取到的结果为：" + matchResult[1]);
                            video_src = matchResult[1]
                            return `<div><video controls crossorigin playsinline data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" id="player"><source src=${video_src} type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default/><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"/><a href=${video_src} download>Download</a></video></div>`
                        } else {
                            console.log("未能提取到小括号中的内容");
                            let r = "未能提取到小括号中的内容"
                            return `<div style="border: 1px solid;border-radius: 15px;background: gold;">${r}</div>`;
                            // return `<div><video controls crossorigin playsinline data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" id="player"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default/><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"/><a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download>Download</a></video></div>`
                        }
                    });
                },
                rule(str) {
                    // return { reg: /\n\+\+(\n[\s\S]+?\n)\+\+\n/g };
                    return {
                        reg: /^!video.*.mp4\)/
                    }
                },
            });
            this.cherrInstance = new Cherry({
                id: 'markdown-container',                    
                value: defaultValue,
                customSyntax: {
                    // importHook: {
                    //     syntaxClass: this.CustomHookA, // 将自定义语法对象挂载到 importHook.syntaxClass上
                    //     force: true, // true： 当cherry自带的语法中也有一个“importHook”时，用自定义的语法覆盖默认语法； false：不覆盖
                    //     before: 'fontEmphasis', // 定义该自定义语法的执行顺序，当前例子表明在加粗/斜体语法前执行该自定义语法
                    // },
                    myBlock: {
                        syntaxClass: myBlockHook,
                        force: true,
                        before: 'blockquote',
                    },
                },
                fileUpload: this.fileUpload,
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
                        urlProcessor: this.urlProcessor,
        
                        /**
                         * 额外允许渲染的html标签
                         * 标签以英文竖线分隔，如：htmlWhiteList: 'iframe|script|style'
                         * 默认为空，默认允许渲染的html见src/utils/sanitize.js whiteList 属性
                         * 需要注意：
                         *    - 启用iframe、script等标签后，会产生xss注入，请根据实际场景判断是否需要启用
                         *    - 一般编辑权限可控的场景（如api文档系统）可以允许iframe、script等标签
                         */
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
                                        return `<goplus-code half-code language="gop" style="width: 85vw">${src}</goplus-code> `;
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
                        toc: {
                            /** 默认只渲染一个目录 */
                            allowMultiToc: false
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
                    toolbarRight: ['fullScreen', '|', 'customMenuFName'],
                    bubble: ['bold', 'italic', 'underline', 'strikethrough', 'sub', 'sup', 'quote', 'ruby', '|', 'size',
                        'color'
                    ],
                    // array or false
                    "float": ['h1', 'h2', 'h3', '|', 'checklist', 'quote', 'quickTable', 'code'], // array or false
                    sidebar: ['mobilePreview', 'theme'], // 'copy',
                    customMenu: {
                        // customMenuAName: customMenuA,
                        // customMenuBName: customMenuB,
                        // customMenuCName: customMenuC,
                        customMenuFName: customMenuF
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
                    afterChange: this.afterChange,
                    afterInit: this.afterInit,
                    beforeImageMounted: this.beforeImageMounted,
                    // 预览区域点击事件，previewer.enablePreviewerBubble = true 时生效
                    onClickPreview: this.onClickPreview,
                    // 复制代码块代码时的回调
                    onCopyCode: this.onCopyCode,
                    // 把中文变成拼音的回调，当然也可以把中文变成英文、英文变成中文
                    // changeString2Pinyin: this.changeString2Pinyin
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
        },
        // 上传通用接口
        fileUpload(file) {
            // 提交的时候代码
            let type = file.type
            console.log("upload", file.type)
            var formData = new FormData()
            formData.append('file', file)
            // axios post 请求                
            axios({
                method: 'post',
                url: '/upload', // 后端接口地址，需要根据实际情况修改
                data: formData,
                headers: {
                    'Content-Type': 'multipart/form-data' // 设置请求头部类型为 multipart/form-data
                }
            })
                .then(response => {
                    console.log('文件上传成功！');
                    console.log(response.data); // 输出服务器返回的数据（可选） 获取 返回的ID
                    axios({
                        method: 'get',
                        url: '/getMediaUrl/'+response.data,
                        // params: {
                        //     id: media_id
                        // }
                    }).then(response => {
                        console.log("请求media成功")
                        console.log()
                        let url = response.data.url
                        console.log(url)
                        // 返回文件地址的话需要
                        this.onloadCallback(type, url)                        
                    })
                })
                .catch(error => {
                    console.error('文件上传失败！');
                    console.error(error);
                });

            // // 本地测试代码
            // console.log("type", file.type)
            // let type = file.type
            // let url = 'https://img.mp.itc.cn/upload/20161030/0bf4745fadf046b3a68ece0763fc3010_th.jpeg'
            // this.onloadCallback(type, url)                        

        },
        onloadCallback(type, url) {
            console.log("call back")
            let imgMdStr = ""
            if (/mp4|avi|rmvb/i.test(type)) {
                // 会渲染成video标签 需要替换video1标签 <video controls src=""> <video>
                imgMdStr = `!video[${type}](${url})`;
            } else if (/mp3/i.test(type)) {
                imgMdStr = `!audio[${type}](${url})`;
            } else if (/bmp|gif|jpg|jpeg|png/i.test(type)) {
                imgMdStr = `![${type}](${url})`
            } else {
                imgMdStr = `[${type}](${url})`
            } 
            this.cherrInstance.insert(imgMdStr)
            
            // test 代码 成功
            // console.log("call back")
            // let imgMdStr = ''
            // imgMdStr = `![${'image/jpg'}](https://up.deskcity.org/pic_source/d6/6a/c5/d66ac5159da4a416a910061f3df99add.jpg)`;
            
            // this.cherrInstance.insert(imgMdStr)
        },
        // 全局的URL处理器
        urlProcessor(url, srcType) {
            return url;
        },
        // 变更事件回调
        afterChange(text, html) {
            this.content = text
            // console.log("markdown", this.getCherryContent())
            // console.log("html", this.getCherryHtml())
            this.$emit('mdChange', html, text)
            this.$emit('input', text)
        },
        // 初始化事件回调
        afterInit(e) {},
        // 图片加载回调
        beforeImageMounted(e, src) {
            return {
                [e]: src
            }
        },
        // 预览区域点击事件
        onClickPreview(event) {},
        // 粘贴事件
        onCopyCode(event, code) {
            // 阻止默认的粘贴事件
            // return false;
            // 对复制内容进行额外处理
            return code;
        },
        // 获取中文的拼音
        // changeString2Pinyin(string) {
        //     /**
        //      * 推荐使用这个组件：https://github.com/liu11hao11/pinyin_js
        //      *
        //      * 可以在 ../scripts/pinyin/pinyin_dist.js 里直接引用
        //      */
        //     var pinyin = require("./pinyin/pinyin.js");
        //     return pinyin.pinyin(string, " ");
        // },
        setMarkdown(content) {
            console.log("init req", content)
            if (!this.cherrInstance) { // 未加载则重新初始化
                this.initCherryMD(content)
                return
            }
            // this.cherrInstance.setMarkdown(content, 0)
            this.cherrInstance.insert(content, false, false, true)

        },
        // setValue(content, keepCursor = false) {
        //     let editor = this.cherrInstance.editor
        //     editor.storeDocumentScroll();
        //     if (keepCursor === false) {
        //     return editor.editor.setValue(content);
        //     }
        //     const codemirror = editor.editor;
        //     const old = this.getCherryContent();
        //     const pos = codemirror.getDoc().indexFromPos(codemirror.getCursor());
        //     const newPos = getPosBydiffs(pos, old, content);
        //     codemirror.setValue(content);
        //     const cursor = codemirror.getDoc().posFromIndex(newPos);
        //     codemirror.setCursor(cursor);
        // },
        getCherryContent() {
            var result = this.cherrInstance.getMarkdown() // 获取markdown内容
            return result
        },
        getCherryHtml() {
            var result = this.cherrInstance.getHtml()

            //   添加成完整的html 以及 script标签 
            var parser = new DOMParser();
            var doc = parser.parseFromString(result, 'text/html');
            // console.log(doc)

            let t = "https://goplus.org/_next/static/widgets/code.85827e18ab6a0fa63bdc.js"
            let e = "https://goplus.org/_next/static/widgets/code.f81abac15122c88e65d7.css"
            var o=document.createElement("script");
            o.src=t,null!=e&&o.setAttribute("data-style-url",e),document.body.appendChild(o)
        
            // var s = document.createElement('script'); 
            // s.type = 'text/javascript'; 
            // var code = 'alert("hello world!");'; 
            // s.appendChild(document.createTextNode(code)); 
            doc.body.appendChild(o);
            let new_content = doc.body.innerHTML
            console.log(doc)
            console.log(new_content)
            return new_content
        },
        getData() {
            var result = this.cherrInstance.getHtml()
            return result
        },
        getToc() {
            var result = this.cherrInstance.getToc()
            return result
        },
        /**
         * type：{'pdf'|'img'}
         */
        exportMD(type = 'pdf') {
            this.cherrInstance.export(type)
        },
        /**
         * model{'edit&preview'|'editOnly'|'previewOnly'}
         */
        switchModel(model) {
            if (this.isInit()) {
                this.cherrInstance.switchModel(model)
            }
        },

        insert(content, isSelect = false, anchor = [], focus = true) {
            this.cherrInstance.insert(content, isSelect, anchor, focus)
        },
        isInit() {
            if (this.cherrInstance) {
                return true
            }
            this.$message.warning('编辑器未初始化，请检查')
            return false
        },
    }
}
  </script>
  
  <style scoped>
  #markdown-container {
    position: absolute;
    left: 0px;
    top: 0px;
    width: 100%;
  }
  </style>