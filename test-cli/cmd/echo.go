/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
// 参数验证

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	// 参数验证（最小需要1个参数）
	// cobra.NoArgs: 不允许任何参数。
	// cobra.ExactArgs(n): 必须有 n 个参数。
	// cobra.MinimumNArgs(n): 至少要有 n 个参数。
	// cobra.MaximumNArgs(n): 最多只能有 n 个参数。
	// cobra.RangeArgs(min, max): 参数个数必须在 min 和 max 之间
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("echo called")
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// echoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// echoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
