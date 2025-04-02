<template>
  <div class="mb-3">
    <label class="form-label">验证码</label>
    <div class="input-group">
      <input
        v-model="form.captcha"
        type="text"
        class="form-control"
        placeholder="输入验证码"
      >
      <img 
        :src="captchaImage" 
        @click="refreshCaptcha"
        class="captcha-image"
      >
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      captchaId: '',
      captchaImage: '',
      form: {
        // ...其他字段
        captcha: ''
      }
    }
  },
  created() {
    this.refreshCaptcha()
  },
  methods: {
    async refreshCaptcha() {
      try {
        const res = await this.$axios.get('/api/api/captcha')
        this.captchaId = res.data.captcha_id
        this.captchaImage = res.data.image_data
      } catch (error) {
        console.error('获取验证码失败:', error)
      }
    },
    async handleRegister() {
      const data = {
        captcha_id: this.captchaId,
        captcha_value: this.form.captcha
      }
      
      try {
        await this.$axios.post('/api/register', data)
        this.$router.push('/login')
      } catch (error) {
        this.refreshCaptcha() // 失败时刷新验证码
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
</style>