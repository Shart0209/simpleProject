import { createRouter, createWebHistory } from 'vue-router';

export const routes = [
  {
    path: '/',
    component: () => import('@/layouts/Default.vue'),
    children: [
      {
        path: '',
        name: 'home',
        component: () => import('@/views/HomeViews.vue'),
      },
    ],
  },
  {
    path: '/board',
    component: () => import('@/layouts/Default.vue'),
    children: [
      {
        path: '',
        name: 'board',
        component: () => import('@/views/BoardViews.vue'),
      },
    ],
  },
  {
    path: '/add',
    component: () => import('@/layouts/Default.vue'),
    children: [
      {
        path: '',
        name: 'add',
        component: () => import('@/views/CreateViews.vue'),
      },
    ],
  },
  {
    path: '/board/:id',
    component: () => import('@/layouts/Default.vue'),
    children: [
      {
        path: '',
        name: 'board_detail_id',
        component: () => import('@/views/BoardDetailViews.vue'),
      },
    ],
  },
  {
    path: '/login',
    component: () => import('@/layouts/Page.vue'),
    children: [
      {
        path: '',
        name: 'login',
        component: () => import('@/views/LoginViews.vue'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    component: () => import('@/layouts/Default.vue'),
    children: [
      {
        path: '',
        name: '404',
        component: () => import('@/views/NotFoundViews.vue'),
      },
    ],
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes: routes,
});

router.beforeEach((to, from, next) => {
  if (!to.matched.length) console.warn('no match');
  next();
});
