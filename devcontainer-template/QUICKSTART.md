# 代码加载快速参考

## 一键自动克隆（最简单）

**在 `.devcontainer/devcontainer.json` 中设置：**

```json
{
  "containerEnv": {
    "GIT_REPO_URL": "https://github.com/your-username/your-repo.git",
    "GIT_BRANCH": "main"
  }
}
```

启动容器即可自动克隆！ ✅

---

## 常用命令速查

### 容器运行中加载代码

```bash
# Git 克隆
git clone <repo-url> /workspace

# 从本地复制
docker cp /local/path devcontainer-app:/workspace/

# 从备份恢复
docker run --rm -v devcontainer-workspace:/data -v $(pwd):/backup \
  debian:trixie tar xzf /backup/backup.tar.gz -C /data
```

### 查看和管理 Volume

```bash
# 查看 Volume
docker volume ls
docker volume inspect devcontainer-workspace

# 进入 Volume 查看
docker run --rm -it -v devcontainer-workspace:/data debian:trixie /bin/bash

# 备份 Volume
docker run --rm -v devcontainer-workspace:/data -v $(pwd):/backup \
  debian:trixie tar czf /backup/workspace-backup.tar.gz -C /data

# 删除 Volume（清空代码）
docker volume rm devcontainer-workspace
```

### 切换到绑定挂载（本地开发）

**修改 `.devcontainer/docker-compose.yml`：**

```yaml
volumes:
  # 注释这行
  # - devcontainer-workspace:/workspace

  # 添加这行
  - .:/workspace
```

---

## 场景推荐

| 场景 | 推荐方式 | 配置 |
|------|---------|------|
| **新项目** | 自动克隆 | 设置 `GIT_REPO_URL` |
| **本地项目** | 绑定挂载 | 修改为 `.:/workspace` |
| **远程开发** | Volume + 手动克隆 | 使用默认配置 |
| **大型项目** | Volume + tar 归档 | `docker cp` + `tar` |
| **团队协作** | Volume + Git | 团队共享同一个 Volume |

---

## 完整示例

### 场景：从 GitHub 克隆现有项目

1. **创建项目目录**
   ```bash
   mkdir my-app && cd my-app
   ```

2. **添加 devcontainer 配置**
   ```bash
   cp -r /path/to/devcontainer-template .devcontainer
   ```

3. **配置自动克隆**
   ```json
   // .devcontainer/devcontainer.json
   {
     "containerEnv": {
       "GIT_REPO_URL": "https://github.com/username/my-app.git",
       "GIT_BRANCH": "main"
     }
   }
   ```

4. **在 VS Code 中启动**
   - 打开 VS Code
   - `F1` -> `Dev Containers: Reopen in Container`
   - 等待自动克隆完成 ✅

---

## 故障排除

### 容器启动但 /workspace 是空的？

**原因**：没有配置代码加载方式

**解决**：选择以下任一方式：
1. 设置 `GIT_REPO_URL` 环境变量
2. 手动在容器中 `git clone`
3. 使用 `docker cp` 复制代码
4. 修改为绑定挂载

### 找不到 devcontainer-app 容器？

**原因**：容器名称可能不同

**解决**：
```bash
# 查看实际容器名
docker ps

# 使用实际容器名
docker cp /local/path actual-container-name:/workspace/
```

### Volume 权限问题？

**解决**：
```bash
# 修正权限
docker exec -it devcontainer-app \
  sudo chown -R admin:admin /workspace
```

---

## 更多帮助

详细文档：[USAGE.md](./USAGE.md)
主文档：[README.md](./README.md)
