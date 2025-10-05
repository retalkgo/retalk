# Docker 部署指南

## 快速开始

### 使用 Docker Compose（推荐）

这是最简单的部署方式，会自动启动 PostgreSQL 和 Redis。

1. 准备配置文件：
```bash
cp retalk.docker.yaml retalk.yaml
```

2. 编辑 `retalk.yaml`，至少需要修改：
   - `secret`: 修改为一个长且安全的密钥
   - 其他配置根据需要调整

3. 启动服务：
```bash
docker-compose up -d
```

4. 查看日志：
```bash
docker-compose logs -f retalk
```

5. 停止服务：
```bash
docker-compose down
```

### 仅使用 Docker

如果你已有 PostgreSQL 和 Redis 服务，可以只运行 retalk 容器：

1. 构建镜像：
```bash
docker build -t retalk:latest .
```

2. 准备配置文件 `retalk.yaml`，配置数据库和缓存连接信息

3. 运行容器：
```bash
docker run -d \
  --name retalk \
  -p 2716:2716 \
  -v $(pwd)/retalk.yaml:/app/retalk.yaml:ro \
  retalk:latest
```

## 配置说明

### 环境变量

当前版本通过配置文件 `retalk.yaml` 进行配置，需要挂载到容器的 `/app/retalk.yaml` 路径。

### 配置文件示例

参考 `retalk.docker.yaml` 文件，主要配置项：

- `server.host`: 建议设置为 `0.0.0.0` 以便容器外访问
- `server.port`: 默认 2716
- `database`: 数据库连接配置
  - 使用 Docker Compose 时，host 应设为 `postgres`
  - 使用外部数据库时，设置为实际的数据库地址
- `cache`: 缓存配置
  - 使用 Docker Compose 时，addr 应设为 `redis:6379`
  - 使用外部 Redis 时，设置为实际的 Redis 地址

## 数据持久化

Docker Compose 配置已经包含了数据卷：
- `postgres_data`: PostgreSQL 数据
- `redis_data`: Redis 数据

数据会在容器重启后保留，除非执行 `docker-compose down -v`（会删除卷）。

## 健康检查

服务启动后，可以访问健康检查接口：
```bash
curl http://localhost:2716/api/healthz
```

## 常见问题

### 端口冲突

如果 2716 端口已被占用，可以修改 `docker-compose.yml` 中的端口映射：
```yaml
ports:
  - "8080:2716"  # 将主机的 8080 端口映射到容器的 2716 端口
```

### 数据库连接失败

检查：
1. 数据库服务是否正常运行
2. `retalk.yaml` 中的数据库配置是否正确
3. 网络连接是否正常（Docker Compose 会自动创建网络）

### 查看容器日志

```bash
docker-compose logs -f retalk
```

或单独查看某个服务：
```bash
docker-compose logs -f postgres
docker-compose logs -f redis
```
