package cmd

import (
	"bb/config"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	uploadFileCmd.Flags().StringVarP(&filePath, "file_path", "f", "your_file/dic_path", "path of your local file to upload")
	uploadFileCmd.MarkFlagRequired("file_path")
	uploadFileCmd.Flags().StringVarP(&loginName, "login_name", "n", "zl", "your login name")

	rootCmd.AddCommand(uploadFileCmd)
}

var (
	filePath string
	loginName string

	uploadFileCmd = &cobra.Command{
		Use:   "upload_file",
		Short: "get upload file command",
		Long: `show the command to upload file from local to clusters BMC`,
		Run: func(cmd *cobra.Command, args []string) {
			localPath := config.GetWorkPath() + "file/"
			localIp := config.GetBaseIp()
			localPort := config.GetLocalPort()
			fmt.Println("use command:")
			fmt.Println("scp -r -P " + localPort + " " +
				filePath + " " + loginName + "@" +
				localIp + ":" + localPath)
		},
	}
)
