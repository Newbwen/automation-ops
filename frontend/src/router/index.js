import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      meta: {
        requiresAuth: false //明确标记不需要登录即可访问的页面
      }
    },
    {
      path: '/error',
      name: 'error',
      component: () => import('../views/errors/ErrorView.vue'),
      meta: {
        requiresAuth: false //需要登录才能访问的页面
      }
    },
    {path: '/index',
      name: 'index',
      component: () => import('../views/IndexView.vue')
    },
    {
      path: '/host/list',
      name: 'listHost',
      component: () => import('../views/hosts/ListView.vue')
    },
    {
      path: '/welcome',
      name: 'welcome',
      component: () => import('../views/WelcomeView.vue'),
      meta: {
        requiresAuth: true //需要登录才能访问的页面
      }
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
    },
    {
      path: '/user/profile',
      name: 'user-profile',
      component: () => import('../views/UserProfileView.vue'),
      meta: {
        requiresAuth: true //需要登录才能访问的页面
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const isAuthenticated = checkAuthStatus(); // 检查用户是否已认证

  // 如果目标路由需要认证，但用户未认证，则跳转到首页
  if (to.matched.some(record => record.meta.requiresAuth) && !isAuthenticated) {
    next({ name: 'home' });
    return;
  }

  // 如果用户已认证，尝试访问登录页，则跳转到用户个人资料页
  if (to.name === 'login' && isAuthenticated) {
    next({ name: 'welcome' });
    return;
  }


  // 如果目标路由只允许未认证用户访问，但用户已认证，则跳转到首页
  if (to.matched.some(record => record.meta.guestOnly) && isAuthenticated) {
    next({ name: 'welcome' });
    return;
  }

  // 其他情况，允许访问
  next();
});

// 检查用户认证状态的方法
function checkAuthStatus() {
  // 从 localStorage 中获取认证令牌
  const token = localStorage.getItem('authToken');

  // 返回令牌是否存在的布尔值
  return !!token;
}




export default router