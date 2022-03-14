package cmd

import (
	"bb/config"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var(
	author string
	rootCmd = &cobra.Command{
		Use:   "bb",
		Short: "batch_bench",
		Long:  "batch_bench",
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			fmt.Println(config.GetSocIpList())
			fmt.Println(config.GetWorPath())
			fmt.Println(author)
			fmt.Printf("running batch_bench command.\n")
		},
	}
)

func init() {
	if err := config.InitConfig(); err != nil {
		fmt.Println("Init config error: " + err.Error())
	}
	rootCmd.AddCommand(fileCmd)
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "Pad", "author name for copyright attribution")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
