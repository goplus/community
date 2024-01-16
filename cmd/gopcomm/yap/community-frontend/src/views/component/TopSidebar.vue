<template>
    <a-row :gutter="8" class="header-menu" style="margin: 0;">
        <!-- 搜索框 -->
        <a-col :span="12">
            <a-input v-model:value="searchContent" placeholder="Search">
                <template #prefix>
                    <SearchOutlined style="color: rgba(0, 0, 0, 0.45)" />
                </template>
            </a-input>
        </a-col>

        <!-- 登录 -->
        <a-col :span="6">
            <a-dropdown>
                <a-button type="text" style="color: #fff;">
                    <UserOutlined />
                    Sign In
                </a-button>
                <template #overlay>
                    <a-menu>
                        <a-menu-item>
                            <a href="javascript:;">1st menu item</a>
                        </a-menu-item>
                        <a-menu-item>
                            <a href="javascript:;">2nd menu item</a>
                        </a-menu-item>
                        <a-menu-item>
                            <a href="javascript:;">3rd menu item</a>
                        </a-menu-item>
                    </a-menu>
                </template>
            </a-dropdown>
        </a-col>

        <!-- 设置 -->
        <a-col :span="2">
            <a-button type="text" shape="circle" :icon="h(SettingOutlined)" style="color: #fff;"></a-button>
        </a-col>

        <!-- 通知 -->
        <a-col :span="2">
            <a-button type="text" shape="circle" :icon="h(BellOutlined)" style="color: #fff;"></a-button>
        </a-col>
    </a-row>
</template>

<script setup>
import { ref, h } from 'vue';
import { SearchOutlined, UserOutlined, SettingOutlined, BellOutlined } from '@ant-design/icons-vue';
const searchContent = ref('');

</script>

<style scoped>
.header-menu {
    float: right;
}
</style>