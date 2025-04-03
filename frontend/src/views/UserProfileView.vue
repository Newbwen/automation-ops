<template>
  <div class="container mt-4">
    <div class="row justify-content-center">
      <div class="col-md-8">
        <div class="card">
          <div class="card-header bg-primary text-white">
            <h4>个人信息</h4>
          </div>
          <div class="card-body">
            <!-- 个人信息展示 -->
            <div v-if="!isEditing" class="user-info">
              <div class="row mb-3">
                <div class="col-md-4 fw-bold">用户名:</div>
                <div class="col-md-8">{{ user.username }}</div>
              </div>
              <div class="row mb-3">
                <div class="col-md-4 fw-bold">邮箱:</div>
                <div class="col-md-8">{{ user.email }}</div>
              </div>
              <div class="row mb-3">
                <div class="col-md-4 fw-bold">角色:</div>
                <div class="col-md-8">{{ user.role }}</div>
              </div>
              <div class="row mb-3">
                <div class="col-md-4 fw-bold">注册时间:</div>
                <div class="col-md-8">{{ formatDate(user.created_at) }}</div>
              </div>
              <div class="now mb-3">
                <div class="col-md-4 fw-bold">上次登录时间:</div>
                <div class="col-md-8">{{ formatDate(user.last_login) }}</div>
              </div>
              <button @click="startEditing" class="btn btn-primary me-2">
                编辑信息
              </button>
            </div>

            <!-- 信息编辑表单 -->
            <div v-else class="edit-form">
              <form @submit.prevent="saveProfile">
                <div class="mb-3 row">
                  <label for="username" class="col-md-4 col-form-label fw-bold">用户名</label>
                  <div class="col-md-8">
                    <input type="text" class="form-control" id="username" v-model="editUser.username">
                  </div>
                </div>
                <div class="mb-3 row">
                  <label for="email" class="col-md-4 col-form-label fw-bold">邮箱</label>
                  <div class="col-md-8">
                    <input type="email" class="form-control" id="email" v-model="editUser.email">
                  </div>
                </div>
                <div class="text-end">
                  <button type="button" @click="cancelEditing" class="btn btn-outline-secondary me-2">
                    取消
                  </button>
                  <button type="submit" class="btn btn-primary" :disabled="isSaving">
                    <span v-if="isSaving" class="spinner-border spinner-border-sm" role="status" aria-hidden="true"></span>
                    {{ isSaving ? '保存中...' : '保存' }}
                  </button>
                </div>
              </form>
            </div>
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
      user: {
        username: '',
        email: '',
        role: '',
        createdAt: '',
        last_login: ''

      },
      editUser: {
        username: '',
        email: ''
      },
      isEditing: false,
      isSaving: false
    }
  },
  async created() {
    await this.fetchUserInfo()
  },
  methods: {
    async fetchUserInfo() {
      try {
        const response = await this.$axios.get('/api/user/profile', {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('authToken')}`
          }
        })
        this.user = response.data
      } catch (error) {
        console.error('获取用户信息失败:', error)
        this.$toast.error('获取用户信息失败')
      }
    },
    startEditing() {
      this.editUser = { ...this.user }
      this.isEditing = true
    },
    cancelEditing() {
      this.isEditing = false
    },
    async saveProfile() {
      this.isSaving = true
      try {
        const response = await this.$axios.put('/api/user/profile', this.editUser, {
          headers: {
            Authorization: `Bearer ${localStorage.getItem('authToken')}`
          }
        })
        this.user = response.data
        this.isEditing = false
        this.$toast.success('个人信息更新成功')
      } catch (error) {
        console.error('更新失败:', error)
        this.$toast.error(error.response?.data?.message || '更新失败')
      } finally {
        this.isSaving = false
      }
    },
    formatDate(dateString) {
      return new Date(dateString).toLocaleString()
    }
  }
}
</script>

<style scoped>
.user-info {
  padding: 15px;
}

.edit-form {
  padding: 15px;
}

.card {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.card-header {
  padding: 15px 20px;
}
</style>