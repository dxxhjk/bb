package cmd

import (
	"bb/config"
	"bb/handler"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	execCmd.Flags().StringVarP(&command, "command", "c", "", "command to exec")
	execCmd.MarkFlagRequired("command")
	execCmd.Flags().StringVarP(&startSoc, "start_soc", "s",
		config.GetSocIpListInternal()[0], "It is used to specify the port number or IP of the starting soc. If internal mode is enabled, specify the IP")
	execCmd.Flags().StringVarP(&socNum, "soc_num", "n", strconv.Itoa(len(config.GetSocPortList())), "The name of the file to be distributed in the \"file\" folder")
	execCmd.Flags().BoolVarP(&energyMonitor, "energy_monitor", "e", false, "Whether to monitor the command energy consumption")
	execCmd.Flags().StringVarP(&energyMonitorOutput, "energy_monitor_output_file", "o", "", "Energy monitor output file name, effective only when -e flag is selected")

	rootCmd.AddCommand(execCmd)
}

var (
	command string
	energyMonitor bool
	energyMonitorOutput string

	execCmd = &cobra.Command{
		Use:   "exec",
		Short: "exec a command on designated soc",
		Long: `exec a command on designated soc`,
		Run: func(cmd *cobra.Command, args []string) {
			if internal {
				handler.ExecInternal(startSoc, socNum, command, energyMonitor, energyMonitorOutput)
			} else {
				handler.Exec(startSoc, socNum, command, energyMonitor, energyMonitorOutput)
			}
		},
	}
)
