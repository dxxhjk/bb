package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	fileCmd = &cobra.Command{
		Use:   "upload_file",
		Short: "upload file",
		Long: `upload file from local to clusters BMC`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("upload_file")
		},
	}
)
