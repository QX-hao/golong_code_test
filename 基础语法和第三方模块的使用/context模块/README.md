# Go context 模块学习

这里不是 6 个项目，而是一个轻量学习包：一个 Go module，6 个递进案例。每个案例都能单独运行，适合从简单到复杂理解 `context`。

先进入目录：

```bash
cd /Users/sealos/data/golong_code_test/基础语法和第三方模块的使用/context模块
```

## 学习顺序

### 1. `Background` 和 `TODO`

目标：知道 `context.Context` 通常作为第一个参数向下传递。

```bash
go run ./examples/01_background_todo
```

重点：

- `context.Background()` 一般作为根 context。
- `context.TODO()` 表示这里以后应该补上合适的 context。
- 普通函数约定写成 `func Do(ctx context.Context, ...)`。

### 2. `WithCancel`

目标：理解手动取消，以及 goroutine 如何监听 `ctx.Done()` 退出。

```bash
go run ./examples/02_with_cancel
```

重点：

- `cancel()` 会关闭 `ctx.Done()`。
- goroutine 不会自动停止，必须自己监听 `ctx.Done()`。

### 3. `WithTimeout`

目标：理解超时控制。

```bash
go run ./examples/03_with_timeout
```

重点：

- 超过指定时间后，`ctx.Err()` 通常是 `context deadline exceeded`。
- 有 `cancel` 就 `defer cancel()`，即使超时会自动触发，也要释放资源。

### 4. `WithDeadline`

目标：理解截止时间可以在多层调用中共享。

```bash
go run ./examples/04_with_deadline
```

重点：

- `WithTimeout` 是“从现在开始多久”。
- `WithDeadline` 是“到某个具体时间点为止”。

### 5. `WithValue`

目标：理解请求级元数据如何沿调用链传递。

```bash
go run ./examples/05_with_value
```

重点：

- 适合传 `request_id`、`trace_id`、认证信息等。
- 不适合传普通业务参数、数据库连接、大对象。
- key 不建议直接用字符串，应该自定义类型，避免冲突。

### 6. HTTP 请求调用链

目标：把 context 放到真实服务调用链里看。

```bash
go run ./examples/06_http_request_chain
```

这个例子模拟：

```text
HTTP Handler -> Service -> Repository
```

同一个 `ctx` 会一路传下去。Repository 如果太慢，Handler 设置的超时会让它提前结束。

## 建议练习方式

1. 先按顺序运行 6 个例子。
2. 每个例子只改一个时间参数，比如把 `500*time.Millisecond` 改成 `50*time.Millisecond`。
3. 观察输出里的 `context canceled` 和 `context deadline exceeded`。
4. 在第 2、3、4、6 个例子里，试着删除 `case <-ctx.Done()`，看程序行为有什么变化。

## 常见规则

- `ctx` 通常放在函数第一个参数。
- 不要传 `nil` context，不确定时用 `context.TODO()`。
- 有 `cancel` 就 `defer cancel()`。
- 不要把 context 长期存在 struct 字段里。
- 不要用 context 传普通业务参数。
- goroutine 必须主动监听 `ctx.Done()` 才能响应取消。
