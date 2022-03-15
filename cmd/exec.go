package cmd

import (
	"bb/adb"
	"bb/config"
	"github.com/spf13/cobra"
)

func init() {
	execCmd.Flags().StringVarP(&command, "command", "c", "pwd", "command to exec")
	uploadFileCmd.MarkFlagRequired("command")

	rootCmd.AddCommand(execCmd)
}

var (
	command string

	execCmd = &cobra.Command{
		Use:   "exec",
		Short: "exec a command on soc",
		Long: `exec a command on soc`,
		Run: func(cmd *cobra.Command, args []string) {
			socIp := config.GetBaseIp()
			socPortList := config.GetSocPortList()
			adb.Shell(socIp, socPortList, command)
		},
	}
)
