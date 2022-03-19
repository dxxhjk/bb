package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var(
	rootCmd = &cobra.Command{
		Use:   "bb",
		Short: "batch_bench",
		Long:  "batch_bench",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("使用 bb -h 或者 bb help 查看使用帮助")
		},
	}
)

func init() {}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
