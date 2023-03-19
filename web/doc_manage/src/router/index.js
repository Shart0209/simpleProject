import { createRouter, createWebHistory } from "vue-router";
import HomeViews from '@/views/HomeViews.vue'

const baseUrl = import.meta.env.VITE_BUILD_ADDRESS;

export const routes = [
    {
        path: `${baseUrl}/`,
        name: "home",
        component: HomeViews
    },
];

export const router = createRouter({
    history: createWebHistory(),
    routes: routes,
});
