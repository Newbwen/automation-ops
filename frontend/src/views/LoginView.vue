<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3>登录</h3>
          </div>
          <div class="card-body">
            <div class="error mb-3">
              {{ errMsg }}
            </div>
            <form @submit.prevent="handleLogin">
              <div class="mb-3">
                <label class="form-label">邮箱地址</label>
                <input 
                  type="email" 
                  class="form-control"
                  v-model="form.email"
                  :class="{ 'is-invalid': errors.email }"
                >
                <div class="invalid-feedback">{{ errors.email }}</div>
              </div>
              
              <div class="mb-3">
                <label class="form-label">密码</label>
                <input 
                  type="password" 
                  class="form-control"
                  v-model="form.password"
                  :class="{ 'is-invalid': errors.password }"
                >
                <div class="invalid-feedback">{{ errors.password }}</div>
              </div>

              <button type="submit" class="btn btn-primary">登录</button>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      form: {
        email: '',
        password: ''
      },
      errors: {
        email: '',
        password: ''
      },
      errMsg: ''
    }
  },
  methods: {
    validateForm() {
      let isValid = true
      
      // Email validation
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      if (!this.form.email) {
        this.errors.email = 'Email is required'
        isValid = false
      } else if (!emailRegex.test(this.form.email)) {
        this.errors.email = 'Invalid email format'
        isValid = false
      } else {
        this.errors.email = ""
      }

      // Password validation
      if (!this.form.password) {
        this.errors.password = 'Password is required'
        isValid = false
      } else {
        this.errors.password = ''
      }

      return isValid
    },
    handleLogin() {
      if (this.validateForm()) {
        this.$axios.post('/api/login', this.form)
        .then(response => {
          localStorage.setItem('authToken', response.data.token)
          //跳转到欢迎页
          this.$router.push({ 
          name: 'welcome',
          params: { freshLogin: true } // 可选参数，用于显示欢迎提示
        })
      })
          .catch(error => {
            // 明确处理不同错误类型
            if (error.response) {
              // HTTP 状态码 2xx 之外的情况
              this.errMsg = error.response?.data?.error || '用户名或者密码错误，登录失败！'
            } else if (error.request) {
              // 请求已发出但无响应
              this.errMsg = '服务器无响应'
            } else {
              // 其他错误
              this.errMsg = error.message
            }
          })
      }
    }
  }
}
</script>