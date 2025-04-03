import { createApp } from 'vue'
import App from './App.vue'

import router from './router'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap'
import axios from 'axios'

const app = createApp(App)
app.config.globalProperties.$axios = axios//定义全局属性，其他文件使用方式this.$axios

axios.interceptors.request.use(config => {
  const token = localStorage.getItem('authToken');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

app.use(router).mount('#app')