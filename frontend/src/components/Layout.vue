<template>
  <div class="d-flex flex-column h-100">
    <!-- 引入 NavBar 组件 -->
    <NavBar />

    <!-- 主体布局 -->
    <div class="d-flex flex-grow-1">
      <!-- 左侧菜单，只有用户登录后才显示 -->
      <div v-if="isLoggedIn" :class="['bg-light', {'d-none': !menuVisible}, 'p-3', 'menu', 'min-vh-100', 'border-end']">
        <ul class="nav flex-column">
          <li class="nav-item" v-for="(menu, index) in menuList" :key="index">
            <router-link :to="menu.link" class="btn btn-link">
              {{ menu.name }}
            </router-link>
          </li>
        </ul>
      </div>

      <!-- 分割线，点击切换左侧菜单显示/隐藏 -->
      <div
        class="bg-dark d-flex justify-content-center align-items-center"
        style="cursor: pointer; width: 5px; height: 100vh;"
        @click="toggleMenu"
      >
        <span class="text-white" style="font-size: 20px;">&#9776;</span>
      </div>

      <!-- 右侧内容区域 -->
      <div class="flex-fill p-3">
        <router-view />
      </div>
    </div>
  </div>
</template>

<script>
import NavBar from './NavBar.vue'; // 引入 NavBar 组件

export default {
  components: {
    NavBar,
  },
  data() {
    return {
      isLoggedIn: false,  // 默认未登录
      menuVisible: true,  // 菜单默认显示
      menuList: [
        { name: '首页', link: '/home' },
        { name: '功能1', link: '/function1' },
        { name: '功能2', link: '/function2' },
        { name: '设置', link: '/settings' },
      ],
    };
  },
  created() {
    // 页面加载时检查登录状态
    this.checkLoginStatus();
  },
  methods: {
    toggleMenu() {
      this.menuVisible = !this.menuVisible; // 切换菜单显示/隐藏
    },
    checkLoginStatus() {
      // 如果 localStorage 中有 authToken，则表示用户已登录
      this.isLoggedIn = !!localStorage.getItem('authToken');
    }
  }
};
</script>

<style scoped>
/* 样式与之前一致，确保布局正常 */
.menu {
  width: auto;
  transition: width 0.3s ease;
}

.bg-dark {
  background-color: #343a40;
}

.bg-light {
  background-color: #f8f9fa;
}

.min-vh-100 {
  min-height: 100vh;
}

.border-end {
  border-right: 1px solid #ddd;
}

.nav-item a {
  text-align: left;
  width: 100%;
  background: none;
  border: none;
  padding: 10px;
  color: #007bff;
  font-size: 16px;
  text-decoration: none;
}

.nav-item a:hover {
  background-color: #f1f1f1;
}

.flex-fill {
  flex: 1;
}
</style>
