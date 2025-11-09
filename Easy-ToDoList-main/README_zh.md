# To-Do List API (待办事项清单 API)

这是一个使用 Go (Golang) 和 Gin 框架构建的、简单而强大的待办事项清单 API 服务器。它提供了用户管理（注册、登录）以及针对待-办事项的完整 CRUD（创建、读取、更新、删除）操作，并使用 JWT 进行安全认证。

## 功能特性

- **用户认证**: 基于 JWT (JSON Web Tokens) 的安全注册和登录系统。
- **待办事项 CRUD**: 对待办事项提供完整的创建、读取、更新和删除功能。
- **RESTful API**: 简洁且符合惯例的 RESTful API 设计。
- **结构化日志**: 详细的日志记录，便于监控和调试。
- **配置管理**: 使用 `config.toml` 文件进行外部化配置。
- **API 文档**: 集成 Swagger，方便进行 API 探索和测试。
- **分层架构**: 组织良好的项目结构 (`cmd`, `internal`, `pkg`)，易于维护。

## 快速开始

请按照以下说明在您的本地机器上启动并运行该项目。

### 环境要求

- [Go](https://golang.org/dl/) (推荐使用 1.18 或更高版本)
- 一个正在运行的数据库实例（项目默认配置为本地文件数据库，如 SQLite）。

### 安装与设置

1.  **克隆仓库:**
    ```sh
    git clone <your-repository-url>
    cd todolist
    ```

2.  **安装依赖:**
    项目使用 Go Modules。当您构建或运行项目时，依赖项将自动下载。您也可以手动下载它们：
    ```sh
    go mod tidy
    ```

3.  **配置应用程序:**
    - 将 `config.example.toml` 重命名为 `config.toml` (或创建您自己的 `config.toml` 文件)。
    - 在 `config.toml` 文件中更新您的数据库连接信息、JWT 密钥和服务器端口。

    ```toml
    # config.toml
    [Server]
    Port = "8848"

    [Database]
    Type = "sqlite3"
    DSN = "todo.db"

    [Jwt]
    SecretKey = "your-secret-key"
    Expire = 3600 # 过期时间（秒），例如：1小时
    ```

### 运行应用

在项目根目录中执行以下命令来运行 API 服务器：

```sh
go run cmd/v1/main.go
```

服务器将在您 `config.toml` 文件中指定的端口上启动 (例如, `localhost:8848`)。

## API 端点

所有 API 端点的基础路径为 `/api/v1`。

### Swagger 文档

要访问交互式 API 文档，请访问：
`http://localhost:8848/swagger/index.html`

### 用户管理

-   **注册新用户:**
    -   `POST /user/register`
    -   **请求体**: `{"name": "YourName", "account": "user_account", "password": "your_password"}`

-   **登录:**
    -   `POST /user/login`
    -   **请求体**: `{"account": "user_account", "password": "your_password"}`
    -   **返回**: 用于后续请求认证的 JWT 令牌。

-   **更新用户信息:**
    -   `PUT /user/update`
    -   **认证**: `Bearer <token>`
    -   **请求体**: `{"name": "NewName", "password": "new_password"}`

### 待办事项管理

*所有待办事项相关的端点都需要在请求头中包含 `Authorization: Bearer <token>`。*

-   **创建新的待办事项:**
    -   `POST /todos/add`
    -   **请求体**: `{"title": "我的任务", "description": "任务详情", "status": 0, "end_time": "2025-12-31T23:59:59Z"}`

-   **列出用户的所有待办事项:**
    -   `GET /todos/list`

-   **根据 ID 获取特定待办事项:**
    -   `GET /todos/get/:id`

-   **更新待办事项:**
    -   `PUT /todos/update`
    -   **请求体**: `{"title": "更新后的标题", "description": "更新后的描述", "status": 1}`

-   **根据 ID 删除待办事项:**
    -   `DELETE /todos/delete/:id`

## 项目结构

该项目遵循标准的 Go 应用布局：

-   `cmd/`: 主程序入口。
    -   `v1/main.go`: 启动 API 服务器的 main 函数。
-   `internal/`: 应用的私有代码，不能被其他项目导入。
    -   `database/`: 数据库初始化。
    -   `dto/`: 用于请求体的数据传输对象 (Data Transfer Objects)。
    -   `handler/`: 处理 HTTP 请求的 Gin 处理器。
    -   `middleware/`: Gin 中间件 (例如 JWT 认证)。
    -   `model/`: 数据库模型/实体。
    -   `repository/`: 用于数据库操作的数据访问层。
    -   `router/`: API 路由定义。
    -   `service/`: 业务逻辑层。
    -   `vo/`: 用于 API 响应的视图对象 (View Objects)。
-   `pkg/`: 可被外部项目导入的公共库代码。
    -   `jwtUtil/`: JWT 生成和解析工具。
    -   `logUtil/`: 日志记录器配置。
-   `docs/`: Swagger 文档文件。
-   `config.toml`: 应用配置文件。
-   `go.mod`: Go 模块定义和依赖项。

