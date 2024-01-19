<template>
    <a-config-provider :theme="{ token: { colorPrimary: '#3182ce' } }">
        <a-layout class="page-container">
            <!-- 顶部区域 -->
            <a-layout-header class="header-container" style="height: 9vh;">
               <TopSidebar></TopSidebar>
            </a-layout-header>

            <!-- 中部区域 -->
            <arco-tabs type="capsule" size="large" default-active-key="1" lazy-load>
                <!-- BLOG -->
                <arco-tab-pane key="1">
                    <template #title>
                        <MyIcon type="icon-logo-blog-new" style="font-size: 30px; margin-right: 10px;" />
                        BLOG
                    </template>

                    <a-layout class="page-container">
                        <!-- 左栏 -->
                        <a-layout-content class="content-left-container" style="height: 75vh;">
                            <!-- 文章列表 -->
                            <BlogList></BlogList>
                        </a-layout-content>

                        <!-- 右栏 -->
                        <a-layout-sider class="content-right-container" width="20vw" style="height: 75vh;">
                            <BlogListSider></BlogListSider>
                        </a-layout-sider>
                    </a-layout>
                </arco-tab-pane>

                <!-- VIDEO -->
                <arco-tab-pane key="2">
                    <template #title>
                        <MyIcon type="icon-logo-video" style="font-size: 30px; margin-right: 10px;" />
                        VIDEO
                    </template>
                    Content of VIDEO
                </arco-tab-pane>

                <!-- STEM -->
                <arco-tab-pane key="3">
                    <template #title>
                        <MyIcon type="icon-logo-stem-new" style="font-size: 30px; margin-right: 10px;" />
                        STEM
                    </template>
                    Content of STEM
                </arco-tab-pane>

                <!-- ACTIVITY -->
                <arco-tab-pane key="4" title="">
                    <template #title>
                        <MyIcon type="icon-logo-activity" style="font-size: 30px; margin-right: 10px;" />
                        ACTIVITY
                    </template>
                    Content of ACTIVITY
                </arco-tab-pane>
            </arco-tabs>

            <!-- 底部区域 -->
            <a-layout-footer class="footer-container">
                ©2024, provided by 
                <a href="https://www.qiniu.com/" target="_blank" style="color: #3182ce;">Qiniu Tech</a>
            </a-layout-footer>
        </a-layout>
    </a-config-provider>
</template>

<script setup>
import { ref, h } from 'vue';
import { createFromIconfontCN } from '@ant-design/icons-vue';

import TopSidebar from '@/views/component/TopSidebar.vue'
import BlogList from '@/views/blog/BlogList.vue'
import BlogListSider from '@/views/blog/BlogListSider.vue'

const MyIcon = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/c/font_4410595_l5qolflxp8b.js', // 在 iconfont.cn 上生成
});


</script>

<style scoped>
.page-container {
    background-color: transparent;
    color: black;
    height: 100%;
    width: 100%;
    font-family: 'Inter-Regular', '思源黑体-Regular';
    background-image: url("@/assets/pics/home/home-background.png");
    background-repeat: no-repeat;
    background-size: 100% 34%;
    background-position: 50% 0%;
    background-attachment: fixed;
}

.header-container {
    background-image: url("@/assets/pics/logo-gop-community.png");
    background-repeat: no-repeat;
    background-size: 15%;
    background-position: 3% 2%;
    background-attachment: fixed;
    padding: 0;
    background-color: transparent;
}

:deep(.arco-tabs-nav-type-capsule .arco-tabs-nav-tab:not(.arco-tabs-nav-tab-scroll)) {
    justify-content: center;
}

:deep(.arco-tabs-nav-tab-list) {
    width: 70vw;
    height: 8vh;
    opacity: 0.9;
    border-radius: 15px;
    background: linear-gradient(136.8deg, rgba(255, 255, 255, 0.82) 0%, rgba(255, 255, 255, 0.8) 100%);
    border: 2px solid rgba(255, 255, 255, 1);
    box-shadow: 0px 2px 5.5px rgba(0, 0, 0, 0.02);
    backdrop-filter: blur(21px);
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
    padding: 5px 80px;
}

:deep(.arco-tabs-tab-title) {
    font-size: larger;
    font-family: '优设标题黑';
    display: flex;
    flex-direction: row;
    justify-content: center;
}

:deep(.arco-tabs-nav-size-large.arco-tabs-nav-type-capsule .arco-tabs-tab) {
    padding: 0 20px;
    line-height: 40px;
    border-radius: 5px;
}

:deep(.arco-tabs-content) {
    padding: 15px 30px 0 30px;
}

.content-left-container {
    margin: 0 10px 10px 0;
    padding: 10px;
    border-radius: 15px;
    background: rgba(255, 255, 255, 1);
    box-shadow: 0px 5px 14px rgba(0, 0, 0, 0.05);
}

.content-right-container {
    margin: 0 0 10px 5px;
    padding: 20px;
    border-radius: 15px;
    background: rgba(255, 255, 255, 1);
    box-shadow: 0px 5px 14px rgba(0, 0, 0, 0.05);
    display: flex;
    flex-direction: column;
    align-items: center;
}

.footer-container {
    background: transparent;
    padding: 0;
    margin-top: 5px;
    text-align: center;
    font-size: 12px;
}

</style>