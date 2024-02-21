<script setup>

</script>
<template>
    <div id="markdown-container"></div>
</template>
<script>
import Cherry from 'cherry-markdown'
import axios from 'axios'


import Plyr from 'plyr';
import 'plyr/dist/plyr.css'
// const player = new Plyr('#player');
// const p = new Plyr('video', {captions: {active: true}});


// let v_s = "https://cdn.plyr.io/3.6.8/plyr.js"
//             let v_css = "https://cdn.plyr.io/3.6.8/plyr.css"
// import {  UploadOutlined, SendOutlined } from '@ant-design/icons-vue';


axios.defaults.baseURL = '/';

var cherrInstance = null

var cherryEngine = null

var vtt_src = ""

function fileUpload(file) {
    // 提交的时候代码
    let type = file.type
    console.log("upload", file.type)
    var formData = new FormData()
    formData.append('file', file)
    // axios post 请求                
    axios({
        method: 'post',
        url: '/api/medias', // 后端接口地址，需要根据实际情况修改
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data' // 设置请求头部类型为 multipart/form-data
        }
    })
        .then(response => {
            console.log('文件上传成功！', type);
            console.log(response.data); // 输出服务器返回的数据（可选） 获取 返回的ID  
            // 根据类型执行不同的请求
            if(type.includes("video")) {
                axios({
                    method: 'get',
                    url: '/api/video-with-caption/'+response.data,
                }).then(response => {
                    console.log("请求视频")
                    let url = response.data.url.fileKey
                    let stitle = response.data.url.subtitle
                    vtt_src = stitle
                    console.log("视频URL", url)
                    // 返回文件地址的话需要
                    onloadCallback(type, url)                        
                })
            } else {
                axios({
                    method: 'get',
                    url: '/api/medias/url/'+response.data,
                }).then(response => {
                    console.log("请求media成功")
                    let url = response.data.url
                    console.log("图片", url)
                    // 返回文件地址的话需要
                    onloadCallback(type, url)                        
                })
            }
           
        })
        .catch(error => {
            console.error('文件上传失败！');
            console.error(error);
        });

    // // 本地测试代码
    // console.log("type", file.type)
    // let type = file.type
    // let url = 'https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4'
    // vtt_src = "https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt"
    // onloadCallback(type, url)                        

}
function onloadCallback(type, url) {
    console.log("call back")
    let imgMdStr = ""
    if (/mp4|avi|rmvb/i.test(type)) {
        // 会渲染成video标签 需要替换video1标签 <video controls src=""> <video>
        imgMdStr = `!video[${type}](${url})(${vtt_src})`;
    } else if (/mp3/i.test(type)) {
        imgMdStr = `!audio[${type}](${url})`;
    } else if (/bmp|gif|jpg|jpeg|png/i.test(type)) {
        imgMdStr = `![${type}](${url})`
    } else {
        imgMdStr = `[${type}](${url})`
    } 
    cherrInstance.insert(imgMdStr) 
    // test 代码 成功
    // console.log("call back")
    // let imgMdStr = ''
    // imgMdStr = `![${'image/jpg'}](https://up.deskcity.org/pic_source/d6/6a/c5/d66ac5159da4a416a910061f3df99add.jpg)`;
    
    // cherrInstance.insert(imgMdStr)
}
            // 全局的URL处理器
function urlProcessor(url, srcType) {
    return url;
}
        
function afterChange(text, html) {
    this.content = text

    console.log("markdown", getCherryContent())
    console.log("html", getCherryHtml())
    const p = new Plyr('video', {captions: {active: true}});

    // this.$emit('mdChange', html, text)
    // this.$emit('input', text)
}

function onCopyCode(event, code) {
    // 阻止默认的粘贴事件
    // return false;
    // 对复制内容进行额外处理
    return code;
}
// 获取中文的拼音
function changeString2Pinyin(string) {
    /**
     * 推荐使用这个组件：https://github.com/liu11hao11/pinyin_js
     *
     * 可以在 ../scripts/pinyin/pinyin_dist.js 里直接引用
     */
    var pinyin = require("./pinyin/pinyin.js");
    return pinyin.pinyin(string, " ");
}
function setMarkdown(content) {
    console.log("init req", content)
    if (!cherrInstance) { // 未加载则重新初始化
        initCherryMD(content)
        return
    }
    // this.cherrInstance.setMarkdown(content, 0)
    cherrInstance.insert(content, false, false, true)

}

function getCherryContent() {
    if (!cherrInstance) { // 未加载则重新初始化
        initCherryMD("")
        // return
    }
    var result = cherrInstance.getMarkdown() // 获取markdown内容
    return result
}
function getCherryHtml() {
    var result = cherrInstance.getHtml()
    //   添加成完整的html 以及 script标签 
    // var parser = new DOMParser();
    // var doc = parser.parseFromString(result, 'text/html');
    // // console.log(doc)

    // let t = "https://goplus.org/_next/static/widgets/code.85827e18ab6a0fa63bdc.js"
    // let e = "https://goplus.org/_next/static/widgets/code.f81abac15122c88e65d7.css"
    // var o=document.createElement("script");
    // o.src=t,null!=e&&o.setAttribute("data-style-url",e),document.body.appendChild(o) 
    // doc.body.appendChild(o);
    // let new_content = doc.body.innerHTML
    // console.log(doc)
    // console.log(new_content)
    return result
}
// 
function makeHtml(str) {
    var result = cherrInstance.engine.makeHtml(str)
    console.log("make html", result)
    return result
}
function getData() {
    var result = cherrInstance.getHtml()
    return result
}
function getToc() {
    var result = cherrInstance.getToc()
    return result
}

function exportMD(type = 'pdf') {
    cherrInstance.export(type)
}
function switchModel(model) {
    if (this.isInit()) {
        cherrInstance.switchModel(model)
    }
}

function insert(content, isSelect = false, anchor = [], focus = true) {
    cherrInstance.insert(content, isSelect, anchor, focus)
}
function isInit() {
    if (cherrInstance) {
        return true
    }
    this.$message.warning('编辑器未初始化，请检查')
    return false
}

function initCherryMD(value, config) {
    // var {
    //     afterChange,
    //     afterInit,
    //     beforeImageMounted,
    //     fileUpload,
    //     // mdId
    // } = this
    var defaultValue = value || ""

    // 自定义工具栏 - 翻译按钮
    // var customMenuF = Cherry.createMenuHook('Translate', {
    //     noIcon: true,
    //     iconName: true,
    //     name: 'Translate',
    //     onClick: (selection) => {
    //         console.log("点击翻译")
    //         // 调用 提交函数
    //         this.tranlate_md()
    //     },
    // }); 
    //  测试数据
    // !video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)

    var myBlockHook = Cherry.createSyntaxHook('myBlock', Cherry.constants.HOOKS_TYPE_LIST.PAR, {
        makeHtml(str) {
            return str.replace(this.RULE.reg, function(whole, m1) {
                // 匹配whole括号中的内容
                console.log("提取video内容", whole)
                const regex = /\((.*?)\)/; // 定义正则表达式，其中 \() 表示左括号，(.*?) 表示非贪婪模式匹配任意字符，\)) 表示右括号
                const matchResult = whole.match(regex); // 使用 match 函数进行匹配
                let video_src = ""
                if (matchResult && matchResult[1]) {
                    console.log("提取到的结果为：" + matchResult[1]);
                    // console.log("player", player)
                    video_src = matchResult[1]
                    console.log("video src 已经替换了", video_src)
                    return `<div><video controls="" crossorigin="" playsinline="" poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg"><source src=${video_src} type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src=${vtt_src} default/><track kind="captions" label="Français" srclang="fr" src=${vtt_src} /><a href=${video_src} download>Download</a></video></div>`
                    // return `<vue-plyr ref="plyr"><video ref="plyr" class="plyr" :autoplay="autoplay" :crossorigin="crossorigin" :poster="poster" :playsinline="playsinline"><source :src="source" type="video/mp4"></video></vue-plyr>`
                    // return `<div class="plyr plyr--full-ui plyr--video plyr--html5 plyr--fullscreen-enabled plyr--paused plyr--stopped plyr--pip-supported plyr--captions-enabled plyr__poster-enabled"><video controls crossorigin playsinline poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-720p.mp4" type="video/mp4" size="720"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-1080p.mp4" type="video/mp4" size="1080"><!-- Caption files --><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt"default><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"><!-- Fallback for browsers that don't support the <video> element --><a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download>Download</a></video>`
                    // return `<div class="plyr plyr--full-ui plyr--video plyr--html5 plyr--fullscreen-enabled plyr--paused plyr--stopped plyr--pip-supported plyr--captions-enabled plyr__poster-enabled"><div class="plyr__controls"><button class="plyr__controls__item plyr__control" type="button" data-plyr="play" aria-pressed="false" aria-label="Play"><svg class="icon--pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-pause"></use></svg><svg class="icon--not-pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-play"></use></svg><span class="label--pressed plyr__sr-only">Pause</span><span class="label--not-pressed plyr__sr-only">Play</span></button><div class="plyr__controls__item plyr__progress__container"><div class="plyr__progress"><input data-plyr="seek" type="range" min="0" max="100" step="0.01" value="0" autocomplete="off" role="slider" aria-label="Seek" aria-valuemin="0" aria-valuemax="183.125333" aria-valuenow="0" id="plyr-seek-8215" aria-valuetext="00:00 of 03:03" style="--value: 0%;"><progress class="plyr__progress__buffer" min="0" max="100" value="0" role="progressbar" aria-hidden="true">% buffered</progress><span class="plyr__tooltip">00:00</span></div></div><div class="plyr__controls__item plyr__time--current plyr__time" aria-label="Current time" role="timer">-03:03</div><div class="plyr__controls__item plyr__volume"><button type="button" class="plyr__control" data-plyr="mute" aria-pressed="false"><svg class="icon--pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-muted"></use></svg><svg class="icon--not-pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-volume"></use></svg><span class="label--pressed plyr__sr-only">Unmute</span><span class="label--not-pressed plyr__sr-only">Mute</span></button><input data-plyr="volume" type="range" min="0" max="1" step="0.05" value="1" autocomplete="off" role="slider" aria-label="Volume" aria-valuemin="0" aria-valuemax="100" aria-valuenow="100" id="plyr-volume-8215" aria-valuetext="100.0%" style="--value: 100%;"></div><button class="plyr__controls__item plyr__control" type="button" data-plyr="captions" aria-pressed="false"><svg class="icon--pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-captions-on"></use></svg><svg class="icon--not-pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-captions-off"></use></svg><span class="label--pressed plyr__sr-only">Disable captions</span><span class="label--not-pressed plyr__sr-only">Enable captions</span></button><div class="plyr__controls__item plyr__menu"><button aria-haspopup="true" aria-controls="plyr-settings-8215" aria-expanded="false" type="button" class="plyr__control" data-plyr="settings" aria-pressed="false"><svg aria-hidden="true" focusable="false"><use xlink:href="#plyr-settings"></use></svg><span class="plyr__sr-only">Settings</span></button><div class="plyr__menu__container" id="plyr-settings-8215" hidden=""><div><div id="plyr-settings-8215-home"><div role="menu"><button data-plyr="settings" type="button" class="plyr__control plyr__control--forward" role="menuitem" aria-haspopup="true"><span>Captions<span class="plyr__menu__value">Disabled</span></span></button><button data-plyr="settings" type="button" class="plyr__control plyr__control--forward" role="menuitem" aria-haspopup="true"><span>Quality<span class="plyr__menu__value">576p</span></span></button><button data-plyr="settings" type="button" class="plyr__control plyr__control--forward" role="menuitem" aria-haspopup="true"><span>Speed<span class="plyr__menu__value">Normal</span></span></button></div></div><div id="plyr-settings-8215-captions" hidden=""><button type="button" class="plyr__control plyr__control--back"><span aria-hidden="true">Captions</span><span class="plyr__sr-only">Go back to previous menu</span></button><div role="menu"><button data-plyr="language" type="button" role="menuitemradio" class="plyr__control" aria-checked="true" value="-1"><span>Disabled</span></button><button data-plyr="language" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="0"><span>English<span class="plyr__menu__value"><span class="plyr__badge">EN</span></span></span></button><button data-plyr="language" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="1"><span>Français<span class="plyr__menu__value"><span class="plyr__badge">FR</span></span></span></button></div></div><div id="plyr-settings-8215-quality" hidden=""><button type="button" class="plyr__control plyr__control--back"><span aria-hidden="true">Quality</span><span class="plyr__sr-only">Go back to previous menu</span></button><div role="menu"><button data-plyr="quality" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="1080"><span>1080p<span class="plyr__menu__value"><span class="plyr__badge">HD</span></span></span></button><button data-plyr="quality" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="720"><span>720p<span class="plyr__menu__value"><span class="plyr__badge">HD</span></span></span></button><button data-plyr="quality" type="button" role="menuitemradio" class="plyr__control" aria-checked="true" value="576"><span>576p<span class="plyr__menu__value"><span class="plyr__badge">SD</span></span></span></button></div></div><div id="plyr-settings-8215-speed" hidden=""><button type="button" class="plyr__control plyr__control--back"><span aria-hidden="true">Speed</span><span class="plyr__sr-only">Go back to previous menu</span></button><div role="menu"><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="0.5"><span>0.5×</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="0.75"><span>0.75×</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="true" value="1"><span>Normal</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="1.25"><span>1.25×</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="1.5"><span>1.5×</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="1.75"><span>1.75×</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="2"><span>2×</span></button><button data-plyr="speed" type="button" role="menuitemradio" class="plyr__control" aria-checked="false" value="4"><span>4×</span></button></div></div></div></div></div><button class="plyr__controls__item plyr__control" type="button" data-plyr="pip" aria-pressed="false"><svg aria-hidden="true" focusable="false"><use xlink:href="#plyr-pip"></use></svg><span class="plyr__sr-only">PIP</span></button><button class="plyr__controls__item plyr__control" type="button" data-plyr="fullscreen" aria-pressed="false"><svg class="icon--pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-exit-fullscreen"></use></svg><svg class="icon--not-pressed" aria-hidden="true" focusable="false"><use xlink:href="#plyr-enter-fullscreen"></use></svg><span class="label--pressed plyr__sr-only">Exit fullscreen</span><span class="label--not-pressed plyr__sr-only">Enter fullscreen</span></button></div><div class="plyr__video-wrapper"><video crossorigin="" playsinline="" poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-720p.mp4" type="video/mp4" size="720"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-1080p.mp4" type="video/mp4" size="1080"><!-- Caption files --><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default=""><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"><!-- Fallback for browsers that don't support the <video> element --><a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download="">Download</a></video><div class="plyr__poster" style="background-image: url(&quot;https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg&quot;);"></div></div><div class="plyr__captions" dir="auto"></div><button type="button" class="plyr__control plyr__control--overlaid" data-plyr="play" aria-pressed="false" aria-label="Play"><svg aria-hidden="true" focusable="false"><use xlink:href="#plyr-play"></use></svg><span class="plyr__sr-only">Play</span></button></div>`
                } else {
                    console.log("未能提取到小括号中的内容");
                    let r = "未能提取到小括号中的内容"
                    return `<div style="border: 1px solid;border-radius: 15px;background: gold;">${r}</div>`;
                    // return `<div><video controls crossorigin playsinline data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" id="player"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default/><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"/><a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download>Download</a></video></div>`
                }
            });
        },
        
        rule(str) {
            return {
                // reg: /^!video.*.mp4\)/
                reg: /^!video.*\)/
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
            // toolbarRight: ['fullScreen', '|', 'customMenuFName'],
            toolbarRight: ['fullScreen'],
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
    // import js
    // import 'https://cdn.plyr.io/3.6.8/plyr.js'
    // import "https://goplus.org/_next/static/widgets/code.85827e18ab6a0fa63bdc.js"
export default {
    name: "GoplusMarkdown",
    props: {

    },
    components: {
        // UploadOutlined,
        // SendOutlined
    },
    data() {
        return {
            // content: null,
            // cherrInstance: null,
            // child_con: 'Hellll',
            content_id: -1,
            isAdd: false,
            // input_title: "",
            // req_suc: false,
            // open: false,
            // formState: {
            //     resource: '',
            //     tags: '',
            //     abs: '',

            // },
            // resources: [
            //     "original",
            //     "reposted",
            //     "translated from the other ariticle"
            // ],
            // labelCol: { span: 10 },
            // wrapperCol: { span: 14 },       
            // 自定义 
            // CustomHookA: Cherry.createSyntaxHook('important', Cherry.constants.HOOKS_TYPE_LIST.SEN, {
            //     makeHtml(str) {
            //         console.log("custom hook", str)
            //         return str.replace(this.RULE.reg, function(whole, m1, m2) {
            //             return `<span style="color: green;"><strong>${m2}</strong></span>`;
            //         });
            //     },
            //     rule(str) {
            //         console.log("rule", str)
            //         return { reg: /(\*\*\*)([^\*]+)\1/g };
            //     },
            // })
        }
    },
    mounted() {
        initCherryMD()
        this.init_req()
        // const p = new Plyr('video', {captions: {active: true}});

        // 不需要
        // let t = "https://goplus.org/_next/static/widgets/code.85827e18ab6a0fa63bdc.js"
        // let e = "https://goplus.org/_next/static/widgets/code.f81abac15122c88e65d7.css"
        // var o=document.createElement("script");
        // o.src=t,null!=e&&o.setAttribute("data-style-url",e),document.body.appendChild(o)


        // let v_s = "https://cdn.plyr.io/3.6.8/plyr.js"
        // let v_css = "https://cdn.plyr.io/3.6.8/plyr.css"
        // var v=document.createElement("script");
        // v.src=v_s,null!=v_css&&v.setAttribute("data-style-url",v_css),document.body.appendChild(v)

        // let v_c = document.createElement("link")
        // v_c.href = v_css
        // v_c.setAttribute("href", "stylesheet")
        // document.head.appendChild(v_c)
    },
    methods: {    
         getMarkdown() {
            return getCherryContent()
        },         
        // 弹窗关闭事件
        // handleOk(e) {
        //     console.log(e);
        //     console.log(this.formState)
        //     // 提交
        //     this.submit_markdown()
        //     this.open = false

        // },
        
        // // 点击事件
        // submit_article() {
        //     this.open = true
        // },
        // 初始化请求
        init_req() {
            // this.setMarkdown("hello")
            let url_len = window.location.href.split("/")
            if (url_len.length > 4) {
                // ID 不为空 编辑页面
                let id = url_len[4]
                axios({
                    method: 'get',
                    url: '/api/article/'+id,
                    params: {
                        // id: id
                    }
                    }).then(response => {
                        let article = response.data.data
                        console.log(article)
                        if(article) {
                            // 是编辑 要设定文章 id 与 内容     
                            setMarkdown(article.content)
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
            let title = getToc()
            if (title.length > 0) {
                title = title[0].text
            } else {
                title = "Default Title"
            }

            if(this.input_title != "") {
                title = this.input_title
            }

            formData.append("title",title)
            formData.append("content",getCherryContent())
            formData.append("html",getCherryHtml())


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

            axios({
                method: 'post',
                // url: '/api/commit', // 后端接口地址，需要根据实际情况修改
                url: '/api/commit', // 后端接口地址，需要根据实际情况修改
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
                    window.location.href="/page/p/" + response.data.data; 
                })
                .catch(error => {
                    console.error('内容发送失败');
                    console.error(error);
                });

        },
        tranlate_md() {
            let formData = new FormData()
            formData.append("content",getCherryContent())
            formData.append("html",getCherryHtml())

            if (this.content_id > 0) {
                formData.append("id",this.content_id) 
            }
            axios({
                method: 'post',
                url: '/translations', // 后端接口地址，需要根据实际情况修改
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
                    setMarkdown(trans_content.content)
                    this.content_id = trans_content.id
                })
                .catch(error => {
                    console.error('内容发送失败');
                    console.error(error);
                });
        },  

        setSubmit(sub) {
            sub()
        },
        setTrans(tran) {
            tran()
        },
        insertMarkdown(content) {
            setMarkdown(content)
        },  
        setMarkdown(content) {
            // setMarkdown(content)
            cherrInstance.setMarkdown(content)
            cherrInstance.editor.editor.setCursor(cherrInstance.getMarkdown().length)
        }

        //获取cookie
        // getcookie(sname){
        //     var acookie = document.cookie.split("; ");
        //     for (var i = 0; i < acookie.length; i++) {
        //         var arr = acookie[i].split("=");
        //         if (sname == arr[0]) {
        //             if (arr.length > 1)
        //                 return unescape(arr[1]);
        //             else
        //                 return "";
        //         }
        //     }
        //     return "";
        // },
    }
}

// const MarkdownEditor = {
//         getHtml: function() {
//             console.log("Get Html")
//             return getCherryHtml()
//         },
//         getMarkdown: function() {
//             console.log("Get Markdown")
//             return getCherryContent()
//         },
//         setInitValue: function(content) {
//             setMarkdown()
//         },
//         setTranslation: function(trans) {
//             trans()
//         },
//         setSubmit: function(sub) {
//             sub()
//         }
// }
// export {MarkdownEditor}

</script>
  
<style>
/* @import url('https://cdn.plyr.io/3.6.8/plyr.css');     */
</style>