# file_operations_demo02

这个 demo 演示文件写入，包含 3 种方式：直接写、`bufio` 缓冲写、一次性写文件。

## 涵盖的操作

- 打开或创建文件
- 写入日志内容
- 关闭文件

## 方法对比

| 方法 | 代码位置 | 核心 API | 写入机制 | 行为特点 | 适用场景 |
| --- | --- | --- | --- | --- | --- |
| 方法一 | `FileWrite1` | `os.OpenFile` + `Write/WriteString` | 直接写入文件 | 代码直观，立即写；可同时写字节和字符串 | 少量写入、简单场景 |
| 方法二 | `FileWrite2` | `os.OpenFile` + `bufio.NewWriter` + `Flush` | 先写缓冲区再刷盘 | 小块高频写时通常更高效；必须调用 `Flush()` | 日志、批量写文本、频繁写入 |
| 方法三 | `FileWrite3` | `ioutil.WriteFile`（已废弃） | 一次性覆盖写入 | 简洁，但会覆盖原文件内容 | 一次性落盘小文件（推荐改 `os.WriteFile`） |

## 关键区别说明

- 方法一/二通过 `os.OpenFile` 控制打开模式，示例中使用 `os.O_CREATE|os.O_RDWR|os.O_APPEND`，因此是“创建 + 读写 + 追加”。
- 方法二如果忘记 `Flush()`，缓存内容可能不会真正写入文件。
- 方法三（`WriteFile`）语义是“把给定内容写成文件最终内容”，通常配合覆盖写，不适合做追加日志。

## 运行方式

在当前目录执行：

```bash
go run .
```

默认输出文件：

- `log/YYYY-MM-DD.log`（方法一、方法二追加写）
- `log/YYYY-MM-DD.log1`（方法三写入）

## 代码细节提示

- 当前示例 `FileWrite3` 使用 `ioutil.WriteFile`，建议迁移为 `os.WriteFile`。
- 如果希望方法三也“追加写”，应改成 `os.OpenFile(..., os.O_APPEND|os.O_CREATE|os.O_WRONLY, perm)` 再 `WriteString`，而不是 `WriteFile`。
