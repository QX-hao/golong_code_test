# 数据可视化管理平台

基于 Go-Admin 框架构建的专业数据可视化与管理系统，提供强大的数据管理、图表展示和仪表板功能。

## 功能特性

- 📊 **多维度数据可视化** - 支持折线图、柱状图、饼图、雷达图等多种图表类型
- 🎯 **实时数据监控** - 实时数据更新和监控，支持自动刷新
- 🔐 **权限管理** - 基于角色的访问控制（RBAC）
- 📈 **仪表板定制** - 可自定义的仪表板布局和图表组合
- 📱 **响应式设计** - 支持桌面和移动设备访问
- 🔄 **数据导入导出** - 支持 CSV、Excel 等格式数据导入
- 📋 **系统日志** - 完整的操作日志记录和审计

## 技术栈

### 后端
- **框架**: Go-Admin + Gin
- **数据库**: MySQL
- **ORM**: XORM
- **认证**: JWT + Session

### 前端
- **UI框架**: Bootstrap 5
- **图表库**: Chart.js
- **图标**: Bootstrap Icons

## 快速开始

### 环境要求

- Go 1.22+
- MySQL 5.7+
- Git

### 安装步骤

1. **克隆项目**
```bash
git clone <项目地址>
cd data-visualization-platform
```

2. **安装依赖**
```bash
go mod tidy
```

3. **数据库配置**
```bash
# 创建数据库
mysql -u root -p < scripts/init.sql
```

4. **修改配置文件**
编辑 `config/config.yaml`，更新数据库连接信息：
```yaml
database:
  default:
    host: 127.0.0.1
    port: 3306
    user: root
    pwd: your_password
    name: data_visualization
```

5. **启动服务**
```bash
go run main.go
```

6. **访问系统**
- 管理后台: http://localhost:9033/admin
- API文档: http://localhost:9033/api
- 前端页面: http://localhost:9033/

## 默认账号

- 管理员: admin / admin123
- 普通用户: user1 / user123

## 项目结构

```
data-visualization-platform/
├── config/                 # 配置文件
│   └── config.yaml        # 应用配置
├── handlers/              # 请求处理器
│   └── handlers.go        # API接口实现
├── models/                # 数据模型
│   └── models.go          # 数据库模型定义
├── routes/                # 路由配置
│   └── routes.go          # 路由定义
├── scripts/               # 脚本文件
│   └── init.sql           # 数据库初始化脚本
├── static/                # 静态资源
├── templates/             # 模板文件
│   ├── index.html         # 主页面
│   └── dashboard.html     # 仪表板页面
├── uploads/               # 文件上传目录
├── go.mod                 # 依赖管理
├── go.sum                 # 依赖校验
└── main.go               # 程序入口
```

## API 接口

### 数据管理
- `GET /api/data` - 获取业务数据列表
- `POST /api/data` - 创建业务数据
- `PUT /api/data/:id` - 更新业务数据
- `DELETE /api/data/:id` - 删除业务数据

### 图表管理
- `GET /api/charts` - 获取图表列表
- `POST /api/charts` - 创建图表配置
- `PUT /api/charts/:id` - 更新图表配置
- `DELETE /api/charts/:id` - 删除图表配置

### 统计接口
- `GET /api/stats/summary` - 获取统计摘要
- `GET /api/stats/trends` - 获取趋势数据
- `GET /api/stats/categories` - 获取分类统计

### 文件上传
- `POST /api/upload` - 文件上传
- `POST /api/upload/csv` - CSV文件上传

## 数据模型

### 用户表 (users)
- 用户基本信息管理
- 角色权限控制

### 业务数据表 (business_data)
- 存储各类业务数据
- 支持分类和标签

### 图表配置表 (chart_configs)
- 图表类型和配置存储
- 数据源定义

### 仪表板表 (dashboards)
- 仪表板布局配置
- 图表组合管理

## 开发指南

### 添加新的数据源
1. 在 `models/models.go` 中定义新的数据模型
2. 在 `handlers/handlers.go` 中实现对应的 API 接口
3. 在 `routes/routes.go` 中注册路由
4. 在前端页面中添加对应的展示逻辑

### 自定义图表类型
1. 在 `chart_configs` 表中添加新的图表类型
2. 在前端 JavaScript 中实现图表渲染逻辑
3. 更新配置界面支持新的图表类型

### 扩展权限控制
1. 修改用户表的角色字段
2. 在中间件中添加权限验证逻辑
3. 更新前端菜单和按钮的权限控制

## 部署说明

### 生产环境部署

1. **编译应用**
```bash
go build -o data-visualization-platform
```

2. **配置生产环境**
创建 `config/production.yaml`:
```yaml
app:
  env: production
  debug: false
  log_level: warn

server:
  addr: 0.0.0.0
  port: 9033
```

3. **使用进程管理**
使用 systemd 或 supervisor 管理进程：
```ini
[program:data-visualization]
command=/path/to/data-visualization-platform
directory=/path/to/app
autostart=true
autorestart=true
```

### Docker 部署

```dockerfile
FROM golang:1.22-alpine

WORKDIR /app
COPY . .
RUN go build -o main .

EXPOSE 9033
CMD ["./main"]
```

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查数据库服务是否启动
   - 验证连接配置是否正确
   - 确认数据库用户权限

2. **端口被占用**
   - 修改配置文件中的端口号
   - 检查是否有其他服务占用端口

3. **静态资源无法加载**
   - 确认 static 目录存在
   - 检查文件权限设置

### 日志查看

应用日志位于 `logs/` 目录，包含：
- 访问日志
- 错误日志
- 系统日志

## 贡献指南

1. Fork 项目
2. 创建特性分支
3. 提交更改
4. 推送到分支
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证。

## 联系方式

- 项目主页: [项目地址]
- 问题反馈: [Issues]
- 文档地址: [Wiki]

---

**数据可视化管理平台** - 让数据洞察更简单！