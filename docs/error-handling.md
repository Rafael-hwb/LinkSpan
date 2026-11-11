# 统一错误处理约定

## 接口兼容性

业务成功及业务失败继续使用 HTTP 200，响应保留 code、message、data、requestId、success。
admin 和 project 共用错误中间件、访问日志、绑定错误分类和响应结构。各服务处理的错误响应都附带与 X-Request-ID 对应的请求编号。
路由不存在的 HTTP 404、服务启动失败、连接中断等不属于业务错误包装范围。

## 错误码

| code | 场景 | 示例消息 |
| --- | --- | --- |
| A000002 | JSON 格式、字段类型或请求体大小错误 | 请求体格式错误 / 用户名类型错误 |
| A000003 | 字段校验失败 | 密码不能为空 / 邮箱格式不正确 |
| A000004 | JSON 请求体为空 | 请求体不能为空 |
| A000005 | 未指定任何资料更新字段 | 请至少提供一个需要修改的资料字段 |
| A000110 | 用户名业务格式错误 | 用户名须为3至32位小写字母、数字或下划线 |
| A000120 | 密码超过 bcrypt 字节限制 | 密码不能超过72字节 |
| A000300 | 未取得同用户名注册锁 | 当前系统繁忙，请稍后再试 |
| A000400 | 缺少登录身份或凭证 | 请先登录 |
| A000401 | 会话无效、已过期，或退出凭证无效 | 登录状态无效或已过期 |
| A000402 | 用户不存在、没有密码或密码不匹配 | 用户名或密码错误 |
| A000403 | 试图修改其他用户资料 | 无权修改其他用户的资料 |
| B000001 | 未识别的内部错误、绑定程序错误、panic | 使用既有系统错误消息 |
| B000100 | 未另行分类的 context.DeadlineExceeded | 系统执行超时 |

A000002～A000005、A000400～A000403 是 Go 适配层新增编号，不宣称来自原 Java 枚举。
A000110、A000120、A000300、B000100 的编号沿用原版公共错误码；具体消息按 Go 业务规则定义。
B000200～B000203 等原用户业务码保持不变。A000300 当前用于注册锁竞争，不代表已实现登录请求限流。

## 分层职责

- Handler 使用 ShouldBindJSON / ShouldBindQuery，失败交给 exception.FromBindError。
- FromBindError 保留原始 Cause，分别识别空体、格式错误、类型错误和 validator 字段错误。
- 非指针 JSON 绑定、无效校验对象和未知绑定异常属于程序错误，返回 B000001。
- Service 使用明确业务码；需要保留原因时使用 WrapServiceError。
- Repository 继续返回数据库错误，不负责 HTTP 响应。
- ErrorHandler 记录请求编号、业务错误及底层原因，对外不输出 Cause。
- 普通 error 显式通过 c.Error(err) 上报；panic 由请求中间件兜底。
- 响应已发送时只记录错误，不追加第二份 JSON。

资料修改请求改为绑定到普通结构体，JSON null 会因缺少用户名被拒绝，不再解引用空指针。

## 日志

服务入口改用 RequestLogger，不记录原始 Query、Header、Body 或 Gin 的错误集合输出，避免登录凭证进入默认访问日志。
错误日志与响应通过 requestId 关联；panic 记录类型及堆栈，不打印 panic 值。
底层错误用于本地诊断；新增错误实现中不要把密码、密钥、token 或完整用户资料拼进错误字符串。

## Apifox 更新

现有请求路径和参数位置不变，更新错误示例：

1. POST /api/short-link/admin/v1/user/login，Body 为 {"username":"test3"}，预期 A000003、密码不能为空。
2. 登录 Body 为空，预期 A000004；JSON 截断或字段类型错误，预期 A000002。
3. GET /api/short-link/admin/v1/user/check-login 缺少 token，预期 A000003、登录凭证不能为空。
4. PUT /api/short-link/admin/v1/user 不带登录 Header，预期 A000400。
5. 带无效会话修改资料，预期 A000401；有效身份修改其他用户，预期 A000403。
6. DELETE /api/short-link/admin/v1/user/logout 使用无效凭证，预期 A000401。

错误响应示例（message 类型保持现有 string）：

```json
{
  "code": "A000003",
  "message": "密码不能为空",
  "data": null,
  "requestId": "本次响应头X-Request-ID的值",
  "success": false
}
```

## 验证

执行 `go test ./admin/...`。测试不依赖真实 MySQL 或 Redis，不修改业务数据。
覆盖参数解析、字段校验、错误原因链、身份/权限分支、panic 恢复、超时、响应防重写与访问日志凭证隐藏。

## project 短链接接口

创建、分页、修改通过 `c.Error(err)` 把错误交给 project 的 `ErrorHandler`。空请求体使用 `A000004`，JSON 格式或类型错误使用 `A000002`，字段校验和业务参数错误使用 `A000003`，执行超时使用 `B000100`，其他内部错误和 panic 使用 `B000001`。数据库错误由 Repository 返回，客户端收到约定的错误消息。创建、分页、修改的响应体 `requestId` 与响应头 `X-Request-ID` 对应。admin 转发接口会保留 project 响应体里的请求编号，admin 自己的响应头用于查 admin 日志。

### Apifox 验证

在 Apifox 切换到 project 的本地环境 `http://127.0.0.1:8002`，发送以下请求。每项预期 HTTP 200、`success: false`、`data: null`，并可在响应 Header 对照 `X-Request-ID`。

| 请求 | 输入 | 预期 |
| --- | --- | --- |
| POST /api/short-link/v1/create | Body 选 JSON，只填写一个左花括号 | A000002，请求体格式错误 |
| GET /api/short-link/v1/page | Params 留空 | A000003，分组标识不能为空 |
| POST /api/short-link/v1/update | 从可用的修改请求复制 Body，将 originUrl 改成 go.dev/tour/ | A000003，originUrl 须为完整的 HTTP 或 HTTPS 地址 |

在 admin 的修改请求重复最后一个用例，使用当前账号拥有的分组，并在 Headers 里填写有效的 username 和 token。预期收到相同的业务错误码和消息。执行 `go test ./...` 可验证错误响应、超时和 panic，无需连接数据库。
