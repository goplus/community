<script setup>
import MarkdownEditor from './components/MarkdownEditor.vue';
import MarkdownViewerOld from './components/MarkdownViewerOld.vue';
import MarkdownViewer from './components/MarkdownViewer.vue';

</script>



<template>
   
    <markdown-editor></markdown-editor>
    <markdown-viewer :md="c" ></markdown-viewer>
    <!-- <markdown-viewer-old :md="c"></markdown-viewer-old> -->
</template> 


<script>

import Plyr from 'plyr';
import 'plyr/dist/plyr.css'

export default  {
    data() {
        return {
            // c: "!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)"
            c: `!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)


!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)`
// c: `# head
// ## sss

// - 123
// > 456
// # welcome to cherry editor!
// `
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
</style>
