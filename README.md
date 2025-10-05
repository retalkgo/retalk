# retalk

> 轻量简洁的评论系统，基于 hertz 构建

## 部署

### Docker 部署（推荐）

1. 复制配置文件并修改：
```bash
cp retalk.docker.yaml retalk.yaml
# 编辑 retalk.yaml，修改 secret 等配置
```

2. 使用 Docker Compose 启动：
```bash
docker-compose up -d
```

服务将在 `http://localhost:2716` 启动。

### 手动部署

1. 安装 PostgreSQL 和 Redis（可选）
2. 复制配置文件：
```bash
cp retalk.example.yaml retalk.yaml
# 编辑 retalk.yaml，配置数据库和缓存连接
```

3. 编译并运行：
```bash
make app
./bin/retalk
```

## 开发环境配置

```bash
make install-dev
```

## 运行

```bash
make dev
```