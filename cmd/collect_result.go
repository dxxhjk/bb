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
	collectResultCmd.Flags().StringVarP(&fileToCollect, "file_path", "f", "your_file/dic_name", "The name of the file to be distributed in the \"file\" folder")
	collectResultCmd.MarkFlagRequired("file_path")
	collectResultCmd.Flags().StringVarP(&localPathToSaveFiles, "destination", "d", config.GetWorPath() + "/result", "Path to save the file")
	collectResultCmd.Flags().StringVarP(&startSocPort, "start_soc_port", "s",
		config.GetSocPortList()[0], "The name of the file to be distributed in the \"file\" folder")
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
			socIp := config.GetBaseIp()
			socPortList := config.GetSocPortList()
			socPortList, err := util.GetDesignatedPortList(startSocPort, socNum, socPortList)
			if err != nil {
				fmt.Println(err)
				return
			}
			adb.Init(socIp, socPortList)
			adb.Pull(socIp, socPortList, fileToCollect, localPathToSaveFiles)
		},
	}
)

