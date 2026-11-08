# go-server-falsework

基于 [Gin](https://gin-gonic.com/) 的 Go Web 服务脚手架（falsework），内置配置加载、结构化日志、Trace ID、统一响应与错误处理，可作为新服务的起点。

## 功能特性

- **HTTP 服务**：基于 Gin v1.12，路由按业务分组注册。
- **YAML 配置**：支持内置默认值，再由 `config/config.yml` 覆盖。
- **结构化日志**：基于 Zap 输出 JSON 日志，同时写入终端和文件；按天切割，自动清理过期日志。
- **Trace ID**：为每个请求建立独立 Trace ID（优先复用客户端传入的 `X-Trace-ID` / `X-Request-ID`），并写入响应头与日志字段。
- **自定义上下文**：`app.Context` 封装了请求级 logger、Trace ID、业务码和响应数据。
- **统一响应**：`Api` 中间件统一封装成功/失败响应，并集中记录错误日志。
- **结构化错误**：`exceptions.AppError` 携带 HTTP 状态码、业务码和错误信息。
- **异常恢复**：handler 内发生 panic 时自动恢复并返回 500，避免进程崩溃。
- **开发体验**：Air 热重载 + Delve 远程调试，附 VS Code 启动配置。

## 技术栈

| 组件 | 说明 |
| --- | --- |
| [Gin](https://github.com/gin-gonic/gin) | HTTP Web 框架 |
| [Zap](https://github.com/uber-go/zap) | 结构化日志 |
| [file-rotatelogs](https://github.com/lestrrat-go/file-rotatelogs) | 日志文件按天切割、过期清理 |
| [go-yaml](https://github.com/goccy/go-yaml) | YAML 配置解析 |
| [Air](https://github.com/air-verse/air) | 开发期热重载 |
| [Delve](https://github.com/go-delve/delve) | 调试器（Air 以 headless 模式启动，VS Code attach） |

## 目录结构

```text
.
├── cmd/server/main.go          # 程序入口：加载配置、注册路由、启动 HTTP 服务
├── config/
│   ├── example.config.yml      # 配置模板（纳入版本管理）
│   └── config.yml              # 本地实际配置（已 gitignore，需自行复制）
├── internal/
│   ├── app/                    # 应用装配
│   │   ├── config.go           # 全局配置单例
│   │   ├── load_config.go      # Config 结构体、默认值与 YAML 加载
│   │   ├── logger.go           # 全局 Zap logger 初始化
│   │   └── context.go          # 自定义 Context 与 Handle 包装（含 panic 恢复）
│   ├── controller/             # 控制器（业务处理）
│   ├── exceptions/             # 结构化错误定义
│   ├── middleware/             # 中间件
│   │   ├── trace.go            # Trace ID 与请求级 logger
│   │   └── api.go              # 统一响应/错误处理与耗时日志
│   ├── router/                 # 路由注册
│   ├── utils/                  # 工具函数
│   │   └── logger/             # Zap + rotatelogs 初始化
│   ├── database/               # 预留：数据库连接
│   ├── model/                  # 预留：数据模型
│   ├── repository/             # 预留：数据访问层
│   └── service/                # 预留：业务服务层
├── pkg/                        # 预留：可对外复用的包
├── docs/                       # 预留：文档
├── scripts/                    # 预留：脚本
├── logs/                       # 运行期日志输出（已 gitignore）
├── tmp/                        # Air 构建产物（已 gitignore）
└── .air.toml                   # Air 热重载配置
```

## 环境要求

- Go **1.26.10** 及以上
- 可选：`air`、`dlv`（热重载/调试）、VS Code + Go 扩展

## 快速开始

1. 准备配置文件。仓库不包含 `config/config.yml`（含敏感信息，已 gitignore），请从模板复制：

   ```bash
   cp config/example.config.yml config/config.yml
   ```

2. 按需修改端口、日志目录等配置。

3. 下载依赖并启动：

   ```bash
   go mod download
   go run ./cmd/server
   ```

   > 程序按相对路径 `./config/config.yml` 读取配置，请在项目根目录下运行。

4. 验证服务：

   ```bash
   curl http://127.0.0.1:8080/api/test
   ```

## 开发模式（热重载）

`.air.toml` 已配置为编译到 `./tmp/main`，并通过 Delve 以 headless 模式在 `127.0.0.1:2345` 启动，支持 VS Code 随时 attach。

```bash
# 若尚未安装
go install github.com/air-verse/air@latest
go install github.com/go-delve/delve/cmd/dlv@latest

# air 通常安装在 $(go env GOPATH)/bin，需加入 PATH
export PATH="$(go env GOPATH)/bin:$PATH"

air
```

VS Code 中提供两个启动配置（见 `.vscode/launch.json`）：

- **Start server**：直接调试启动服务，无需 Air。
- **Attach to Air**：先运行 `air`，再启动该配置 attach 到 2345 端口。注意文件变更后 Air 会重启进程，需要重新 attach。

## 配置说明

配置结构定义见 [internal/app/load_config.go](internal/app/load_config.go)。未在文件中出现的字段会使用代码中的默认值。

```yaml
server:
  host: 127.0.0.1   # 监听地址
  port: 8080        # 监听端口
  name: iot         # 服务名，同时用作日志文件名前缀
  logLevel: info    # debug | info | warn | error | dpanic | panic | fatal
  logDir: ./logs    # 日志目录
  logMaxAge: 7      # 日志保留天数
  session:
    sec: xxxxxx     # 预留：session 密钥
    name: APP-SESSION
    lastLang: 604800  # 预留：session 最大时长（秒）

db:                 # 预留：数据库配置，当前未接入
  type: mariadb
  host: 127.0.0.1
  port: 3306
  user: root
  password: ""
  database: app
```

> 请勿将包含真实密码的 `config/config.yml` 提交到版本库，该文件已在 `.gitignore` 中忽略。

## API

所有业务接口挂在 `/api` 分组下，并统一经过 `Trace` 和 `Api` 中间件。示例路由位于 `/api/test`：

| 方法 | 路径 | 说明 |
| --- | --- | --- |
| GET | `/api/test` | 返回 `{"method":"GET"}` |
| GET | `/api/test/error` | 演示错误与 panic 恢复 |
| POST | `/api/test` | 回显请求 JSON |
| PUT | `/api/test` | 返回 `{"method":"PUT"}` |
| PATCH | `/api/test` | 返回 `{"method":"PATCH"}` |
| DELETE | `/api/test` | 返回 `{"method":"DELETE"}` |

### 统一响应格式

成功时由 `Api` 中间件封装为：

```json
{
  "code": 0,
  "success": true,
  "data": { "method": "GET" }
}
```

出错时返回 `AppError` 中的业务码、HTTP 状态与信息：

```json
{
  "code": 1000404,
  "success": false,
  "message": "resource not found"
}
```

内置错误定义见 [internal/exceptions/basic.go](internal/exceptions/basic.go)：

| 错误 | HTTP 状态 | 业务码 |
| --- | --- | --- |
| `ErrNotFound` | 200 | 1000404 |
| `ErrUnauthorized` | 200 | 1000401 |
| `ErrBadRequest` | 200 | 1000400 |
| `ErrInternalServer` | 500 | 1000500 |

业务码约定为 `1000` + 对应 HTTP 语义状态码（如 `404 → 1000404`），便于在响应体中区分类别。

### Trace ID

- 客户端可通过 `X-Trace-ID` 或 `X-Request-ID` 传入；缺失或非法时自动生成 32 位十六进制串。
- 合法字符为字母、数字及 `-`、`_`、`.`、`:`，长度不超过 128。
- Trace ID 会写入响应头 `X-Trace-ID`，并作为 `trace_id` 字段附加到该请求的所有日志中。

## 日志

- 输出：终端 + 文件，均为 JSON 格式。
- 文件路径：`<logDir>/<server.name>-YYYY-MM-DD.log`。
- 切割：每 24 小时按日期切分；超过 `logMaxAge` 天的日志自动删除。
- 级别：由 `server.logLevel` 控制；`error` 及以上级别会自动附带调用栈。
- 每个请求的日志通过独立的派生 logger 输出，互不干扰。

## 如何新增接口

1. 在 `internal/controller/` 新增处理函数，签名为 `func(c *app.Context)`，通过 `c.Body` 设置成功响应数据。
2. 在 `internal/router/` 中注册路由，使用 `app.Handle(controller.YourHandler)` 包装。
3. 返回错误时调用 `c.Error(exceptions.YourError)`；响应体由 `Api` 中间件统一输出。

> 模块名为 `main`，包内引用统一使用 `main/internal/...`。

## 已知限制

- `internal/service`、`internal/repository`、`internal/model`、`internal/database`、`pkg`、`docs`、`scripts` 目前为空目录，仅作为分层占位。
- `db` 与 `session` 配置块尚未接入实际逻辑。
- 错误响应中的 `success` 字段当前按 `code != 0` 计算，恒为 `true`，与成功/失败语义不符（见 [internal/middleware/api.go](internal/middleware/api.go)）。
- 尚未提供单元测试与 Docker 部署配置。
