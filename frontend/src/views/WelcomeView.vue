<template>
  <div class="container mt-5">
    <h1>欢迎回来, {{ user.email }}!</h1>
    <div class="card mt-4">
      <div class="card-body">
        <h5 class="card-title">用户信息</h5>
        <p class="card-text">用户名: {{ user.username }}</p>
        <button @click="logout" class="btn btn-danger">退出登录</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        username: '',
        email: '',
      }
    }
  },
  async created() {
    await this.fetchUserInfo();
  },
  methods: {
    async fetchUserInfo() {
      try {
        const response = await this.$axios.get('/api/user', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('authToken')}`
          }
        });
        this.user = response.data;
      } catch (error) {
        console.error('获取用户信息失败:', error);
        this.$router.push('/login');
      }
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleDateString();
    },
    logout() {
      localStorage.removeItem('authToken');
      this.$router.push('/login');
    }
  }
}
</script>
