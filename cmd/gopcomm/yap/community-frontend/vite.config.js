import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path'
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';

export default defineConfig({
    plugins: [
        vue(),
        Components({
            resolvers: [
                AntDesignVueResolver({
                    importStyle: false, // css in js
                }),
            ],
        }),
    ],
    css: {
        preprocessorOptions: {
            less: {
                'arcoblue-6': '#262626',
            },
            javascriptEnabled: true,
        }
    },
    resolve: {
        alias: {
            // 设置别名
            '@': path.resolve(__dirname, 'src')
        },
        extensions: ['.mjs', '.js', '.ts', '.jsx', '.tsx', '.json', '.vue']
    },
    // server: {
    //     host: 'localhost',
    //     hmr: true
    // }
})