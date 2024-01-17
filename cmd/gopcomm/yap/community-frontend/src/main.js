import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import router from './router'

import Antd from 'ant-design-vue';
import 'ant-design-vue/dist/reset.css';
import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';

const app = createApp(App);
app.use(router);
app.use(Antd);
app.use(ArcoVue, {
    componentPrefix: 'arco' // 用于改变使用组件时的前缀名称
});
app.mount('#app');