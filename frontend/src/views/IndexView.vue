<template>
  <div class="d-flex flex-column h-100">

    <!-- 主体布局 -->
    <div class="d-flex flex-grow-1">
      <!-- 左侧菜单 -->
      <div :class="['bg-light', {'d-none': !menuVisible}, 'p-3', 'menu', 'min-vh-100', 'border-end']">
        <ul class="nav flex-column">
          <li class="nav-item" v-for="(menu, index) in menuList" :key="index">
            <button class="btn btn-link" @click="changePage(menu)">
              {{ menu.name }}
            </button>
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
        <h2>{{ currentPage.title }}</h2>
        <div>
          <p>{{ currentPage.content }}</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>


export default {
  data() {
    return {
      menuVisible: true, // 控制菜单的显示与隐藏
      currentPage: {
        title: '首页',
        content: '这里是首页的内容。',
      },
      menuList: [
        { name: '首页', title: '首页', content: '这里是首页的内容。' },
        { name: '功能1', title: '功能1', content: '这里是功能1的内容。' },
        { name: '功能2', title: '功能2', content: '这里是功能2的内容。' },
        { name: '设置', title: '设置', content: '这里是设置的内容。' },
      ],
    };
  },
  methods: {
    toggleMenu() {
      this.menuVisible = !this.menuVisible;
    },
    changePage(menu) {
      this.currentPage = {
        title: menu.title,
        content: menu.content,
      };
    }
  }
};
</script>

<style scoped>
#app {
  height: 100vh;
  display: flex;
  flex-direction: column;
}

/* 调整菜单栏样式，使左侧菜单栏宽度自适应 */
.menu {
  width: auto; /* 菜单宽度自适应 */
  transition: width 0.3s ease; /* 动画效果 */
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

.nav-item button {
  text-align: left;
  width: 100%;
  background: none;
  border: none;
  padding: 10px;
  color: #007bff;
  font-size: 16px;
}

.nav-item button:hover {
  text-decoration: none;
  background-color: #f1f1f1;
}

/* 右侧区域样式 */
.flex-fill {
  flex: 1;
}

h2 {
  font-size: 1.5rem;
  margin-bottom: 20px;
}

p {
  font-size: 1rem;
}

/* 样式控制分割线 */
.bg-dark {
  background-color: #343a40 !important;
}

.text-white {
  color: #ffffff !important;
}

#app {
  margin-top: 0; /* 确保页面顶部与NavBar紧密相连 */
}

.d-flex {
  display: flex;
}

.flex-grow-1 {
  flex-grow: 1;
}

/* 菜单隐藏时右侧内容区域的宽度 */
.menu.d-none {
  width: 0 !important;
  padding: 0;
}

</style>
