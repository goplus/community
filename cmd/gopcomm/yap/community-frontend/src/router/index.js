import { createRouter, createWebHistory } from "vue-router";

const routes = [
    { // 地址重定向
        path: '/', 
        redirect: '/goplus/community/home' 
    }, 
    {
        path: '/goplus/community/home',
        name: 'Home',
        component: () =>
                import("../views/Home.vue"),
    },
    {
        path: '/goplus/community/blog',
        name: 'BlogDetail',
        component: () =>
                import("../views/blog/BlogDetail.vue"),
    },
    // {
    //     path: '/user/login',
    //     name: 'LoginCard',
    //     component: () =>
    //         import("../views/UserLogin/LoginCard.vue"),
    // },
];

const router = createRouter({
    history: createWebHistory(''),
    routes,
});

export default router;