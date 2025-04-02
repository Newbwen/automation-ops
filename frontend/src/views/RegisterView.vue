<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3>注册</h3>
          </div>
          <div v-if="errorMessage" class="alert alert-danger" role="alert">{{ errorMessage }}</div>
          <div v-if="loading" class="text-center"></div>
          <div class="card-body">
            <form @submit.prevent="handleRegister">
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

              <div class="mb-3">
                <label class="form-label">确认密码</label>
                <input 
                  type="password" 
                  class="form-control"
                  v-model="form.confirmPassword"
                  :class="{ 'is-invalid': errors.confirmPassword }"
                >
                <div class="invalid-feedback">{{ errors.confirmPassword }}</div>
              </div>

            <!--<div class="mb-3">
                <label class="form-label">验证码</label>
                <div class="input-group">
                  <input
                    v-model="form.captcha"
                    type="text"
                    class="form-control"
                    placeholder="输入验证码"
                    :class="{ 'is-invalid': errors.captcha }"
                  >
                  <img 
                    :src="captchaImage" 
                    @click="refreshCaptcha"
                    class="captcha-image"
                  >
                  <div class="invalid-feedback">{{ errors.captcha }}</div>
                </div>
              </div>
-->  
              <button type="submit" class="btn btn-primary">注册</button>
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
      //captchaId: '',
      //captchaImage: '',
      form: {
        email: '',
        password: '',
        confirmPassword: ''
        //captcha: ''
      },
      errors: {
        email: '',
        password: '',
        confirmPassword: ''
        //captcha: ''
      },
      errorMessage: '',
      loading: false
    }
  },
  methods: {
    // 表单验证方法
    validateForm() {
      let isValid = true
      this.errorMessage = ''

      // 邮箱验证
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      if (!this.form.email.trim()) {
        this.errors.email = '邮箱不能为空'
        isValid = false
      } else if (!emailRegex.test(this.form.email)) {
        this.errors.email = '邮箱格式不正确'
        isValid = false
      } else {
        this.errors.email = ''
      }

      // 密码验证
      if (!this.form.password) {
        this.errors.password = '密码不能为空'
        isValid = false
      } else if (this.form.password.length < 6) {
        this.errors.password = '密码至少需要6位'
        isValid = false
      } else {
        this.errors.password = ''
      }

      // 确认密码验证
      if (!this.form.confirmPassword) {
        this.errors.confirmPassword = '请确认密码'
        isValid = false
      } else if (this.form.password !== this.form.confirmPassword) {
        this.errors.confirmPassword = '两次输入的密码不一致'
        isValid = false
      } else {
        this.errors.confirmPassword = ''
      }
     
      return isValid
    },
    // 注册处理
    async handleRegister() {
      if (!this.validateForm()) return

      this.loading = true
      try {
        const response = await this.$axios.post('/api/register', {
          email: this.form.email,
          password: this.form.password,
          confirm_password: this.form.confirmPassword,
          captcha_id: this.captchaId,
          captcha_value: this.form.captcha
        })

        if (response.status === 201) {
          this.$router.push('/login?registration=success')
        }
      } catch (error) {
        this.refreshCaptcha()
        this.handleRegistrationError(error)
      } finally {
        this.loading = false
      }
    },

    // 错误处理
    handleRegistrationError(error) {
      if (error.response) {
        switch (error.response.status) {
          case 409:
            this.errorMessage = '该邮箱已被注册'
            break
          case 400:
            this.errorMessage = '请求参数错误'
            break
          default:
            this.errorMessage = '注册失败，请稍后重试'
        }
      } else if (error.request) {
        this.errorMessage = '无法连接到服务器'
      } else {
        this.errorMessage = '发生未知错误'
      }
    }
  }
}
</script>

<style scoped>
.captcha-image {
  height: 40px;
  margin-left: 10px;
  cursor: pointer;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.card {
  border-radius: 15px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.card-header {
  border-radius: 15px 15px 0 0 !important;
}

.btn-primary {
  background-color: #007bff;
  border: none;
  padding: 12px;
  font-size: 1.1rem;
  transition: background-color 0.3s;
}

.btn-primary:hover {
  background-color: #0056b3;
}

.invalid-feedback {
  display: block;
  margin-top: 5px;
}

.spinner-border {
  vertical-align: text-top;
}
</style>
