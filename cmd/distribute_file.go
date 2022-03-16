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
	if err := config.InitConfig(); err != nil {
		fmt.Println("Init config error: " + err.Error())
	}
	distributeFileCmd.Flags().StringVarP(&fileToDistribute, "file_path", "f", "your_file/dic_name", "The name of the file to be distributed in the \"file\" folder")
	distributeFileCmd.MarkFlagRequired("file_path")
	distributeFileCmd.Flags().StringVarP(&destination, "destination", "d", "/data/bb_workspace", "Path to save the file")
	distributeFileCmd.Flags().StringVarP(&startSocPort, "start_soc_port", "s",
		config.GetSocPortList()[0], "The name of the file to be distributed in the \"file\" folder")
	distributeFileCmd.Flags().StringVarP(&socNum, "soc_num", "n", strconv.Itoa(len(config.GetSocPortList())), "The name of the file to be distributed in the \"file\" folder")

	rootCmd.AddCommand(distributeFileCmd)
}

var (
	fileToDistribute string
	destination string
	startSocPort string
	socNum string

	distributeFileCmd = &cobra.Command{
		Use:   "distribute_file",
		Short: "distribute file to soc",
		Long: `batch distribute designated file to designated soc`,
		Run: func(cmd *cobra.Command, args []string) {
			socIp := config.GetBaseIp()
			socPortList := config.GetSocPortList()
			socPortList, err := util.GetDesignatedPortList(startSocPort, socNum, socPortList)
			if err != nil {
				fmt.Println(err)
				return
			}
			adb.Init(socIp, socPortList)
			adb.Push(socIp, socPortList, fileToDistribute, destination)
		},
	}
)

