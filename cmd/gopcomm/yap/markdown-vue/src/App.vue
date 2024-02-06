<script setup>
import Cherry from './components/Cherry.vue';
// import gop from './components/gop.vue'

// import { ref } from 'vue';

// let con = ref("hhhhh")

</script>

<!-- <script>
  !function(){
        var t,e={
            code:[
                "https://goplus.org/_next/static/widgets/code.f81abac15122c88e65d7.css",
                "https://goplus.org/_next/static/widgets/code.85827e18ab6a0fa63bdc.js"
            ],
            footer:[
                "https://goplus.org/_next/static/widgets/footer.5f76b11115203767e7c3.css",
                "https://goplus.org/_next/static/widgets/footer.64aca03204c8ae8ad97e.js"
            ],
            header:[
                "https://goplus.org/_next/static/widgets/header.5445867c6da0bcd5587b.css",
                "https://goplus.org/_next/static/widgets/header.9ea108e482eb1e740a30.js"
            ]
        };
        // ((null===(t=document.currentScript)||void 0===t?void 0:'code')||"").split(",")
        ["code"]
        .map((function(t){return t.trim()}))
        .filter(Boolean)
        .forEach((
            function(t){
                var s=e[t]||[];
                function o(){
                    var t=s.find((function(t){return/\.js$/.test(t)})),
                    e=s.find((function(t){return/\.css$/.test(t)}));
                    if(null!=t){
                        var o=document.createElement("script");
                        o.src=t,null!=e&&o.setAttribute("data-style-url",e),document.body.appendChild(o)
                    }
                }
                "loading"!==document.readyState?o():window.addEventListener("DOMContentLoaded",o)
            }
        ))
    }();
</script> -->

<template>
    <Cherry></Cherry>

    <!-- <gop :cont="con"></gop> -->

</template>

<style scoped>

</style>
