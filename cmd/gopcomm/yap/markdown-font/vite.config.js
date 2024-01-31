import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'

// import { fileURLToPath, URL } from 'node:url'

// import {viteSingleFile} from "vite-plugin-singlefile";
// import AutoImport from 'unplugin-auto-import/vite';
// import Components from 'unplugin-vue-components/vite';
// import {ElementPlusResolver} from "unplugin-vue-components/resolvers";



// https://vitejs.dev/config/
// export default defineConfig(({mode}) => {
//   const env = loadEnv(mode, __dirname)
//   return {
//     plugins: [
//       vue(),
//     ],
//     // base: env.VITE_MODE === 'production' ? './' : '/',
//     base: './',
//     resolve: {
//       alias: {
//         //resolve.alias设置别称 解决@绝对路径引入问题
//         "@": fileURLToPath(new URL('./src', import.meta.url))
//       },
//     }
//   }
  
// })

export default defineConfig({
  plugins: [
    vue(),
    // viteSingleFile(),
    // AutoImport({
    //   resolvers: [ElementPlusResolver()],
    // }),
    // Components({
    //   resolvers: [ElementPlusResolver()],
    // }),
  ],
  // build: {
  //   minify: false,
  // },
  
  // resolve: {
  //   alias: {
  //     '@': fileURLToPath(new URL('./src', import.meta.url)) 
  //   }
  // },
  server: {
    cors: true, // 默认启用并允许任何源
    open: true, // 在服务器启动时自动在浏览器中打开应用程序
    //反向代理配置，注意rewrite写法，开始没看文档在这里踩了坑
    proxy: {// 本地开发环境通过代理实现跨域，生产环境使用 nginx 转发
      '/api': {
        target: 'http://localhost:8080', // 通过代理接口访问实际地址。这里是实际访问的地址。vue会通过代理服务器来代理请求
        changeOrigin: true,
        ws: true,  // 允许websocket代理
        rewrite: (path) => path.replace(/^\/api/, '') // 将api替换为空
      }
    }
  }
})
