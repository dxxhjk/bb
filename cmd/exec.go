package cmd

import (
	"bb/adb"
	"bb/config"
	"bb/util"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	execCmd.Flags().StringVarP(&command, "command", "c", "", "command to exec")
	execCmd.MarkFlagRequired("command")
	execCmd.Flags().StringVarP(&startSocPort, "start_soc_port", "s",
		config.GetSocPortList()[0], "The name of the file to be distributed in the \"file\" folder")
	execCmd.Flags().StringVarP(&socNum, "soc_num", "n", strconv.Itoa(len(config.GetSocPortList())), "The name of the file to be distributed in the \"file\" folder")
	execCmd.Flags().BoolVarP(&energyMonitor, "energy_monitor", "e", false, "Whether to monitor the command energy consumption")

	rootCmd.AddCommand(execCmd)
}

var (
	command string
	energyMonitor bool

	execCmd = &cobra.Command{
		Use:   "exec",
		Short: "exec a command on designated soc",
		Long: `exec a command on designated soc`,
		Run: func(cmd *cobra.Command, args []string) {
			socIp := config.GetBaseIp()
			socPortList := config.GetSocPortList()
			socPortList, err := util.GetDesignatedPortList(startSocPort, socNum, socPortList)
			if err != nil {
				fmt.Println(err)
				return
			}
			adb.Init(socIp, socPortList)
			adb.Shell(socIp, socPortList, command, energyMonitor)
		},
	}
)
