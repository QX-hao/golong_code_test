# 部署指南

本文档详细说明如何部署数据可视化管理平台到不同环境。

## 本地开发环境部署

### 1. 环境准备

确保系统已安装以下软件：
- Go 1.22+
- MySQL 5.7+ 或 8.0+
- Git

### 2. 数据库初始化

```bash
# 登录MySQL
mysql -u root -p

# 执行初始化脚本
source scripts/init.sql
```

### 3. 配置修改

编辑 `config/config.yaml` 文件，更新数据库连接信息：

```yaml
database:
  default:
    host: 127.0.0.1
    port: 3306
    user: root
    pwd: your_password
    name: data_visualization
```

### 4. 启动应用

**方式一：使用启动脚本（推荐）**
```bash
# Windows\.\run.bat

# Linux/macOS
chmod +x run.sh
./run.sh
```

**方式二：手动启动**
```bash
# 安装依赖
go mod tidy

# 编译
go build -o data-visualization-platform

# 运行
./data-visualization-platform
```

### 5. 访问应用

- 主页面: http://localhost:9033
- 管理后台: http://localhost:9033/admin
- API接口: http://localhost:9033/api

## Docker 部署

### 1. 环境准备

确保系统已安装：
- Docker 20.10+
- Docker Compose 2.0+

### 2. 一键部署

```bash
# 使用Docker Compose部署
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f app
```

### 3. 自定义配置

编辑 `docker-compose.yml` 文件，根据需要修改：

```yaml
services:
  mysql:
    environment:
      MYSQL_ROOT_PASSWORD: your_secure_password
      MYSQL_DATABASE: data_visualization
      MYSQL_USER: your_username
      MYSQL_PASSWORD: your_password

  app:
    environment:
      - DB_HOST=mysql
      - DB_PORT=3306
      - DB_USER=your_username
      - DB_PASSWORD=your_password
      - DB_NAME=data_visualization
```

### 4. 访问应用

- 主页面: http://localhost
- 管理后台: http://localhost/admin
- API接口: http://localhost/api

## 生产环境部署

### 1. 服务器准备

**系统要求：**
- 操作系统: Ubuntu 20.04+ / CentOS 8+
- 内存: 4GB+
- 存储: 50GB+
- 网络: 公网IP，开放80/443端口

### 2. 手动部署

```bash
# 创建应用用户
sudo useradd -m -s /bin/bash appuser
sudo passwd appuser

# 切换到应用用户
sudo su - appuser

# 克隆代码
git clone <repository-url>
cd data-visualization-platform

# 安装依赖
go mod tidy

# 编译生产版本
GOOS=linux GOARCH=amd64 go build -o data-visualization-platform

# 创建服务配置
sudo tee /etc/systemd/system/data-visualization.service << EOF
[Unit]
Description=Data Visualization Platform
After=network.target

[Service]
Type=simple
User=appuser
Group=appuser
WorkingDirectory=/home/appuser/data-visualization-platform
ExecStart=/home/appuser/data-visualization-platform/data-visualization-platform
Restart=always
RestartSec=10

[Install]
WantedBy=multi-user.target
EOF

# 启动服务
sudo systemctl daemon-reload
sudo systemctl enable data-visualization
sudo systemctl start data-visualization
sudo systemctl status data-visualization
```

### 3. Nginx 配置

```bash
# 安装Nginx
sudo apt update
sudo apt install nginx

# 配置站点
sudo tee /etc/nginx/sites-available/data-visualization << EOF
server {
    listen 80;
    server_name your-domain.com;
    
    location / {
        proxy_pass http://127.0.0.1:9033;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
EOF

# 启用站点
sudo ln -s /etc/nginx/sites-available/data-visualization /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 4. SSL证书配置

```bash
# 安装Certbot
sudo apt install certbot python3-certbot-nginx

# 获取证书
sudo certbot --nginx -d your-domain.com

# 设置自动续期
sudo crontab -e
# 添加：0 12 * * * /usr/bin/certbot renew --quiet
```

## 数据库备份与恢复

### 1. 备份数据库

```bash
# 备份整个数据库
mysqldump -u root -p data_visualization > backup_$(date +%Y%m%d).sql

# 备份特定表
mysqldump -u root -p data_visualization users business_data > tables_backup.sql
```

### 2. 恢复数据库

```bash
# 恢复整个数据库
mysql -u root -p data_visualization < backup_file.sql

# 恢复特定表
mysql -u root -p data_visualization < tables_backup.sql
```

### 3. 自动备份脚本

创建自动备份脚本 `/opt/backup.sh`：

```bash
#!/bin/bash
BACKUP_DIR="/opt/backups"
DATE=$(date +%Y%m%d_%H%M%S)

mysqldump -u root -p$DB_PASSWORD data_visualization > $BACKUP_DIR/backup_$DATE.sql

# 保留最近7天的备份
find $BACKUP_DIR -name "*.sql" -mtime +7 -delete
```

设置定时任务：
```bash
# 每天凌晨2点执行备份
0 2 * * * /opt/backup.sh
```

## 监控与日志

### 1. 应用日志

日志文件位置：
- 访问日志: `logs/access.log`
- 错误日志: `logs/error.log`
- 系统日志: `logs/system.log`

### 2. 系统监控

使用以下命令监控系统状态：

```bash
# 查看应用进程
ps aux | grep data-visualization

# 查看端口占用
netstat -tlnp | grep 9033

# 查看系统资源
top
htop

# 查看磁盘使用
df -h

# 查看内存使用
free -h
```

### 3. 性能优化

**数据库优化：**
```sql
-- 添加索引
CREATE INDEX idx_business_data_category ON business_data(category);
CREATE INDEX idx_business_data_created_at ON business_data(created_at);

-- 优化查询
ANALYZE TABLE business_data;
OPTIMIZE TABLE business_data;
```

**应用优化：**
- 启用Gzip压缩
- 配置CDN加速静态资源
- 使用Redis缓存热点数据
- 优化数据库连接池配置

## 故障排除

### 常见问题

1. **应用无法启动**
   - 检查端口是否被占用：`netstat -tlnp | grep 9033`
   - 检查配置文件路径和权限
   - 查看应用日志：`tail -f logs/error.log`

2. **数据库连接失败**
   - 检查数据库服务是否运行：`systemctl status mysql`
   - 验证连接配置：`mysql -u username -p -h hostname`
   - 检查防火墙设置

3. **静态资源无法访问**
   - 确认static目录存在且权限正确
   - 检查Nginx配置中的静态文件路径
   - 验证文件权限：`ls -la static/`

4. **性能问题**
   - 检查数据库查询性能
   - 监控系统资源使用情况
   - 优化前端资源加载

### 紧急恢复

1. **服务宕机**
   ```bash
   # 重启服务
   sudo systemctl restart data-visualization
   
   # 查看状态
   sudo systemctl status data-visualization
   sudo journalctl -u data-visualization -f
   ```

2. **数据库故障**
   ```bash
   # 重启数据库
   sudo systemctl restart mysql
   
   # 检查数据库状态
   mysql -u root -p -e "SHOW DATABASES;"
   ```

## 安全建议

1. **定期更新**
   - 及时更新操作系统安全补丁
   - 定期更新应用依赖包
   - 监控安全公告

2. **访问控制**
   - 使用强密码策略
   - 限制数据库远程访问
   - 配置防火墙规则

3. **数据保护**
   - 定期备份重要数据
   - 加密敏感信息
   - 实施访问审计

4. **监控告警**
   - 设置系统监控告警
   - 监控异常访问模式
   - 定期安全扫描

---

如有问题，请查看详细日志文件或联系技术支持。