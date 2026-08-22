/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/

// cobra-cli add <创建子命令>
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// PreRun 只针对当前命令执行，不会被子命令继承. 先执行prerun后run
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("run 的 PreRun")
	},

	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("执行成功,参数为:verbose")
		} else {
			fmt.Println("run 子命令执行成功")
		}
	},

	// PostRun 只针对当前命令执行，不会被子命令继承. 先执行run后postrun
	/*
	输出执行结果
	关闭当前命令创建的资源
	记录耗时
	*/

	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("run 的 PostRun")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
