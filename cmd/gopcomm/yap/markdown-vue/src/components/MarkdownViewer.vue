<script setup>

</script>
<template>
    <div>
        <div v-html="content"></div>
        <!-- <div>{{ message }}</div> -->
    </div>
</template>
<script>
import CherryEngine from 'cherry-markdown/dist/cherry-markdown.engine.core'
// import Plyr from 'plyr';
// import 'plyr/dist/plyr.css'
const cherryEngineInstance = new CherryEngine();

export default {
        name: "MarkdownViewer",
        props: {
            "md": String
        },
        components: {
            
        },
        data() {
            return {
                content: "Hello World",
                // content: '!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)', // HTML内容
                vtts: [],
                videos: []
            }
        },
        watch: {
            md(newValue) {
                // 在这里处理新值的更新逻辑
                console.log("更新子组件值", newValue)
                this.computeOut(newValue)

            }
        },
        beforeMount() {
            // const cherryEngineInstance = new CherryEngine();
            // var html = this.getVttfile(this.md)
            // console.log("++++++++", html)
            // var htmlContent = cherryEngineInstance.makeHtml(html);
            // console.log("html content", htmlContent)
            // htmlContent = this.replaceVideo(htmlContent)
            // htmlContent = this.replaceGoplus(htmlContent)
            // this.content = htmlContent
            // console.log("result", htmlContent)
            this.computeOut(this.md)
        },
        mounted() {
            // console.log("mounted", this.md)
            // const cherryEngineInstance = new CherryEngine();
            // // markdown语法以 <code>gop 开头</code>结尾
            // // video <video>开头 </video>结尾 
            // // <code>gop println&quot;&quot;</code>
            // // var md = this.replaceGoplus(this.md)
            // var html = this.getVttfile(this.md)
            // console.log("++++++++", html)
            // var htmlContent = cherryEngineInstance.makeHtml(html);
            // console.log("html content", htmlContent)
            // htmlContent = this.replaceVideo(htmlContent)
            // htmlContent = this.replaceGoplus(htmlContent)
            // this.content = htmlContent
            // console.log("result", htmlContent)
            // const p = new Plyr('video', {captions: {active: true}});

        },
        methods: {  
            computeOut(md) {
                // const cherryEngineInstance = new CherryEngine();
                var html = this.getVttfile(md)
                console.log("++++++++", html)
                var htmlContent = cherryEngineInstance.makeHtml(html);
                console.log("html content", htmlContent)
                htmlContent = this.replaceVideo(htmlContent)
                htmlContent = this.replaceGoplus(htmlContent)
                this.content = htmlContent
                console.log("result", htmlContent)
            },  
            replaceVideo(htmlContent) {
                var regex = /<video(.*?)<\/video>/g
                var matches = htmlContent.match(regex); // 使用match()函数获取匹配结果
                console.log("m", matches)
                let now_s = this.videos
                let now_v = this.vtts
                if(matches) {
                    for(let i=0; i<matches.length; i++) {
                        console.log("match", matches[i])
                        // htmlContent = htmlContent.replace(matches[i], "video replace" )
                        htmlContent = htmlContent.replace(matches[i], function(h) { 
                            var regex1 = /src="(.*?)">video\/mp4/
                            let src = matches[i].match(regex1)[1]
                            let src_a = now_s[i]
                            // console.log("---------", this.videos)
                            let vtt_a = now_v[i]
                            console.log("=========duibi", src, src_a)
                            // console.log("src", src)
                            return `<div><video controls crossorigin playsinline data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" id="player"><source src=${src_a} type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src=${vtt_a} default/><track kind="captions" label="Français" srclang="fr" src=${vtt_a}/><a href=${src_a} download>Download</a></video></div>`
                        })
                    }
                    console.log("replace")
                } else {

                }
                return htmlContent
            },
            getVttfile(mdContent) {
                // let reg1 = /^!video.*\)/g
                let reg1 = /!video(.*?)\((.*?)\)\((.*?)\)/g
                let sttt = reg1.exec(mdContent)
                console.log("sttt", sttt)
                // let p_str = sttt[0]
                // console.log(p_str)
                var replace_vtt = ""
                if(sttt) {
                    for(let i = 0; i<sttt.length; i++) {
                        let p_str = sttt[i]
                        const regex = /\((.*?)\)/g; // 定义正则表达式，其中 \() 表示左括号，(.*?) 表示非贪婪模式匹配任意字符，\)) 表示右括号
                        const matchResult = p_str.match(regex); // 使用 match 函数进行匹配
                        let video_src = ""
                        if (matchResult && matchResult[1]) {
                            console.log("提取到的结果为：" , matchResult[0],"===", matchResult[1]);
                            
                            replace_vtt = mdContent.replace(matchResult[1], "")

                            matchResult[1] = matchResult[1].replace("(", "")
                            matchResult[0] = matchResult[0].replace("(", "")
                            matchResult[1] = matchResult[1].replace(")", "")
                            matchResult[0] = matchResult[0].replace(")", "")
                            this.vtts.push(matchResult[1])
                            this.videos.push(matchResult[0])
                            console.log('temp', replace_vtt)
                        } 
                    }
                    return replace_vtt
                }
                
                // 替换一下
                console.log(this.videos, this.vtts)
                return mdContent
            },
            replaceGoplus(htmlContent) {    
                // 获取 data-lang 如果是gop 就进行匹配
                var regex_l = /data-lang="(.*?)"/
                var m_lan = htmlContent.match(regex_l)
                console.log("language", m_lan)
                if(m_lan) {
                    // 如果有gop 才替换
                    if(m_lan[1] == 'gop') {
                        var regex = /<code(.*?)<\/code>/g
                        var matches = htmlContent.match(regex); // 使用match()函数获取匹配结果
                        console.log("goplus")
                        if(matches) {
                            for(let i=0; i<matches.length; i++) {
                                console.log("match", matches[i])
                                htmlContent = htmlContent.replace(matches[i], function(h) { 
                                    // var regex1 = /<code(.*?)<\/code>/
                                    var regex1 = /<code class="language-javascript wrap">(.*?)<\/code>/
                                    let src = matches[i].match(regex1)[1]
                                    console.log("src", src)
                                    var regex2 = /<span class="code-line">(.*?)<\/span>/g
                                    let s_match = src.match(regex2)
                                    var iner_src = ""
                                    if(s_match) {
                                        for(let i = 0; i<s_match.length; i++) {
                                            var regex3 = /<span class="code-line">(.*?)<\/span>/
                                            let temp_s = s_match[i].match(regex3)[1]
                                            iner_src += temp_s
                                            iner_src += '\n'
                                        }
                                    }
                                    console.log("iner", iner_src)
                                    return `<goplus-code half-code language="gop" style="width: 85vw">${iner_src}</goplus-code>`
                                })
                            }
                            console.log(htmlContent)
                        } 
                        htmlContent = htmlContent.replace(/<pre(.*?)>/, "")
                        htmlContent = htmlContent.replace(/<\/pre>/, "")
                    }
                }
                return htmlContent
            }          
            
        }
}



</script>
  
<style>
</style>