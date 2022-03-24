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
	collectResultCmd.Flags().StringVarP(&fileToCollect, "file_path", "f", "your_file/dic_name", "The name of the file to be distributed in the \"file\" folder")
	collectResultCmd.MarkFlagRequired("file_path")
	collectResultCmd.Flags().StringVarP(&localPathToSaveFiles, "destination", "d", config.GetWorkPath() + "/result", "Path to save the file")
	collectResultCmd.Flags().StringVarP(&startSoc, "start_soc", "s",
		config.GetSocIpListInternal()[0], "It is used to specify the port number or IP of the starting soc. If internal mode is enabled, specify the IP")
	collectResultCmd.Flags().StringVarP(&socNum, "soc_num", "n", strconv.Itoa(len(config.GetSocPortList())), "The name of the file to be distributed in the \"file\" folder")

	rootCmd.AddCommand(collectResultCmd)
}

var (
	fileToCollect string
	localPathToSaveFiles string

	collectResultCmd = &cobra.Command{
		Use:   "collect_result",
		Short: "collect result from soc",
		Long: `batch collect designated result from designated soc`,
		Run: func(cmd *cobra.Command, args []string) {
			if internal {
				handler.CollectResultInternal(fileToCollect, localPathToSaveFiles, startSoc, socNum)
			} else {
				handler.CollectResult(fileToCollect, localPathToSaveFiles, startSoc, socNum)
			}
		},
	}
)

