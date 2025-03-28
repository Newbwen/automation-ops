<template>
  <div class="container mt-5">
    <div class="row justify-content-center">
      <div class="col-md-6">
        <div class="card">
          <div class="card-header">
            <h3>注册</h3>
          </div>
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
      form: {
        email: '',
        password: '',
        confirmPassword: ''
      },
      errors: {
        email: '',
        password: '',
        confirmPassword: ''
      }
    }
  },
  methods: {
    validateForm() {
      let isValid = true
      
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      if (!this.form.email) {
        this.errors.email = 'Email is required'
        isValid = false
      } else if (!emailRegex.test(this.form.email)) {
        this.errors.email = 'Invalid email format'
        isValid = false
      } else {
        this.errors.email = ''
      }

      if (!this.form.password) {
        this.errors.password = 'Password is required'
        isValid = false
      } else if (this.form.password.length < 6) {
        this.errors.password = 'Password must be at least 6 characters'
        isValid = false
      } else {
        this.errors.password = ''
      }

      if (this.form.password !== this.form.confirmPassword) {
        this.errors.confirmPassword = 'Passwords do not match'
        isValid = false
      } else {
        this.errors.confirmPassword = ''
      }

      return isValid
    },
    handleRegister() {
      if (this.validateForm()) {
        this.$axios.post('/api/register', this.form)
         .then(response => {
            console.log(response)
            this.$router.push('/login')
          })
         .catch(error => {
            console.log(error)
            this.errors.email = error.response.data.email[0]
            this.errors.password = error.response.data.password[0]
            this.errors.confirmPassword = error.response.data.confirmPassword[0]
          })
      }
  }
}
}
</script>