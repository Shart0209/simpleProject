import { createRouter, createWebHistory } from 'vue-router';

import HomeViews from '@/views/HomeViews.vue';
import BoardViews from '@/views/BoardViews.vue';
import CreateViews from '@/views/CreateViews.vue';
import BoardDetailViews from '@/views/BoardDetailViews.vue';
import NotFoundViews from '@/views/NotFoundViews.vue';

import LoginViews from '@/views/LoginViews.vue';

export const routes = [
  {
    path: '/',
    component: () => import('@/layouts/Default.vue'),
    children: [{ path: '', name: 'home', component: HomeViews }],
  },
  {
    path: '/board',
    component: () => import('@/layouts/Default.vue'),
    children: [{ path: '', name: 'board', component: BoardViews }],
  },
  {
    path: '/add',
    component: () => import('@/layouts/Default.vue'),
    children: [{ path: '', name: 'add', component: CreateViews }],
  },
  {
    path: '/board/:id',
    component: () => import('@/layouts/Default.vue'),
    children: [
      { path: '', name: 'board_detail_id', component: BoardDetailViews },
    ],
  },
  {
    path: '/login',
    component: () => import('@/layouts/Default.vue'),
    children: [{ path: '', name: 'login', component: LoginViews }],
  },
  {
    path: '/:pathMatch(.*)*',
    component: () => import('@/layouts/Default.vue'),
    children: [{ path: '', name: '404', component: NotFoundViews }],
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});
