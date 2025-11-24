# Go Starter - 用户认证系统

一个基于 Go 的用户认证 starter 项目，包含注册、登录和 JWT 认证功能。

## 技术栈

- **Web框架**: Gin
- **数据库**: PostgreSQL
- **ORM**: sqlx
- **认证**: JWT (JSON Web Token)
- **密码加密**: bcrypt

## 项目结构

```
go-starter/
├── cmd/
│   └── server/          # 应用程序入口
├── config/              # 配置管理
├── internal/
│   ├── handlers/        # HTTP 处理器
│   ├── middleware/      # 中间件
│   ├── models/          # 数据模型
│   ├── repository/      # 数据访问层
│   └── services/        # 业务逻辑层
├── pkg/
│   ├── database/        # 数据库连接和迁移
│   └── utils/           # 工具函数（JWT、密码加密）
└── migrations/          # SQL 迁移脚本
```

## 快速开始

### 1. 环境准备

确保已安装：
- Go 1.23+
- PostgreSQL

### 2. 安装依赖

```bash
go mod download
```

### 3. 配置数据库

创建 PostgreSQL 数据库：

```sql
CREATE DATABASE go_starter;
```

### 4. 配置环境变量

复制 `.env.example` 到 `.env` 并修改配置：

```bash
cp .env.example .env
```

编辑 `.env` 文件，设置数据库密码和 JWT 密钥：

```env
DB_PASSWORD=your_actual_password
JWT_SECRET=your-very-secret-key-min-32-characters-long
```

### 5. 运行项目

```bash
go run cmd/server/main.go
```

服务器将在 `http://localhost:8080` 启动。

## API 接口

### 健康检查

```
GET /health
```

### 用户注册

```
POST /api/v1/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "username": "username",
  "password": "password123"
}
```

响应：
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "email": "user@example.com",
    "username": "username",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### 用户登录

```
POST /api/v1/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

响应格式同注册接口。

### 获取用户信息（需要认证）

```
GET /api/v1/users/profile
Authorization: Bearer <token>
```

响应：
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "username",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## 测试示例

### 使用 curl

注册用户：
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","username":"testuser","password":"password123"}'
```

登录：
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"password123"}'
```

获取用户信息：
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## 环境变量说明

| 变量 | 说明 | 默认值 |
|------|------|--------|
| SERVER_PORT | 服务器端口 | 8080 |
| GIN_MODE | Gin 运行模式 (debug/release) | debug |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户 | postgres |
| DB_PASSWORD | 数据库密码 | 必填 |
| DB_NAME | 数据库名称 | go_starter |
| DB_SSLMODE | SSL 模式 | disable |
| JWT_SECRET | JWT 密钥 | 必填（生产环境） |

## 安全建议

1. **生产环境**：
   - 使用强密码和长随机字符串作为 JWT_SECRET
   - 将 GIN_MODE 设置为 `release`
   - 启用数据库 SSL (DB_SSLMODE=require)
   - 使用环境变量管理敏感信息，不要提交 .env 文件

2. **密码要求**：
   - 最小长度 6 个字符
   - 建议使用更强的密码策略

3. **Token 过期**：
   - 默认 24 小时过期
   - 可在 config/config.go 中修改

## 扩展功能

可以继续添加：
- 刷新 Token 机制
- 用户角色和权限管理
- 密码重置功能
- 邮箱验证
- OAuth 第三方登录
- API 限流
- 日志记录

## License

MIT
