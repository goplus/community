<script setup>

</script>

<script>
import Plyr from 'plyr';
import 'plyr/dist/plyr.css'

var cherrInstance = null

var vtt_src = ""

let fileType = "video/mp4"


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
                // reg: /^!video.*.mp4\)/
                // reg: /^!video.*\)/
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
        externals: { // externals
        },
        // 解析引擎配置
        engine: {
            global: {
                urlProcessor(url, srcType) {
                    console.log(`url-processor`, url, srcType);
                    return url;
                },
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
        toolbars: {
            toolbar: false,
        },
        editor: {
            defaultModel: 'previewOnly',
        },
        previewer: {
            // 自定义markdown预览区域class
            className: 'viewer'
        },
        keydown: [],
        // 外层容器不存在时，是否强制输出到body上
        forceAppend: true,
        // The locale Cherry is going to use. Locales live in /src/locales/
        locale: "en_US",
    })

}

export default {
        name: "MarkdownViewer",
        props: {
            "md": String    
        },
        components: {
            
        },
        data() {
        },
        watch: {
            md(newValue) {
                cherrInstance.setMarkdown(newValue)
            }
        },
        beforeMount() {
           initCherryMD(this.md)
        },
        mounted() {
        //    this.insert()

        },
        methods: {  
        }
}

</script>
  
 <style>
    #markdown {
      border: none;
      background: #ffffff;
    }
    .cherry.cherry--no-toolbar .cherry-editor, .cherry.cherry--no-toolbar .cherry-previewer {
        background: #ffffff;
        border: none;
        box-shadow: 0 0px 0px #ffffff;
    }
    .cherry{
        box-shadow: none;
    }
</style>