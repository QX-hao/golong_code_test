/*
Copyright © 2026 QX-hao

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// cli 默认run
// 全局flags
// hooks

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-cli",
	Short: "test-cli is a test cli tool ,this is short description",
	Long: `test-cli is a test cli tool ,
	this is long description`,
	// Uncomment the following line if your bare application
	// has an action associated with it:

	/*
	hooks
	执行顺序
		root PersistentPreRun
		run PreRun
		run Run
		run PostRun
	r	oot PersistentPostRun
	*/

	/*
		PersistentPreRun
		在命令的 Run 之前执行，并且可以被子命令继承。
		适合放公共初始化逻辑，例如：
		读取配置文件
		初始化日志
		初始化数据库连接
		检查全局 flag
	*/

	// 会先执行 PersistentPreRun，再执行 runCmd.Run
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("verbose is true,显示详细信息")
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("测试test-cli")
	},

	// PersistentPostRun 在 Run 和 PostRun 之后执行，并且可以被子命令继承
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("全局清理逻辑")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.test-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Persistent Flag(全局flag)
	rootCmd.PersistentFlags().BoolVarP(
		&verbose,
		"verbose",
		"v",
		false,
		"是否显示详细信息",
	)
}
