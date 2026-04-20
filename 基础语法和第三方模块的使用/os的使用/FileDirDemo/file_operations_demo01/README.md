# file_operations_demo01

这个 demo 演示文件读取，包含 3 种常见方式：底层流读取、`bufio` 按行读取、一次性整文件读取。

## 涵盖的操作

- 打开文件（只读）
- 读取文件内容
- 关闭文件

## 方法对比

| 方法 | 代码位置 | 核心 API | 读取粒度 | 内存特征 | 适用场景 |
| --- | --- | --- | --- | --- | --- |
| 方法一 | `FileRead1` | `os.Open` + `file.Read` | 固定字节块（128 字节） | 渐进读取，内存可控 | 需要自己掌控读取过程（协议解析、二进制处理） |
| 方法二 | `FileRead2` | `bufio.NewReader` + `ReadString('\n')` | 按行 | 相对可控，但会累加到 `FileStr` | 文本日志、配置文件等按行处理 |
| 方法三-1 | `FileRead31` | `ioutil.ReadFile`（已废弃） | 整文件 | 一次性读入内存 | 小文件快速读取（历史写法） |
| 方法三-2 | `FileRead32` | `os.ReadFile` | 整文件 | 一次性读入内存 | 小文件快速读取（推荐替代 `ioutil.ReadFile`） |

## 关键区别说明

- `os.File.Read`（方法一）最底层，最灵活，但代码量较大，需要自己处理循环、`EOF`、拼接缓冲区。
- `bufio.Reader`（方法二）更偏文本场景，按行读取更直观。
- `os.ReadFile`（方法三）最简洁，但对大文件不友好，因为会整体加载到内存。
- `ioutil.ReadFile` 已在 Go 1.16 后不推荐使用，功能由 `os.ReadFile` 替代。

## 运行方式

在当前目录执行：

```bash
go run .
```

默认读取文件：

- `config/config.json`

## 代码细节提示

- 方法一里 `buf = append(buf, tempBuff[:n]...)` 这一写法是正确的；如果直接 append 整个 `tempBuff`，最后一块会带入无效填充字节。
- 示例中打开文件和关闭文件分开写，便于教学；实际业务中更推荐在打开成功后立刻 `defer file.Close()`，减少忘记关闭的风险。
