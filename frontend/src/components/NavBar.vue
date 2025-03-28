<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <div class="container">
      <router-link class="navbar-brand" to="/">自动化运维管理平台</router-link>
      <div class="navbar-nav">
        <router-link 
          v-if="!isLoggedIn"
          class="nav-link" 
          to="/login"
        >登录</router-link>
        <router-link 
          v-if="!isLoggedIn"
          class="nav-link" 
          to="/register"
        >注册</router-link>
        <button 
          v-if="isLoggedIn"
          class="btn btn-link nav-link"
          @click="logout"
        >退出登录</button>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  data() {
    return {
      isLoggedIn: false
    }
  },
  methods: {
    logout() {
      localStorage.removeItem('authToken')
      this.isLoggedIn = false
      this.$router.push('/login')
    }
  }
}
</script>