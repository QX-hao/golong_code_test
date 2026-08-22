# Cobra 入门

Cobra 是一个用于创建 cli 命令行程序的库。它可以帮助我们组织根命令、子命令、flag、参数校验以及命令执行前后的生命周期逻辑。

相关代码：[test-cli 项目](https://github.com/QX-hao/golong_code_test/tree/main/test-cli)

## 1. 使用 Cobra 创建 CLI 程序

### 1.1 安装 cobra-cli

cobra-cli 是 Cobra 提供的代码生成工具，可以自动生成项目和命令文件。

~~~
go install github.com/spf13/cobra-cli@latest
~~~

如果终端找不到 cobra-cli，可以使用完整路径：

~~~
"$(go env GOPATH)/bin/cobra-cli" --help
~~~

如果 Go 的可执行文件目录已经加入 PATH，则可以直接使用：

~~~
cobra-cli --help
~~~

需要区分：

- github.com/spf13/cobra 是项目运行时使用的 Go 库。
- cobra-cli 是生成项目骨架和子命令文件的工具。

### 1.2 初始化 Go Module

进入准备创建 CLI 的目录：

~~~
mkdir my-cli
cd my-cli
go mod init github.com/你的用户名/my-cli
~~~

如果目录中已经存在 go.mod，不要重复执行 go mod init。

### 1.3 初始化 Cobra

在 Go Module 根目录执行：

~~~
cobra-cli init --author "你的名字" --license apache
~~~

这条命令会生成基本项目结构：

~~~
my-cli/
├── main.go
├── go.mod
├── go.sum
├── LICENSE
└── cmd/
    └── root.go
~~~

参数说明：

- --author：生成代码中的作者或版权归属信息。
- --license apache：使用 Apache License 2.0 许可证模板。
- 如果暂时不需要作者和许可证信息，可以直接执行 cobra-cli init。

初始化后整理依赖：

~~~
go mod tidy
~~~

### 1.4 main.go 和 cmd.Execute()

Cobra 项目的 main.go 通常很简单：

~~~
package main

import "github.com/你的用户名/my-cli/cmd"

func main() {
    cmd.Execute()
}
~~~

cmd.Execute() 通常定义在 cmd/root.go 中：

~~~
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
~~~

它会启动整个 Cobra 命令树：

~~~
main()
  -> cmd.Execute()
  -> rootCmd.Execute()
  -> 解析命令、flag 和参数
  -> 执行对应命令
~~~

### 1.5 root.go 根命令

根命令是整个 CLI 的入口，通常定义在 cmd/root.go：

~~~
var rootCmd = &cobra.Command{
    Use:   "my-cli",
    Short: "my-cli 的简短描述",
    Long:  "my-cli 的详细描述",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("my-cli started")
    },
}
~~~

cobra.Command 中几个常用字段：

- Use：命令名称。
- Short：命令的简短描述。
- Long：命令的详细描述。
- Run：命令的核心执行逻辑。
- RunE：可以返回错误的执行逻辑。

当前项目的根命令使用了：

~~~
var rootCmd = &cobra.Command{
    Use:   "test-cli",
    Short: "test-cli is a test cli tool ,this is short description",
    Long:  "test-cli is a test cli tool, this is long description",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("测试test-cli")
    },
}
~~~

运行根命令：

~~~
go run .
~~~

查看帮助：

~~~
go run . --help
go run . -h
~~~

--help 和 -h 是 Cobra 自动添加的帮助 flag。它们会显示命令的描述、用法、子命令和 flag。

## 2. 创建子命令

### 2.1 使用 cobra-cli add

在项目根目录执行：

~~~
cobra-cli add run
cobra-cli add exec
~~~

执行后会生成：

~~~
cmd/run.go
cmd/exec.go
~~~

项目命令结构变成：

~~~
test-cli
├── run
└── exec
~~~

### 2.2 注册子命令

生成的命令文件会在 init() 中注册到根命令：

~~~
func init() {
    rootCmd.AddCommand(runCmd)
}
~~~

exec 命令也是相同的方式：

~~~
func init() {
    rootCmd.AddCommand(execCmd)
}
~~~

AddCommand 的作用是建立父子命令关系。没有调用 AddCommand 时，即使定义了一个 cobra.Command，用户也不能通过根命令找到它。

### 2.3 定义子命令

一个基本的子命令如下：

~~~
var runCmd = &cobra.Command{
    Use:   "run",
    Short: "执行 run 命令",
    Long:  "run 命令的详细描述",

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("run 子命令执行成功")
    },
}
~~~

运行子命令：

~~~
go run . run
~~~

查看子命令帮助：

~~~
go run . run --help
go run . run -h
~~~

完整的命令调用形式是：

~~~
根命令 子命令 flag 参数
~~~

例如：

~~~
test-cli run
test-cli exec --info abc
~~~

### 2.4 子命令的执行流程

执行：

~~~
go run . run
~~~

Cobra 会：

1. 启动 main.go。
2. 调用 cmd.Execute()。
3. 找到 runCmd。
4. 解析 run 的 flag 和参数。
5. 执行 runCmd 的执行逻辑。

## 3. 创建 flag

flag 用于修改命令的执行行为，例如：

~~~
test-cli exec --info abc
test-cli run --verbose
~~~

flag 通常分为：

- 字符串 flag：需要接收一个字符串值。
- 布尔 flag：表示打开或关闭一个选项。
- 数字 flag：接收整数或浮点数。
- 本地 flag：只属于某一个命令。
- 持久化 flag：当前命令和所有子命令都可以使用。

### 3.1 子命令 flag：Flags()

当前项目在 exec.go 中定义了 info：

~~~
var info string

func init() {
    rootCmd.AddCommand(execCmd)

    execCmd.Flags().StringVarP(
        &info,
        "info",
        "i",
        "",
        "info 的相关信息",
    )
}
~~~

StringVarP 的参数依次是：

~~~
StringVarP(
    p *string,
    name string,
    shorthand string,
    value string,
    usage string,
)
~~~

对应关系：

- &info：将用户输入的值保存到 info。
- "info"：完整名称，使用 --info。
- "i"：简写名称，使用 -i。
- ""：默认值为空字符串。
- 最后的字符串：帮助信息中的说明。

因此下面两种写法等价：

~~~
go run . exec --info abc
go run . exec -i abc
~~~

在 Run 中读取变量：

~~~
Run: func(cmd *cobra.Command, args []string) {
    if info == "" {
        fmt.Println("info 为空")
        cmd.Help()
        return
    }

    fmt.Println("info:", info)
},
~~~

执行：

~~~
go run . exec -i abc
~~~

输出：

~~~
info: abc
~~~

### 3.2 字符串 flag 必须有值

因为 info 是字符串 flag，下面的命令会报错：

~~~
go run . exec -i
~~~

错误发生在 Cobra 解析 flag 的阶段，Run 函数还没有执行：

~~~
Error: flag needs an argument: 'i' in -i
~~~

正确用法：

~~~
go run . exec -i abc
go run . exec --info abc
~~~

如果想显式传入空字符串：

~~~
go run . exec -i ""
~~~

此时才会进入：

~~~
if info == "" {
    // ...
}
~~~

### 3.3 布尔 flag

如果一个 flag 只是表示“是否开启”，不需要额外的值，应使用布尔 flag：

~~~
var force bool

func init() {
    execCmd.Flags().BoolVarP(
        &force,
        "force",
        "f",
        false,
        "是否强制执行",
    )
}
~~~

使用：

~~~
go run . exec --force
go run . exec -f
~~~

布尔 flag 的特点：

- 不需要写额外的值。
- 不传入时通常是 false。
- 传入后变为 true。

### 3.4 全局 flag：PersistentFlags()

当前项目在 root.go 中定义了全局 verbose：

~~~
var verbose bool

func init() {
    rootCmd.PersistentFlags().BoolVarP(
        &verbose,
        "verbose",
        "v",
        false,
        "是否显示详细信息",
    )
}
~~~

PersistentFlags() 表示这个 flag 不只属于根命令，也可以被所有子命令使用：

~~~
go run . -v
go run . run -v
go run . exec -v
~~~

与之相对的是：

~~~
execCmd.Flags()
~~~

它只会给 exec 命令增加 flag：

~~~
go run . exec -i abc
~~~

run 不能使用 --info：

~~~
go run . run --info abc
~~~

两者的区别：

~~~
Flags()
    只对当前命令生效

PersistentFlags()
    对当前命令和所有子命令生效
~~~

需要注意：PersistentFlags() 只负责让子命令能够接收这个 flag，不会自动改变子命令的业务逻辑。

当前项目的 run.go 中主动读取了 verbose：

~~~
Run: func(cmd *cobra.Command, args []string) {
    if verbose {
        fmt.Println("执行成功,参数为:verbose")
    } else {
        fmt.Println("run 子命令执行成功")
    }
},
~~~

因此：

~~~
go run . run
~~~

输出普通结果，而：

~~~
go run . run -v
~~~

会输出详细模式的结果。

## 4. 参数校验

flag 和参数不是一回事。

flag 带有 - 或 -- 前缀：

~~~
test-cli exec --info abc
~~~

这里的 --info 是 flag，abc 是 flag 的值。

普通参数直接写在命令后面：

~~~
test-cli run file.txt
~~~

这里的 file.txt 是普通参数，不是 flag。

Cobra 可以通过 Args 字段校验普通参数的数量和内容。

### 4.1 不允许参数：NoArgs

如果命令不接受普通参数：

~~~
var runCmd = &cobra.Command{
    Use:  "run",
    Args: cobra.NoArgs,

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("run")
    },
}
~~~

下面的命令可以执行：

~~~
go run . run
~~~

下面的命令会校验失败：

~~~
go run . run abc
~~~

### 4.2 必须有一个参数：ExactArgs

如果命令必须接收一个参数：

~~~
var runCmd = &cobra.Command{
    Use:  "run [name]",
    Args: cobra.ExactArgs(1),

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("name:", args[0])
    },
}
~~~

正确：

~~~
go run . run alice
~~~

错误：

~~~
go run . run
go run . run alice bob
~~~

### 4.3 至少一个参数：MinimumNArgs

如果命令至少需要一个参数：

~~~
var execCmd = &cobra.Command{
    Use:  "exec [command]",
    Args: cobra.MinimumNArgs(1),

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("要执行的内容:", args)
    },
}
~~~

正确：

~~~
go run . exec ls
go run . exec -- ls -l
~~~

错误：

~~~
go run . exec
~~~

常用内置校验器：

~~~
cobra.NoArgs
cobra.ExactArgs(n)
cobra.MinimumNArgs(n)
cobra.MaximumNArgs(n)
cobra.RangeArgs(min, max)
~~~

### 4.4 自定义参数校验

如果内置校验器不能满足需求，可以自己写函数：

~~~
Args: func(cmd *cobra.Command, args []string) error {
    if len(args) != 1 {
        return fmt.Errorf("必须提供一个参数")
    }

    if args[0] == "" {
        return fmt.Errorf("参数不能为空")
    }

    return nil
},
~~~

参数校验失败时，Cobra 不会执行 Run。

这和字符串 flag 缺少值的情况类似：

~~~
参数解析或校验失败
    -> 输出错误
    -> 不执行 Run
~~~

## 5. Hooks 生命周期钩子

Hooks 用于在命令执行前后插入公共逻辑。Cobra 常用的 Hook 有：

~~~
PersistentPreRun
PreRun
Run
PostRun
PersistentPostRun
~~~

执行顺序：

~~~
PersistentPreRun
    ↓
PreRun
    ↓
Run
    ↓
PostRun
    ↓
PersistentPostRun
~~~

### 5.1 PersistentPreRun

PersistentPreRun 在 Run 前执行，并且可以被子命令继承。

当前项目的 root.go 中：

~~~
PersistentPreRun: func(cmd *cobra.Command, args []string) {
    if verbose {
        fmt.Println("verbose is true,显示详细信息")
    }
},
~~~

执行：

~~~
go run . run -v
~~~

会先执行根命令的 PersistentPreRun，然后执行 run 命令的逻辑。

适合放：

- 读取配置文件。
- 初始化日志。
- 初始化数据库连接。
- 检查全局 flag。
- 统一处理 --verbose。

### 5.2 PreRun

PreRun 只针对当前命令执行，不会被子命令继承。

当前项目的 run.go 中：

~~~
PreRun: func(cmd *cobra.Command, args []string) {
    fmt.Println("run 的 PreRun")
},
~~~

### 5.3 Run

Run 是当前命令的核心业务逻辑。

当前项目的 run.go 中：

~~~
Run: func(cmd *cobra.Command, args []string) {
    if verbose {
        fmt.Println("执行成功,参数为:verbose")
    } else {
        fmt.Println("run 子命令执行成功")
    }
},
~~~

### 5.4 PostRun

PostRun 在当前命令的 Run 执行完成后执行，不会被子命令继承。

当前项目的 run.go 中：

~~~
PostRun: func(cmd *cobra.Command, args []string) {
    fmt.Println("run 的 PostRun")
},
~~~

适合放当前命令的收尾逻辑，例如：

- 输出执行结果。
- 记录耗时。
- 释放当前命令创建的资源。

### 5.5 PersistentPostRun

PersistentPostRun 在 Run 和 PostRun 之后执行，并且可以被子命令继承。

当前项目的 root.go 中：

~~~
PersistentPostRun: func(cmd *cobra.Command, args []string) {
    fmt.Println("全局清理逻辑")
},
~~~

因为 run 没有自己的 PersistentPostRun，所以执行：

~~~
go run . run
~~~

默认会执行 root 的 PersistentPostRun。

### 5.6 当前项目的完整执行顺序

当前项目执行：

~~~
go run . run -v
~~~

大致会输出：

~~~
verbose is true,显示详细信息
run 的 PreRun
执行成功,参数为:verbose
run 的 PostRun
全局清理逻辑
~~~

对应关系：

~~~
root PersistentPreRun
    ↓
run PreRun
    ↓
run Run
    ↓
run PostRun
    ↓
root PersistentPostRun
~~~

### 5.7 root 和子命令都有 Persistent Hook 时听谁的

默认情况下，子命令优先。

例如 root 和 run 都定义了 PersistentPreRun：

~~~
root PersistentPreRun
run PersistentPreRun
~~~

执行：

~~~
go run . run
~~~

默认只会执行：

~~~
run PersistentPreRun
~~~

原因是子命令自己的 Hook 会覆盖父命令同名 Hook。所谓“继承”可以理解为：

~~~
子命令没有自己的 PersistentPreRun
    -> 使用父命令的

子命令有自己的 PersistentPreRun
    -> 使用子命令的
~~~

PersistentPostRun 也是同样的规则：默认优先使用距离当前命令最近的 Hook。

如果希望父命令和子命令的持久化 Hook 都执行，可以设置：

~~~
cobra.EnableTraverseRunHooks = true
~~~

开启后，执行子命令时的顺序是：

~~~
父命令 PersistentPreRun
    ↓
子命令 PersistentPreRun
    ↓
子命令 Run
    ↓
子命令 PersistentPostRun
    ↓
父命令 PersistentPostRun
~~~

### 5.8 使用 RunE 返回错误

如果命令执行过程可能失败，可以使用 RunE：

~~~
RunE: func(cmd *cobra.Command, args []string) error {
    if info == "" {
        return cmd.Help()
    }

    fmt.Println("info:", info)
    return nil
},
~~~

Hook 也有对应的错误版本：

~~~
PersistentPreRunE
PreRunE
RunE
PostRunE
PersistentPostRunE
~~~

例如：

~~~
PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
    if verbose {
        fmt.Println("详细模式已开启")
    }

    return nil
},
~~~

如果返回非 nil 错误，后续的命令逻辑通常不会继续执行。

## 常用验证命令

格式化 Go 代码：

~~~
gofmt -w ./cmd/root.go
gofmt -w ./cmd/run.go
gofmt -w ./cmd/exec.go
~~~

运行帮助：

~~~
go run . --help
go run . run --help
go run . exec --help
~~~

运行命令：

~~~
go run .
go run . run
go run . run -v
go run . exec -i abc
~~~

运行测试和编译：

~~~
go test ./...
go build -o test-cli .
./test-cli --help
~~~
