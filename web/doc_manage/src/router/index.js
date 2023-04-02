import { createRouter, createWebHistory } from 'vue-router';

import HomeViews from '@/views/HomeViews.vue';
import BoardViews from '@/views/BoardViews.vue';
import BoardDetailViews from '@/views/BoardDetailViews.vue';
import NotFoundViews from '@/views/NotFoundViews.vue';
import TestViews from '@/views/TestViews.vue';

export const routes = [
  {
    path: '/',
    component: () => import("@/layouts/Default.vue"),
    children: [
      { path: "", name: "home", component: HomeViews}
    ]
  },
  {
    path: '/board',
    component: () => import("@/layouts/Default.vue"),
    children: [
      { path: "", name: "board", component: BoardViews}
    ]   
  },
  {
    path: '/test',
    component: () => import("@/layouts/Default.vue"),
    children: [
      { path: "", name: "test", component: TestViews}
    ]   
  },
  {
    path: '/board/:id',
    component: () => import("@/layouts/Default.vue"),
    children: [
      { path: "", name: "board_detail_id", component: BoardDetailViews}
    ] 
  },
  {
    path: '/:pathMatch(.*)*',
    component: () => import("@/layouts/Default.vue"),
    children: [
      { path: "", name: "404", component: NotFoundViews}
    ] 
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});
