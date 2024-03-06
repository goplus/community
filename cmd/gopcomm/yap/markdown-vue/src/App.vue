<script setup>
import MarkdownEditor from './components/MarkdownEditor.vue';
import MarkdownViewerOld from './components/MarkdownViewerOld.vue';
import MarkdownViewer from './components/MarkdownViewer.vue';

</script>



<template>
   
    <!-- <markdown-editor></markdown-editor> -->
    <div id="container">

         <div id="viewer" style="margin:20px;">ffffffffffffffffffffff</div>
        <markdown-viewer :md="c" ></markdown-viewer>
    </div>
    
    <!-- <markdown-viewer-old :md="c"></markdown-viewer-old> -->
</template> 


<script>

import Plyr from 'plyr';
import 'plyr/dist/plyr.css'

export default  {
    data() {
        return {
            // c: "!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)"
//             c: `!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)


// !video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)`
c: `- 123
> 456
# welcome to cherry editor!
`
        }
    },
    mounted() {
        const video = document.querySelectorAll('video');
	
        for (var i = 0; i < video.length; i++) {
            console.log(video[i])
            
            const player = new Plyr(video[i], {captions: {active: true, update: true, language: 'en'}});
        }
    }
}

</script>

<style>
/* @import url('https://cdn.plyr.io/3.6.8/plyr.css');     */

#container{
        width: 50%;
        margin: 0;
        padding: 0;
        height: 100%;
        overflow: hidden;
        position: relative;
    }
    #markdown-container{
        width: 50%;
        margin: 0;
        padding: 0;
        height: 100%;
        overflow: hidden;
        position: absolute;
    }
</style>







<!-- <video playsinline="" crossorigin="" controls="" poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" ><source size="576" type="video/mp4" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4"><track src="default/" srclang="en" label="English" kind="captions"><track src="/" srclang="fr" label="Français" kind="captions"><a download="" href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4">Download</a></video>
<video controls="" crossorigin="" playsinline="" poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-720p.mp4" type="video/mp4" size="720"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-1080p.mp4" type="video/mp4" size="1080"><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default=""><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"><a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download="">Download</a></video> -->