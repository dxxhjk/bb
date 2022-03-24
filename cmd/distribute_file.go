package cmd

import (
	"bb/config"
	"bb/handler"
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
	distributeFileCmd.Flags().StringVarP(&startSoc, "start_soc", "s",
		config.GetSocIpListInternal()[0], "It is used to specify the port number or IP of the starting soc. If internal mode is enabled, specify the IP")
	distributeFileCmd.Flags().StringVarP(&socNum, "soc_num", "n", strconv.Itoa(len(config.GetSocPortList())), "The name of the file to be distributed in the \"file\" folder")

	rootCmd.AddCommand(distributeFileCmd)
}

var (
	fileToDistribute string
	destination string
	startSoc string
	socNum string

	distributeFileCmd = &cobra.Command{
		Use:   "distribute_file",
		Short: "distribute file to soc",
		Long: `batch distribute designated file to designated soc`,
		Run: func(cmd *cobra.Command, args []string) {
			if internal {
				handler.DistributeFileInternal(startSoc, socNum,fileToDistribute, destination)
			} else {
				handler.DistributeFile(startSoc, socNum,fileToDistribute, destination)
			}
		},
	}
)

