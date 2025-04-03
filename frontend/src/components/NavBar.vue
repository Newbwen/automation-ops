<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container">
      <!-- 新增：用户管理下拉菜单（左侧） -->
      <div class="dropdown" v-if="isLoggedIn">
        <button
          class="btn btn-dark dropdown-toggle"
          type="button"
          id="userDropdown"
          data-bs-toggle="dropdown"
          aria-expanded="false"
        >
          <i class="bi bi-person-circle me-1"></i> 个人中心
        </button>
        <ul class="dropdown-menu" aria-labelledby="userDropdown">
          <li><router-link class="dropdown-item" to="/profile">个人信息</router-link></li>
          <li><hr class="dropdown-divider" /></li>
          <li><a class="dropdown-item" href="#" @click.prevent="openChangePassword">修改密码</a></li>
          <li><a class="dropdown-item" href="#" @click.prevent="logout">退出登录</a></li>
        </ul>
      </div>

      <!-- 原有品牌链接（居中） -->
      <router-link class="navbar-brand mx-auto" to="/">自动化运维管理平台</router-link>

      <!-- 原有登录/注册按钮（右侧） -->
      <div class="navbar-nav">
        <router-link v-if="!isLoggedIn" class="nav-link" to="/login">登录</router-link>
        <router-link v-if="!isLoggedIn" class="nav-link" to="/register">注册</router-link>
      </div>
    </div>
  </nav>

  <!-- 修改密码模态框 -->
  <div class="modal fade" id="changePasswordModal" tabindex="-1" aria-hidden="true">
    <div class="modal-dialog">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">修改密码</h5>
          <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
        </div>
        <div class="modal-body">
          <form @submit.prevent="handleChangePassword">
            <div class="mb-3">
              <label class="form-label">旧密码</label>
              <input type="password" class="form-control" v-model="oldPassword" required />
            </div>
            <div class="mb-3">
              <label class="form-label">新密码</label>
              <input type="password" class="form-control" v-model="newPassword" required />
            </div>
            <div class="mb-3">
              <label class="form-label">确认密码</label>
              <input type="password" class="form-control" v-model="confirmPassword" required />
            </div>
            <button type="submit" class="btn btn-primary">确认修改</button>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { Modal } from 'bootstrap'; // 引入 Bootstrap 的 Modal 控制

export default {
  data() {
    return {
      isLoggedIn: !!localStorage.getItem('authToken'),
      oldPassword: '',
      newPassword: '',
      confirmPassword: ''
    };
  },
  methods: {
    // 打开修改密码模态框
    openChangePassword() {
      const modal = new Modal(document.getElementById('changePasswordModal'));
      modal.show();
    },
    // 提交修改密码
    async handleChangePassword() {
      try {
        //const token = localStorage.getItem('authToken');
        //验证新密码和确认密码是否一致
        if (this.newPassword !== this.confirmPassword) {
          alert('新密码和确认密码不一致，请重新输入！');
          return;
        }
        await this.$axios.post('/api/change-password', {
          old_password: this.oldPassword,
          new_password: this.newPassword,
          confirm_password: this.confirmPassword
        },)
        //{ headers: { 'Authorization': `Bearer ${token}` } });
        alert('密码修改成功！');
        // 关闭模态框
        const modal = Modal.getInstance(document.getElementById('changePasswordModal'));
        modal.hide();
      } catch (error) {
        alert('密码修改失败：' + error.response?.data?.message || error.message);
      }
      //修改完密码后跳转到登录页面
      localStorage.removeItem('authToken');
      this.isLoggedIn = false;
      this.$router.push('/login');
    },
    // 退出登录
    logout() {
      localStorage.removeItem('authToken');
      this.isLoggedIn = false;
      this.$router.push('/login');
    }
  }
};
</script>

<style scoped>
/* 调整下拉菜单样式 */
.dropdown-toggle::after {
  margin-left: 0.5em;
}
</style>