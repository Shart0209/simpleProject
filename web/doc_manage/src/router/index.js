import { createRouter, createWebHistory } from "vue-router";

const baseUrl = import.meta.env.VITE_BUILD_ADDRESS;

export const routes = [
    {
        path: `${baseUrl}/`,
        component: () => import("@/layouts/Default.vue"),
        children: [
            { path: "", name: "Home", component: () => import("@/views/HomeViews.vue") }
        ],
    },
];

export const router = createRouter({
    history: createWebHistory(),
    routes: routes,
});
