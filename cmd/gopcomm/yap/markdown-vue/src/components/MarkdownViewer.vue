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

export default {
        name: "MarkdownViewer",
        props: {
            "md": String
        },
        components: {
            
        },
        data() {
            return {
                content: '<h1>Hello World!</h1>' // HTML内容
            }
        },
        mounted() {
            console.log("mounted", this.md)
            const cherryEngineInstance = new CherryEngine();
            // markdown语法以 <code>gop 开头</code>结尾
            // video <video>开头 </video>结尾 
            // <code>gop println&quot;&quot;</code>
            var htmlContent = cherryEngineInstance.makeHtml(this.md);
            htmlContent = this.replaceGoplus(htmlContent)
            htmlContent = this.replaceVideo(htmlContent)
            // htmlContent = this.replaceGoplus(htmlContent)
            this.content = htmlContent
            console.log(htmlContent)
            
        },
        methods: {  
            replaceVideo(htmlContent) {
                var regex = /<video(.*?)<\/video>/g
                var matches = htmlContent.match(regex); // 使用match()函数获取匹配结果
                console.log("m", matches)
                if(matches) {
                    for(let i=0; i<matches.length; i++) {
                        console.log("match", matches[i])
                        // htmlContent = htmlContent.replace(matches[i], "video replace" )
                        htmlContent = htmlContent.replace(matches[i], function(h) { 
                            var regex1 = /src="(.*?)">video\/mp4/
                            let src = matches[i].match(regex1)[1]
                            console.log("src", src)
                            return `<div><video controls crossorigin playsinline data-poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" id="player"><source src=${src} type="video/mp4" size="576"/><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default/><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"/><a href=${src} download>Download</a></video></div>`
                        })
                    }
                    console.log("replace")
                } else {

                }
                return htmlContent
            },
            replaceGoplus(htmlContent) {

                var regex = /<code>(.*?)<\/code>/g
                var matches = htmlContent.match(regex); // 使用match()函数获取匹配结果
                if(matches) {
                    for(let i=0; i<matches.length; i++) {
                        console.log("match", matches[i])
                        // htmlContent = htmlContent.replace(matches[i], "video replace" )
                        htmlContent = htmlContent.replace(matches[i], function(h) { 
                            var regex1 = /<code>(.*?)<\/code>/
                            let src = matches[i].match(regex1)[1]
                            console.log("src", src)
                            return `<goplus-code half-code language="gop" style="width: 85vw">${src}</goplus-code>`
                        })
                    }
                    console.log(htmlContent)
                } 
                
                return htmlContent
            }          
            
        }
}



</script>
  
<style>
</style>