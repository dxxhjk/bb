package cmd

import (
	"bb/adb"
	"bb/config"
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func init() {
	if err := config.InitConfig(); err != nil {
		fmt.Println("Init config error: " + err.Error())
	}
	distributeFileCmd.Flags().StringVarP(&fileToDistribute, "file_path", "f", "your_file/dic_name", "The name of the file to be distributed in the \"file\" folder")
	distributeFileCmd.MarkFlagRequired("file_path")
	distributeFileCmd.Flags().StringVarP(&destination, "destination", "d", "/data/dd_workspace", "Path to save the file")
	distributeFileCmd.Flags().StringVarP(&startSocPort, "start_soc_port", "s",
		config.GetSocPortList()[0], "The name of the file to be distributed in the \"file\" folder")
	distributeFileCmd.Flags().StringVarP(&socNum, "soc_num", "n", "60", "The name of the file to be distributed in the \"file\" folder")

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
			pwdCmd := exec.Command("bash", "-c", "pwd")
			var stdout bytes.Buffer
			pwdCmd.Stdout = &stdout
			if err := pwdCmd.Run(); err != nil {
				fmt.Println("get pwd error: " + err.Error())
			}
			localPath := fileToDistribute
			if fileToDistribute[0] != '/' {
				pwd := stdout.String()
				pwd = pwd[:len(pwd) - 1] + "/"
				localPath = pwd + localPath
			}
			socIp := config.GetBaseIp()
			socPortList := config.GetSocPortList()
			adb.Init(socIp, socPortList)
			adb.Push(socIp, socPortList, localPath, destination)
		},
	}
)

