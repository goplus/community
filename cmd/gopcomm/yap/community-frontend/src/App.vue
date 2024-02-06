<template>
    <router-view></router-view>
</template>

<script setup>
const debounce = function (callback, delay) {
    let tid;
    return function (...args) {
        const ctx = this;
        if (tid) clearTimeout(tid);
        tid = setTimeout(() => {
            callback.apply(ctx, args);
        }, delay);
    };
};

const OriginalResizeObserver = window.ResizeObserver;

window.ResizeObserver = class ResizeObserver extends OriginalResizeObserver {
    constructor(callback) {
        super(debounce(callback, 20));
    }
};
</script>

<style>
@import './assets/fonts/font.css';
</style>
