<template>
  <div class="host-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <h3>主机列表</h3>
          <div>
            <el-input
              v-model="searchQuery"
              placeholder="搜索主机名称或IP"
              style="width: 240px; margin-right: 10px"
              clearable
              @clear="handleSearchClear"
              @keyup.enter="fetchHosts"
            />
            <el-button type="primary" @click="handleCreate">新增主机</el-button>
          </div>
        </div>
      </template>

      <!-- 主机列表表格 -->
      <el-table
        :data="tableData"
        v-loading="loading"
        border
        stripe
        style="width: 100%"
      >
        <el-table-column prop="name" label="主机名称" />
        <el-table-column prop="ip" label="IP地址" />
        <el-table-column prop="port" label="端口" width="100" />
        <el-table-column prop="status" label="状态" />
        <el-table-column prop="os" label="操作系统" />
        <el-table-column prop="created_at" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.created_at) }}
          </template>
        </el-table-column>
        <el-table-column label="操作" width="180">
          <template #default="{ row }">
            <el-button size="small" @click="handleEdit(row)">编辑</el-button>
            <el-button size="small" type="danger" @click="handleDelete(row.ID)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页组件 -->
      <el-pagination
        class="pagination"
        :current-page="currentPage"
        :page-size="pageSize"
        :page-sizes="[10, 20, 50, 100]"
        :total="total"
        layout="total, sizes, prev, pager, next, jumper"
        @size-change="handleSizeChange"
        @current-change="handlePageChange"
      />
    </el-card>

    <!-- 主机表单对话框 -->
    <el-dialog
      :title="dialogTitle"
      v-model="dialogVisible"
      width="600px"
      @closed="handleDialogClosed"
    >
      <el-form
        ref="hostFormRef"
        :model="hostForm"
        :rules="formRules"
        label-width="100px"
      >
        <el-form-item label="主机名称" prop="name">
          <el-input v-model="hostForm.name" placeholder="请输入主机名称" />
        </el-form-item>
        <el-form-item label="IP地址" prop="ip">
          <el-input v-model="hostForm.ip" placeholder="请输入IP地址" />
        </el-form-item>
        <el-form-item label="端口" prop="port">
          <el-input v-model.number="hostForm.port" placeholder="请输入端口号" />
        </el-form-item>
        <el-form-item label="操作系统" prop="os">
          <el-select v-model="hostForm.os" placeholder="请选择操作系统" style="width: 100%">
            <el-option label="Linux" value="Linux" />
            <el-option label="Windows" value="Windows" />
            <el-option label="MacOS" value="MacOS" />
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">确认</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'HostList',
  data() {
    return {
      loading: false,
      tableData: [],
      currentPage: 1,
      pageSize: 10,
      total: 0,
      searchQuery: '',
      dialogVisible: false,
      dialogTitle: '新增主机',
      currentHostId: null,
      hostForm: {
        name: '',
        ip: '',
        port: 22,
        os: ''
      },
      formRules: {
        name: [
          { required: true, message: '请输入主机名称', trigger: 'blur' },
          { min: 2, max: 50, message: '长度在2到50个字符', trigger: 'blur' }
        ],
        ip: [
          { required: true, message: '请输入IP地址', trigger: 'blur' },
          { 
            pattern: /^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$/,
            message: '请输入正确的IP地址格式',
            trigger: 'blur'
          }
        ],
        port: [
          { required: true, message: '请输入端口号', trigger: 'blur' },
          { type: 'number', message: '端口必须为数字', trigger: 'blur' },
          { type: 'number', min: 1, max: 65535, message: '端口范围1-65535', trigger: 'blur' }
        ],
        os: [
          { required: true, message: '请选择操作系统', trigger: 'change' }
        ]
      }
    }
  },
  created() {
    this.fetchHosts()
  },
  methods: {
    async fetchHosts() {
      this.loading = true
      try {
        const params = {
          page: this.currentPage,
          pageSize: this.pageSize,
          search: this.searchQuery
        }
        
        const response = await this.$axios.get('/api/host/list', { params })
        this.tableData = response.data.data
        this.total = response.data.total
      } catch (error) {
        this.$message.error('获取主机列表失败')
      } finally {
        this.loading = false
      }
    },
    
    handleCreate() {
      this.dialogTitle = '新增主机'
      this.currentHostId = null
      this.dialogVisible = true
    },
    
    handleEdit(row) {
      this.dialogTitle = '编辑主机'
      this.currentHostId = row.ID
      this.hostForm = {
        name: row.name,
        ip: row.ip,
        port: row.port,
        os: row.os
      }
      this.dialogVisible = true
    },
    
    async handleDelete(ID) {
      try {
        await this.$confirm('确定要删除该主机吗?', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        })
        
        await this.$axios.delete(`/api/host/delete/${ID}`)
        this.$message.success('删除成功')
        this.fetchHosts()
      } catch (error) {
        // 用户取消删除
      }
    },
    
    handleSizeChange(val) {
      this.pageSize = val
      this.fetchHosts()
    },
    
    handlePageChange(val) {
      this.currentPage = val
      this.fetchHosts()
    },
    
    handleSearchClear() {
      this.searchQuery = ''
      this.currentPage = 1
      this.fetchHosts()
    },
    
    submitForm() {
      this.$refs.hostFormRef.validate(async (valid) => {
        if (valid) {
          try {
            if (this.currentHostId) {
              // 更新主机
              await this.$axios.put(`/api/host/edit/${this.currentHostId}`, this.hostForm)
              this.$message.success('更新主机成功')
              //this.$router.push('/api/host/list')
            } else {
              // 新增主机
              await this.$axios.post('/api/host/add', this.hostForm)
              this.$message.success('新增主机成功')
            }
            this.dialogVisible = false
            this.fetchHosts()
          } catch (error) {
            this.$message.error(error.response?.data?.message || '操作失败')
          }
        }
      })
    },
    
    handleDialogClosed() {
      this.$refs.hostFormRef.resetFields()
      this.currentHostId = null
      this.hostForm = {
        name: '',
        ip: '',
        port: 22,
        os: ''
      }
    },
    
    formatDate(dateString) {
      if (!dateString) return ''
      const date = new Date(dateString)
      return date.toLocaleString()
    }
  }
}
</script>

<style scoped>
.host-container {
  padding: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.pagination {
  margin-top: 20px;
  justify-content: flex-end;
}

.box-card {
  margin-bottom: 20px;
}
</style>