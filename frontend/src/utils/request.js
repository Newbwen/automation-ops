import axios from 'axios'

const service = axios.create({
  baseURL: process.env.NODE_ENV === 'development' 
    ? '/api'  // 开发环境使用代理
    : 'https://your-domain.com/api', // 生产环境地址
  timeout: 5000
})

export default service