<template>
    <arco-list class="list-demo-action-layout" :data="dataSource" :max-height="450" @reach-bottom="fetchData" scrollbar
        hoverable :bordered="false">
        <template #item="{ item, index }">
            <arco-list-item class="list-demo-item" action-layout="vertical" :key="index" style="padding: 8px 20px;">
                <!-- 左侧上部（标题/标签/简介） -->
                <arco-list-item-meta>
                    <template #title>
                        <span style="font-size: 17px; font-family: '思源黑体-Bold';">{{ item.title }}</span>
                    </template>
                    <template #description>
                        <div style="display: flex; flex-direction: row;">
                            <TagsTwoTone style="font-size: 18px; margin-right: 10px;" />
                            <span style="color: #808080;">BUG修复, 开发, 框架, 异常处理</span>
                        </div>
                        <span style="color: #a0aec0; font-size: 13px;">{{ item.description }}</span>
                    </template>
                </arco-list-item-meta>

                <!-- 左侧下部（头像/昵称/时间） -->
                <arco-list-item-meta>
                    <template #avatar>
                        <div>
                            <arco-avatar shape="circle" :size="26">
                                <img alt="avatar" :src="item.avatar" />
                            </arco-avatar>
                            <span style="margin-left: 10px;">{{ item.name }}</span>
                            <span style="margin-left: 10px; color: #a0aec0; font-size: 13px;">· 2024-01-10</span>
                        </div>
                    </template>
                </arco-list-item-meta>

                <!-- 右侧（题图/互动数量） -->
                <template #extra>
                    <div className="image-area">
                        <img alt="arco-design" :src="item.imageSrc" />
                    </div>
                    <span style="margin-left: 10px;">
                        <EyeTwoTone /> 83
                    </span>
                    <span style="margin-left: 25px;">
                        <LikeTwoTone /> {{ item.index }}
                    </span>
                    <span style="margin-left: 25px;">Share</span>
                </template>
            </arco-list-item>
        </template>
    </arco-list>
</template>

<script setup>
import { ref } from 'vue'
import { TagsTwoTone, EyeTwoTone, LikeTwoTone } from '@ant-design/icons-vue';

defineProps({
    msg: String,
})



const names = ['Socrates', 'Balzac', 'Plato'];
const avatarSrc = [
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/a8c8cdb109cb051163646151a4a5083b.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/e278888093bef8910e829486fb45dd69.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/9eeb1800d9b78349b24682c3518ac4a3.png~tplv-uwbnlip3yd-webp.webp',
];
const imageSrc = [
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/29c1f9d7d17c503c5d7bf4e538cb7c4f.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/04d7bc31dd67dcdf380bc3f6aa07599f.png~tplv-uwbnlip3yd-webp.webp',
    '//p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/1f61854a849a076318ed527c8fca1bbf.png~tplv-uwbnlip3yd-webp.webp',
];
const dataSource = (
    new Array(15).fill(null).map((_, index) => {
        return {
            index: index,
            avatar: avatarSrc[index % avatarSrc.length],
            name: names[index % names.length],
            title: '异常≠错误，正如BUG≠事故，详解业务开发中的异常处理',
            description:
                '软件开发中遇到异常才是正常，很少有人能写出完美的程序跑在任何机器上都不会报错。但极为正常的软件异常，却经常出自不同的原因，导致不同的结果。怎么样科学地认识异常、处理异常，是很多研发同学需要解同学需要',
            imageSrc: imageSrc[index % imageSrc.length],
        };
    })
);

</script>

<style scoped>
.list-demo-action-layout .image-area {
    width: 13vw;
    height: 16vh;
    border-radius: 2px;
    overflow: hidden;
    margin-bottom: 5px;
}

.list-demo-action-layout .list-demo-item {
    border-bottom: 1px solid var(--color-fill-3);
    cursor: pointer;
}

.list-demo-action-layout .image-area img {
    width: 100%;
    margin-left: 10px;
}

.list-demo-action-layout .arco-list-item-action .arco-icon {
    margin: 0 4px;
}
</style>