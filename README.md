# 自动化运维管理平台

## 项目结构
```
automation-ops/
│-- backend/               # 后端代码（Golang + Gin + GORM）
│   │-- main.go            # 入口文件
│   │-- config/            # 配置管理
│   │-- controllers/       # 控制器
│   │-- models/           # 数据模型
│   │-- routes/           # 路由定义
│   │-- services/         # 业务逻辑
│   │-- database/         # 数据库连接
│-- frontend/              # 前端代码（Vue + Bootstrap）
│   │-- src/              # Vue 源代码
│   │   │-- components/    # 组件
│   │   │-- views/         # 页面视图
│   │   │-- router.js      # 路由配置
│   │-- public/           # 静态资源
│   │-- package.json      # 前端依赖管理
│   │-- vite.config.js    # Vue 开发配置
│-- deploy/                # 部署相关（Docker、K8s）
│-- docs/                  # 文档（Markdown）
│-- README.md              # 项目介绍
```

## 后端代码示例

### `main.go`
```go
package main

import (
    "automation-ops/config"
    "automation-ops/routes"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.Use(cors.Default()) // 允许跨域请求
    routes.SetupRoutes(r)
    r.Run(":8080")
}
```

### `routes/routes.go`
```go
package routes

import (
    "automation-ops/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    api := r.Group("/api")
    {
        api.GET("/hosts", controllers.GetHosts)
        api.GET("/k8s/clusters", controllers.GetK8sClusters)
        api.GET("/monitoring", controllers.GetMonitoringData)
    }
}
```

## 前端代码示例

### `src/router.js`
```javascript
import { createRouter, createWebHistory } from 'vue-router';
import Home from './views/Home.vue';
import Hosts from './views/Hosts.vue';

const routes = [
    { path: '/', component: Home },
    { path: '/hosts', component: Hosts }
];

const router = createRouter({
    history: createWebHistory(),
    routes
});

export default router;
```

### `src/views/Home.vue`
```vue
<template>
  <div>
    <h1>自动化运维管理平台</h1>
    <p>欢迎使用</p>
  </div>
</template>

<script>
export default {
  name: 'Home'
};
</script>
```

### `README.md`
```markdown
# 自动化运维管理平台

## 介绍
本项目是一个基于 Golang (Gin + GORM) 和 Vue.js 的前后端分离自动化运维管理平台，支持主机管理、K8s 集群管理、Prometheus 监控等功能。

## 主要功能
- **主机管理**: 远程操作、状态监控
- **K8s 集群管理**: 资源管理、监控告警
- **Prometheus 监控**: 指标采集、告警

## 运行方式

### 启动后端
```sh
git clone https://github.com/yourrepo/automation-ops.git
cd automation-ops/backend
go run main.go
```

### 启动前端
```sh
cd automation-ops/frontend
npm install
npm run dev
```

### 访问前端
```
打开浏览器访问：http://localhost:5173
```
```
