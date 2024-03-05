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
                content: "",
                vtts: [],
                videos: [],
                types: [],
            }
        },
        watch: {
            md(newValue) {
                console.log("update markdown", newValue)
                this.computeOut(newValue)

            }
        },
        beforeMount() {
            this.computeOut(this.md)
        },
        mounted() {


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
                let now_t = this.types
                if(matches) {
                    for(let i=0; i<matches.length; i++) {
                        console.log("match", matches[i])
                        // htmlContent = htmlContent.replace(matches[i], "video replace" )
                        htmlContent = htmlContent.replace(matches[i], function(h) { 
                            let src_a = now_s[i]
                            // console.log("---------", this.videos)
                            let vtt_a = now_v[i]
                            let type_a = now_t[i]
                            console.log("=========duibi", vtt_a, src_a)
                            // console.log("src", src)
                            // return `<div><video controls crossorigin playsinline data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" id="player"><source src=${src_a} type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src=${vtt_a} default/><track kind="captions" label="Français" srclang="fr" src=${vtt_a}/><a href=${src_a} download>Download</a></video></div>`
                            let poster = src_a + "/vframe/jpg/offset/7"
                            // poster = "https://th.bing.com/th/id/OIP.YwQZJ_SoLGm-kVT-e-Xc2AHaEo?rs=1&pid=ImgDetMain"
                            return `<div><video controls="" crossorigin="" playsinline="" data-poster=${poster}><source src=${src_a} type=${type_a} size="576"/><track kind="captions" label="English" srclang="en" src=${vtt_a} default/><a href=${src_a} download>Download</a></video></div>`

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
                // console.log("sttt", sttt)
                var replace_vtt = ""
                if(sttt) {
                    for(let i = 0; i<sttt.length; i++) {
                        let p_str = sttt[i]
                        const regexType = /\[(.*?)\]/g; 
                        const matchType = p_str.match(regexType); // 使用 match 函数进行匹配
                        let fileType = "video/mp4"
                        if (matchType){
                            fileType = matchType[0].replace("[","").replace("]","")
                            if (fileType === "video") {
                                fileType = "video/mp4"
                            }
                        }
                        const regex = /\((.*?)\)/g; // 定义正则表达式，其中 \() 表示左括号，(.*?) 表示非贪婪模式匹配任意字符，\)) 表示右括号
                        const matchResult = p_str.match(regex); // 使用 match 函数进行匹配
                        if (matchResult) {
                            console.log("matchResult" , matchResult);
                            matchResult[0] = matchResult[0].replace("(", "").replace(")", "")
                            replace_vtt = mdContent.replace(matchResult[1], "")

                            if (matchResult[1]){
                                matchResult[1] = matchResult[1].replace("(", "").replace(")", "")
                                this.vtts.push(matchResult[1])
                                this.videos.push(matchResult[0])
                                this.types.push(fileType)
                                console.log('have vtt', replace_vtt)
                                
                            } else{
                                replace_vtt = mdContent.replace(matchResult[1], "")
                                this.vtts.push("")
                                this.videos.push(matchResult[0])
                                this.types.push(fileType)
                                console.log('have no vtt', replace_vtt)
                            }
                        } 
                    }
                    return replace_vtt
                }
                
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