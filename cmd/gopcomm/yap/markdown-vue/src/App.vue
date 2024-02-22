<script setup>
import MarkdownEditor from './components/MarkdownEditor.vue';
import MarkdownViewer from './components/MarkdownViewer.vue';

</script>



<template>
    <markdown-editor></markdown-editor>
    <markdown-viewer :md="c"></markdown-viewer>
    <!-- <markdown-viewer md="!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)"></markdown-viewer> -->
    <!-- <video id="player" playsinline controls data-poster="/path/to/poster.jpg" ref="plyr">
        <source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" />
    </video> -->
    <!-- <vue-plyr ref="plyr">
        <video ref="plyr" class="plyr" :autoplay="autoplay" :crossorigin="crossorigin" :poster="poster" :playsinline="playsinline">
            <source :src="source" type="video/mp4">
        </video>
    </vue-plyr> -->
    <!-- <video controls crossorigin playsinline poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg">
		<source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576">
		<source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-720p.mp4" type="video/mp4" size="720">
		<source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-1080p.mp4" type="video/mp4" size="1080">

		<track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt"
				default>
		<track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt">
		<a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download>Download</a>
	</video> -->

  

</template> 


<script>
// console.log(MarkdownViewer.methods.convertMdToHtml("!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)"))
// import Plyr from 'plyr';

// const player = new Plyr('#player');
// import VuePlyr from 'vue-plyr'
// import 'vue-plyr/dist/vue-plyr.css'
// import Plyr from 'plyr';
// import "https://cdn.plyr.io/3.7.8/plyr.css"

export default  {
    data() {
        return {
            // source: 'https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4', // 修改为你的本地视频路径
            // autoplay: true,
            // crossorigin: 'anonymous',
            // poster: '',
            // playsinline: true
            c: "!video[video/mp4](https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4)(https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt)"
        }
    },
    mounted() {
        setTimeout(() => {
            this.c = "```gop ```"
            console.log("set time out")
        }, 5000)
        // const player = new Plyr(this.$refs.plyr, {
        //     controls: ['play-large', 'play', 'progress', 'current-time', 'mute', 'volume', 'airplay', 'fullscreen']
        // })
        // 播放视频
        // player.play()

        // const player = new Plyr('video', {captions: {active: true}});
        // // Expose player so it can be used from the console
        // window.player = player;
    }
}

</script>

<style>
/* @import url('https://cdn.plyr.io/3.6.8/plyr.css');     */

</style>







<!-- <video playsinline="" crossorigin="" controls="" poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg" ><source size="576" type="video/mp4" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4"><track src="default/" srclang="en" label="English" kind="captions"><track src="/" srclang="fr" label="Français" kind="captions"><a download="" href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4">Download</a></video>
<video controls="" crossorigin="" playsinline="" poster="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.jpg"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" type="video/mp4" size="576"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-720p.mp4" type="video/mp4" size="720"><source src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-1080p.mp4" type="video/mp4" size="1080"><track kind="captions" label="English" srclang="en" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.en.vtt" default=""><track kind="captions" label="Français" srclang="fr" src="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-HD.fr.vtt"><a href="https://cdn.plyr.io/static/demo/View_From_A_Blue_Moon_Trailer-576p.mp4" download="">Download</a></video> -->