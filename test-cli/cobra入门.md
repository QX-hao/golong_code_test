# Cobra 入门

## 1. Cobra 是什么

Cobra 是一个用于构建 Go 命令行程序的库，主要帮助我们处理：

- 根命令和子命令，例如 my-cli exec。
- 命令参数和 flag，例如 my-cli exec --info abc。
- 自动生成帮助信息，例如 my-cli exec --help。
- 命令参数校验和错误处理。
- 子命令的层级组织。

Cobra 相关工具分为两部分：

- github.com/spf13/cobra：项目运行时使用的 Go 库。
- cobra-cli：用于生成 Cobra 项目和子命令文件的脚手架工具。

## 2. 安装 Cobra CLI 工具

安装 Cobra 的代码生成工具：

~~~
go install github.com/spf13/cobra-cli@latest
~~~

这条命令安装的是 cobra-cli 生成器，不是把 Cobra 业务代码写入项目。

如果系统提示找不到 cobra-cli，可以使用完整路径执行：

~~~
"$(go env GOPATH)/bin/cobra-cli" --help
~~~

如果 Go 的可执行文件目录已经加入 PATH，则可以直接使用：

~~~
cobra-cli --help
~~~

## 3. 初始化 Cobra 项目

新建目录并初始化 Go Module：

~~~
mkdir my-cli
cd my-cli
go mod init github.com/你的用户名/my-cli
~~~

如果项目已经存在 go.mod，不要重复执行 go mod init。

然后初始化 Cobra：

~~~
go install github.com/spf13/cobra-cli@latest
cobra-cli init --author "你的名字" --license apache
~~~

其中：

- --author 是源码中的作者或版权归属信息，不是账号认证信息。
- --license apache 表示使用 Apache License 2.0 许可证模板。
- 如果暂时不需要作者和许可证参数，可以直接执行 cobra-cli init。

初始化完成后整理依赖：

~~~
go mod tidy
~~~

Module 路径会影响 main.go 中导入 cmd 包的路径，也会影响将来其他项目引用该模块的方式。

## 4. main.go 和 cmd.Execute()

main.go 通常只负责启动根命令：

~~~
package main

import "github.com/你的用户名/my-cli/cmd"

func main() {
    cmd.Execute()
}
~~~

这里的 cmd 是 cmd 包的名称。cmd.Execute() 会启动根命令：

~~~
main()
  -> cmd.Execute()
  -> rootCmd.Execute()
  -> Cobra 解析命令和 flag
  -> 执行对应命令的 Run 或 RunE
~~~

cmd/root.go 中的 Execute 通常是：

~~~
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}
~~~

因此，main.go 不需要直接处理具体子命令的业务逻辑。

## 5. 根命令 root.go

根命令通常定义在 cmd/root.go：

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

cobra.Command 中几个常用字段的含义：

- Use：命令名称和使用形式。
- Short：短描述，通常显示在父命令的子命令列表中。
- Long：详细描述，通常显示在当前命令的帮助信息中。
- Run：命令被执行时运行的函数。
- RunE：可以返回错误的执行函数，适合需要显式处理错误的场景。

运行根命令：

~~~
go run .
~~~

程序会执行根命令的 Run 函数并输出：

~~~
my-cli started
~~~

查看根命令帮助：

~~~
go run . --help
go run . -h
~~~

--help 和 -h 是 Cobra 自动提供的帮助 flag。它们会显示命令描述、使用方式、可用 flag 和子命令，不会执行当前命令的 Run 函数。

## 6. 使用 cobra-cli add 添加子命令

在项目根目录执行：

~~~
cobra-cli add run
cobra-cli add exec
~~~

每条命令都会在 cmd 目录生成对应的 Go 文件：

~~~
cmd/run.go
cmd/exec.go
~~~

生成的子命令会通过 init 函数注册到根命令。例如 cmd/run.go 中：

~~~
func init() {
    rootCmd.AddCommand(runCmd)
}
~~~

exec.go 中也是同样的注册方式：

~~~
func init() {
    rootCmd.AddCommand(execCmd)
}
~~~

添加子命令后，命令关系可以理解为：

~~~
my-cli
├── run
└── exec
~~~

运行子命令：

~~~
go run . run
go run . exec
~~~

查看子命令帮助：

~~~
go run . run --help
go run . exec --help
go run . exec -h
~~~

子命令自己的 Short、Long 和 flag 会显示在对应的帮助页面中。

## 7. 给 exec 添加 flag

在 exec.go 中定义一个字符串 flag：

~~~
var info string

func init() {
    rootCmd.AddCommand(execCmd)

    execCmd.Flags().StringVarP(
        &info,
        "info",
        "i",
        "",
        "info的相关信息,这里可以使用 cli exec --help/-h查看",
    )
}
~~~

StringVarP 的函数形式如下：

~~~
StringVarP(
    p *string,
    name string,
    shorthand string,
    value string,
    usage string,
)
~~~

对应示例代码中的参数：

- &info：将用户输入的值保存到 info 变量。
- "info"：完整 flag 名称，即 --info。
- "i"：简写 flag 名称，即 -i。
- ""：默认值为空字符串。
- 最后的字符串：在 --help 中显示的说明文字。

因此下面两种写法等价：

~~~
go run . exec --info abc
go run . exec -i abc
~~~

在 Run 函数中读取变量即可使用 flag 的值：

~~~
Run: func(cmd *cobra.Command, args []string) {
    if info == "" {
        fmt.Println("info为空,显示 exec 帮助信息")
        _ = cmd.Help()
        return
    }

    fmt.Println("执行exec --info/-i逻辑")
    fmt.Println("info", info)
},
~~~

执行：

~~~
go run . exec -i abc
~~~

输出类似：

~~~
执行exec --info/-i逻辑
info abc
~~~

如果不传入 info：

~~~
go run . exec
~~~

由于 info 的默认值是空字符串，会进入 if info == "" 分支，并显示 exec 命令的帮助信息。

### -i 为什么必须有值

info 是通过 StringVarP 定义的字符串 flag，因此下面的命令会报错：

~~~
go run . exec -i
~~~

原因是 -i 后面缺少字符串值。这个错误发生在 Cobra 解析参数阶段，早于 Run 函数执行，所以 Run 中的判断不会被执行。

正确示例：

~~~
go run . exec -i abc
go run . exec --info abc
go run . exec -i ""
~~~

如果 flag 不需要值，而是表示一个开关，应使用布尔 flag：

~~~
var force bool

execCmd.Flags().BoolVarP(
    &force,
    "force",
    "f",
    false,
    "是否强制执行",
)
~~~

使用方式：

~~~
go run . exec --force
go run . exec -f
~~~

## 8. Flags 和 PersistentFlags

### 本地 flag：Flags()

当前的 info 使用的是：

~~~
execCmd.Flags().StringVarP(...)
~~~

这表示 info 只属于 exec 命令：

~~~
go run . exec --info abc
~~~

其他命令不能使用它：

~~~
go run . run --info abc
~~~

### 持久化 flag：PersistentFlags()

如果一个 flag 需要被根命令和所有子命令使用，可以定义在根命令上：

~~~
var verbose bool

func init() {
    rootCmd.PersistentFlags().BoolVarP(
        &verbose,
        "verbose",
        "v",
        false,
        "输出详细日志",
    )
}
~~~

这样根命令、run 和 exec 都可以使用：

~~~
go run . --verbose
go run . run --verbose
go run . exec --verbose
~~~

简单记忆：

~~~
Flags()             只对当前命令生效
PersistentFlags()   对当前命令和所有子命令生效
~~~

## 9. Run 和 RunE

Run 适合不需要返回错误的简单逻辑：

~~~
Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("执行命令")
},
~~~

RunE 可以返回错误：

~~~
RunE: func(cmd *cobra.Command, args []string) error {
    if info == "" {
        return cmd.Help()
    }

    fmt.Println("info:", info)
    return nil
},
~~~

cmd.Help() 会返回一个 error。使用 Run 时可以暂时忽略：

~~~
_ = cmd.Help()
~~~

但更推荐在 RunE 中直接返回：

~~~
return cmd.Help()
~~~

_ 表示明确丢弃函数返回值。正常显示帮助时，_ = cmd.Help() 和 cmd.Help() 的输出没有区别；区别只在于是否处理 cmd.Help() 返回的错误。

## 10. 编译和运行

开发阶段可以使用：

~~~
go run .
go run . run
go run . exec -i abc
~~~

格式化代码：

~~~
gofmt -w ./cmd/root.go
gofmt -w ./cmd/run.go
gofmt -w ./cmd/exec.go
~~~

测试整个项目：

~~~
go test ./...
~~~

编译为可执行文件：

~~~
mkdir -p bin
go build -o bin/my-cli .
~~~

运行编译后的程序：

~~~
./bin/my-cli --help
./bin/my-cli run
./bin/my-cli exec -i abc
~~~

go run . 会临时编译并运行，go build 会生成可执行文件，但不会自动把命令安装到系统 PATH 中。

## 11. 推荐开发流程

以后为项目增加功能时，可以按照下面的顺序：

~~~
1. 使用 cobra-cli add 命令名 创建子命令
2. 在 cmd/命令名.go 中修改 Use、Short、Long
3. 使用 Flags() 或 PersistentFlags() 添加 flag
4. 在 Run 或 RunE 中读取 flag 并实现业务逻辑
5. 使用 --help 检查命令说明
6. 使用 gofmt 格式化代码
7. 使用 go test ./... 和 go build 验证项目
~~~

例如新增 config 命令：

~~~
cobra-cli add config
go run . config --help
~~~

## 参考资料

- [李文周：Go CLI 开发利器：Cobra 简明教程](https://liwenzhou.com/posts/go/cobra/)
- [Cobra 官方仓库](https://github.com/spf13/cobra)
- [Cobra CLI 官方说明](https://github.com/spf13/cobra-cli)
