<template>

</template>

<script>
    import CherryEngine from 'cherry-markdown/dist/cherry-markdown.engine.core.esm';

    const MarkdownViewer = {
        convertMdToHtml: function(str) {
            var engine = new CherryEngine( {
                // 配置同 cherrymarkdown
            })
            var result = engine.makeHtml(str)
            return result
        }
    }
    export {MarkdownViewer}
</script>

<style scoped>
</style>