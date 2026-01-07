# VS Code Dev Container 配置模板

这是一个完整的 VS Code Dev Container 配置模板，使用 Docker Compose 并将代码存储在独立的 Docker Volume 中。

## 特性

- ✅ 使用 Docker Compose 管理容器
- ✅ 代码存储在独立的 Docker Volume 中（持久化）
- ✅ 支持自定义项目依赖
- ✅ 预装常用 VS Code 扩展
- ✅ 自动配置开发环境

## 使用方法

### 1. 复制配置到你的项目

将此目录复制到你的项目根目录，重命名为 `.devcontainer`：

```bash
# 在你的项目根目录执行
cp -r /path/to/devcontainer-template .devcontainer
```

### 2. 调整配置（可选）

根据项目需求，可以修改以下文件：

#### `docker-compose.yml`
- 修改端口映射
- 调整资源限制
- 添加环境变量

#### `devcontainer.json`
- 修改 VS Code 扩展列表
- 调整端口转发设置
- 自定义编辑器设置

### 3. 配置项目依赖

创建 `.devcontainer-dependencies` 文件来定义项目需要的语言环境：

```bash
# 复制示例文件
cp .devcontainer/.devcontainer-dependencies.example .devcontainer-dependencies

# 编辑文件，取消需要的环境的注释
vim .devcontainer-dependencies
```

示例内容：

```bash
#!/usr/bin/env bash

# 安装 Node.js
bash ~/scripts/nvm.sh

# 安装 Go
bash ~/scripts/gvm.sh

# 安装 Rust
bash ~/scripts/rustup.sh
```

### 4. 启动 Dev Container

在 VS Code 中：
1. 按 `F1` 打开命令面板
2. 选择 `Dev Containers: Reopen in Container`
3. 等待容器构建和配置完成

## Volume 管理

### 查看代码 Volume

```bash
docker volume inspect devcontainer-workspace
```

### 进入 Volume 查看代码

```bash
docker run --rm -it -v devcontainer-workspace:/data debian:trixie /bin/bash
cd /data
ls -la
```

### 备份代码

```bash
# 从 Volume 备份到本地
docker run --rm -v devcontainer-workspace:/data -v $(pwd):/backup debian:trixie \
  tar czf /backup/workspace-backup.tar.gz -C /data .

# 从本地恢复到 Volume
docker run --rm -v devcontainer-workspace:/data -v $(pwd):/backup debian:trixie \
  tar xzf /backup/workspace-backup.tar.gz -C /data
```

### 清理 Volume

```bash
# 删除 Volume（会删除所有代码！）
docker volume rm devcontainer-workspace
```

## 目录结构

```
your-project/
├── .devcontainer/
│   ├── devcontainer.json           # VS Code Dev Container 配置
│   ├── docker-compose.yml          # Docker Compose 配置
│   ├── post-create.sh              # 容器创建后执行的脚本
│   ├── .devcontainer-dependencies  # 项目依赖配置（需要创建）
│   └── README.md                   # 本文档
└── your-code/                      # 你的项目代码会通过 Volume 映射到容器内的 /workspace
```

## 工作原理

1. **代码持久化**：使用 Docker Volume `devcontainer-workspace` 存储代码
2. **首次启动**：`post-create.sh` 会检查并安装系统依赖
3. **项目依赖**：通过 `.devcontainer-dependencies` 定义项目特定的环境
4. **VS Code 集成**：自动安装扩展、配置端口转发、设置环境变量

## 常见问题

### Q: 为什么使用 Volume 而不是直接挂载本地目录？

A: 使用 Volume 的优势：
- ✅ 隔离宿主机和容器环境
- ✅ 更好的性能（尤其是文件监听）
- ✅ 避免权限问题
- ✅ 可以在多个容器间共享

### Q: 如何在宿主机访问代码？

A: Volume 默认存储在 Docker 管理的目录中。如果需要同步到本地，可以：
1. 定期备份 Volume（见上方备份命令）
2. 使用 Git 将代码推送到远程仓库
3. 修改 `docker-compose.yml` 使用绑定挂载而非 Volume

### Q: 如何修改 Volume 中的代码？

A: 推荐的方式：
1. 通过 VS Code 连接到 Dev Container
2. 在 VS Code 中直接编辑（会实时同步到 Volume）
3. 使用终端在容器内操作

## 高级配置

### 添加数据库服务

在 `docker-compose.yml` 中添加：

```yaml
services:
  db:
    image: postgres:15
    environment:
      POSTGRES_PASSWORD: example
    volumes:
      - postgres-data:/var/lib/postgresql/data
    networks:
      - devcontainer-network

volumes:
  postgres-data:
```

### 使用 Docker in Docker

取消 `devcontainer.json` 中的注释：

```json
"mounts": [
  "source=/var/run/docker.sock,target=/var/run/docker.sock,type=bind"
]
```

### 自定义 VS Code 扩展

编辑 `devcontainer.json` 中的 `extensions` 数组：

```json
"extensions": [
  "dbaeumer.vscode-eslint",
  "your-extension-id.your-extension"
]
```

## 相关文档

- [VS Code Dev Containers 文档](https://code.visualstudio.com/docs/devcontainers/containers)
- [Docker Compose 文档](https://docs.docker.com/compose/)
- [本项目主文档](../README.md)
