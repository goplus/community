<template>
    <a-config-provider :theme="{ token: { colorPrimary: '#3182ce' } }">
        <a-layout class="page-container">
            <!-- 顶部区域 -->
            <a-affix :offset-top="top">
                <a-layout-header class="header-container" style="height: 9vh;">
                    <TopSidebar></TopSidebar>
                </a-layout-header>
            </a-affix>

            <!-- 内容区域 -->
            <a-layout-content class="content-container">
                <!-- 左侧 -->
                <div class="left-content">
                    <!-- 面包屑 -->
                    <div class="breadcrumb">
                        <a-breadcrumb style="font-family: 'Inter-Regular', '思源黑体-Regular';">
                            <template #separator>
                                <span style="color: #3182ce;">/</span>
                            </template>
                            <a-breadcrumb-item>Home</a-breadcrumb-item>
                            <a-breadcrumb-item href="">Blog</a-breadcrumb-item>
                            <a-breadcrumb-item>异常≠错误，正如BUG≠事故，详解业务开发中的异常处理</a-breadcrumb-item>
                        </a-breadcrumb>
                    </div>

                    <div style="display: flex; flex-direction: row;">
                        <!-- 互动按钮组 -->
                        <div class="button-group">
                            <!-- 点赞 -->
                            <a-badge count="27" :number-style="{
                                backgroundColor: '#fff',
                                color: '#f76969',
                                borderColor: '#f76969',
                                boxShadow: '0 0 0 1px #d9d9d9 inset',
                                fontFamily: 'Inter-Regular'
                            }">
                                <a-button size="large" class="btn-1">
                                    <template #icon>
                                        <LikeOutlined />
                                    </template>
                                </a-button>
                            </a-badge>

                            <!-- 评论 -->
                            <a-badge count="7" :number-style="{
                                backgroundColor: '#fff',
                                color: '#83a5fa',
                                borderColor: '#83a5fa',
                                boxShadow: '0 0 0 1px #d9d9d9 inset',
                                fontFamily: 'Inter-Regular'
                            }">
                                <a-button size="large" class="btn-2" @click="openCommentArea">
                                    <template #icon>
                                        <CommentOutlined />
                                    </template>
                                </a-button>
                            </a-badge>
                            <!-- 评论区 -->
                            <arco-drawer :width="400" v-model:visible="showComment" :footer="false">
                                <template #title>
                                    Comment
                                </template>
                                <div>
                                    <arco-comment 
                                        author="Socrates"
                                        avatar="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/3ee5f13fb09879ecb5185e440cef6eb9.png~tplv-uwbnlip3yd-webp.webp"
                                        content="Comment body content." 
                                        datetime="1 hour"
                                        align="right"
                                    >
                                        <template #actions>
                                            <span class="action">
                                                <IconMessage /> Reply
                                            </span>
                                        </template>
                                        <arco-comment 
                                            author="Balzac"
                                            avatar="https://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/9eeb1800d9b78349b24682c3518ac4a3.png~tplv-uwbnlip3yd-webp.webp"
                                            content="Comment body content." 
                                            datetime="1 hour"
                                            align="right"
                                        >
                                            <template #actions>
                                                <span class="action">
                                                    <IconMessage /> Reply
                                                </span>
                                            </template>
                                        </arco-comment>
                                    </arco-comment>
                                </div>
                            </arco-drawer>

                            <!-- 收藏 -->
                            <a-badge count="13" :number-style="{
                                backgroundColor: '#fff',
                                color: '#05a3b6',
                                borderColor: '#05a3b6',
                                boxShadow: '0 0 0 1px #d9d9d9 inset',
                                fontFamily: 'Inter-Regular'
                            }">
                                <a-button size="large" class="btn-3">
                                    <template #icon>
                                        <StarOutlined />
                                    </template>
                                </a-button>
                            </a-badge>

                            <!-- 分享 -->
                            <a-dropdown :placement="placement" arrow>
                                <a-button size="large">
                                    <template #icon>
                                        <ShareAltOutlined />
                                    </template>
                                </a-button>
                                <template #overlay>
                                    <a-menu>
                                        <a-menu-item key="1">
                                            <CodeOutlined />
                                            Share to Community
                                        </a-menu-item>
                                        <a-menu-item key="2">
                                            <a href="javascript:;">Share to ...</a>
                                        </a-menu-item>
                                        <a-menu-item key="3">
                                            <a href="javascript:;">Share to ...</a>
                                        </a-menu-item>
                                    </a-menu>
                                </template>
                            </a-dropdown>
                        </div>

                        <!-- 文章内容 -->
                        <div class="article-content">
                            <Article></Article>
                        </div>
                    </div>
                </div>

                <!-- 右侧 -->
                <div class="right-content">
                    <!-- 作者信息 -->
                    <div class="creator-info">
                        <!-- 头像/昵称 -->
                        <div style="display: flex; flex-direction: row;">
                            <arco-avatar :size="46" :style="{ backgroundColor: '#3182ce' }">Nick</arco-avatar>
                            <div>
                                <span class="creator-name">Nick_1977<br></span>
                                <span class="creator-desc">Go+官方认证开发者</span>
                            </div>
                        </div>

                        <!-- 用户数据 -->
                        <div style="display: flex; flex-direction: row; justify-content: space-between; margin: 10px;">
                            <div style="text-align: center;">
                                <span class="creator-dataname">Blogs<br></span>
                                <span class="creator-data">18</span>
                            </div>
                            <div style="text-align: center;">
                                <span class="creator-dataname">Views<br></span>
                                <span class="creator-data">1.3k</span>
                            </div>
                            <div style="text-align: center;">
                                <span class="creator-dataname">Likes<br></span>
                                <span class="creator-data">279</span>
                            </div>
                            <div style="text-align: center;">
                                <span class="creator-dataname">Followers<br></span>
                                <span class="creator-data">161</span>
                            </div>
                        </div>

                        <!-- 关注按钮 -->
                        <div style="padding: 0 30px;">
                            <a-button block>
                                <template #icon>
                                    <PlusOutlined />
                                </template>
                                Follow
                            </a-button>
                        </div>
                    </div>

                    <!-- 文章目录 -->
                    <div class="article-catalog">
                        TODO 文章目录
                    </div>
                </div>
            </a-layout-content>
        </a-layout>
    </a-config-provider>
</template>

<script setup>
import { ref } from 'vue';
import { LikeOutlined, CommentOutlined, StarOutlined, ShareAltOutlined, CodeOutlined, PlusOutlined } from '@ant-design/icons-vue';
import TopSidebar from '@/views/component/TopSidebar.vue';
import Article from '@/views/component/Article.vue';

const showComment = ref(false);
const openCommentArea = () => {
    showComment.value = true;
};
const handel = () => {
    showComment.value = true;
};
</script>

<style scoped>
.page-container {
    background-color: transparent;
    color: black;
    height: 100%;
    width: 100%;
    font-family: 'Inter-Regular', '思源黑体-Regular';
    background-image: url("@/assets/pics/blog/blog-background.png");
    background-repeat: no-repeat;
    background-size: 100% 9%;
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

.header-menu {
    float: right;
}

.content-container {
    padding: 10px 20px;
    display: flex;
    flex-direction: row;
}

.left-content {}

.breadcrumb {
    width: 72vw;
    height: 6vh;
    padding: 10px 20px;
    border-radius: 10px;
    background: rgba(255, 255, 255, 1);
    box-shadow: 0px 5px 14px rgba(0, 0, 0, 0.05);
}

.button-group {
    margin: 20px 30px 0 0;
    height: 35vh;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}

.btn-1:hover {
    border-color: #f76969;
    color: #f76969;
}

.btn-2:hover {
    border-color: #83a5fa;
    color: #83a5fa;
}

.btn-3:hover {
    border-color: #05a3b6;
    color: #05a3b6;
}

.article-content {
    width: 66vw;
    height: 78vh;
    overflow: auto;
    margin-top: 20px;
    padding: 0 20px;
    background: rgba(128, 128, 128, 0.2);
}

.creator-info {
    width: 23vw;
    height: 40vh;
    margin-left: 15px;
    padding: 20px;
    border-radius: 10px;
    background: rgba(255, 255, 255, 1);
    box-shadow: 0px 5px 14px rgba(0, 0, 0, 0.05);
}

.creator-name {
    font-size: 16px;
    margin-left: 10px;
}

.creator-desc {
    font-size: 13px;
    color: #a0aec0;
    margin-left: 10px;
}

.creator-dataname {
    font-size: 13px;
    color: #a0aec0;
}

.creator-data {
    font-size: 15px;
}

.article-catalog {
    width: 23vw;
    height: 45vh;
    margin: 15px 0 0 15px;
    padding: 20px;
    border-radius: 10px;
    background: rgba(255, 255, 255, 1);
    box-shadow: 0px 5px 14px rgba(0, 0, 0, 0.05);
}

.footer-container {
    background: transparent;
    padding: 0;
    margin-top: 5px;
    text-align: center;
    font-size: 12px;
}
</style>