# 如何加载代码到 Volume

使用 Volume 存储代码时，有几种方式可以将代码加载到容器中。

## 方式 1: 通过环境变量自动克隆（推荐）

在 `devcontainer.json` 中配置 Git 仓库地址：

```json
{
  "containerEnv": {
    "NODE_ENV": "development",
    "GIT_REPO_URL": "https://github.com/username/repo.git",
    "GIT_BRANCH": "main"
  }
}
```

容器启动时会自动克隆代码到 `/workspace`。

## 方式 2: 在容器启动后手动克隆

### 步骤 1: 启动容器

在 VS Code 中按 `F1` -> `Dev Containers: Reopen in Container`

### 步骤 2: 在终端中克隆代码

```bash
# 克隆到 /workspace
git clone https://github.com/username/repo.git /workspace

# 如果需要切换分支
cd /workspace
git checkout feature-branch
```

## 方式 3: 使用 docker cp 从本地复制

如果你已经在本地有代码：

```bash
# 在宿主机终端执行
# 复制整个项目目录到容器
docker cp /path/to/local/project devcontainer-app:/workspace/

# 如果需要，可以进入容器调整权限
docker exec -it devcontainer-app chown -R admin:admin /workspace
```

## 方式 4: 使用 tar 归档（推荐用于大型项目）

```bash
# 在宿主机打包
tar czf project.tar.gz -C /path/to/local/project .

# 复制到容器并解压
docker cp project.tar.gz devcontainer-app:/tmp/
docker exec devcontainer-app tar xzf /tmp/project.tar.gz -C /workspace
docker exec devcontainer-app rm /tmp/project.tar.gz
```

## 方式 5: 修改 docker-compose.yml 使用绑定挂载

如果你希望直接编辑本地文件，可以修改 `docker-compose.yml`：

```yaml
services:
  app:
    image: ghcr.io/kyicy/devcontainer:latest
    volumes:
      # 注释掉 volume 方式
      # - devcontainer-workspace:/workspace

      # 使用绑定挂载
      - type: bind
        source: .  # 当前目录
        target: /workspace
```

**注意**：这种方式会直接挂载本地目录，不再使用 Volume 持久化。

## 方式 6: 使用 SSH 挂载（适用于远程开发）

如果代码在远程服务器上：

### 1. 在 devcontainer.json 中启用 SSH 挂载

```json
{
  "mounts": [
    "source=${localEnv:HOME}/.ssh,target=/home/admin/.ssh,readonly,type=bind,consistency=cached"
  ]
}
```

### 2. 在容器中使用 SSHFS

```bash
# 安装 sshfs
sudo apt-get install sshfs

# 挂载远程目录
mkdir -p /tmp/remote-code
sshfs user@remote-server:/path/to/code /tmp/remote-code

# 复制到 workspace
cp -r /tmp/remote-code/* /workspace/

# 卸载
fusermount -u /tmp/remote-code
```

## 方式 7: 从备份恢复（如果你有之前的备份）

```bash
# 从备份恢复到 Volume
docker run --rm \
  -v devcontainer-workspace:/data \
  -v $(pwd):/backup \
  debian:trixie \
  tar xzf /backup/workspace-backup.tar.gz -C /data
```

## 最佳实践建议

### 新项目推荐流程

```bash
# 1. 创建本地项目目录
mkdir my-project && cd my-project

# 2. 复制 devcontainer 配置
cp -r /path/to/devcontainer-template .devcontainer

# 3. 编辑 devcontainer.json，设置 Git 仓库
vim .devcontainer/devcontainer.json
# 设置: "GIT_REPO_URL": "https://github.com/username/new-repo.git"

# 4. 在 VS Code 中打开项目
code .

# 5. 启动 Dev Container（会自动克隆代码）
# F1 -> Dev Containers: Reopen in Container
```

### 现有项目推荐流程

```bash
# 1. 进入现有项目目录
cd existing-project

# 2. 复制 devcontainer 配置
cp -r /path/to/devcontainer-template .devcontainer

# 3. 修改 docker-compose.yml 使用绑定挂载
vim .devcontainer/docker-compose.yml
# 将: devcontainer-workspace:/workspace
# 改为: .:/workspace

# 4. 在 VS Code 中启动容器
# F1 -> Dev Containers: Reopen in Container
```

## Volume vs 绑定挂载选择

### 使用 Volume（devcontainer-workspace:/workspace）
✅ 优点：
- 隔离性好，不影响宿主机
- 性能更好（文件监听更快）
- 避免权限问题
- 适合远程开发

❌ 缺点：
- 需要手动加载代码
- 宿主机无法直接访问文件

**适用场景**：
- 远程开发（Codespaces、远程 SSH）
- 性能要求高的项目
- 需要容器间共享代码

### 使用绑定挂载（.:/workspace）
✅ 优点：
- 宿主机可以直接访问文件
- 可以使用本地工具（IDE、Git 客户端等）
- 配置简单

❌ 缺点：
- 性能较差（大量文件时）
- 可能有权限问题
- 依赖宿主机环境

**适用场景**：
- 本地开发
- 小型项目
- 需要频繁使用宿主机工具

## 常见问题

### Q: 如何在 Volume 和本地文件之间同步？

A: 有几种方式：
1. 使用 Git 推送到远程仓库
2. 定期备份 Volume（见方式 7）
3. 使用 rsync 同步

### Q: Volume 中的代码在哪里？

A: Volume 存储在 Docker 管理的目录中：
```bash
# 查看 Volume 详情
docker volume inspect devcontainer-workspace

# 查看内容（需要 root 权限）
sudo ls -la /var/lib/docker/volumes/devcontainer-workspace/_data
```

### Q: 如何清空 Volume 重新开始？

A:
```bash
# 停止并删除容器
docker-compose down

# 删除 Volume
docker volume rm devcontainer-workspace

# 重新启动（会创建新的空 Volume）
docker-compose up -d
```

### Q: 能同时在多个容器中使用同一个 Volume 吗？

A: 可以！这正是 Volume 的优势之一：
```yaml
services:
  app1:
    volumes:
      - devcontainer-workspace:/workspace

  app2:
    volumes:
      - devcontainer-workspace:/workspace
```

两个容器共享同一个代码目录。
