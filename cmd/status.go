/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/romberli/go-template/config"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "status command",
	Long:  `print server status.`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			err       error
			isRunning bool
		)

		// init config
		err = initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf("%s\n%s", config.Messages[config.ErrInitConfig].Error(), err.Error()))
		}

		// check if given pid is running
		if serverPid != constant.DefaultRandomInt {
			isRunning, err = linux.IsRunningWithPid(serverPid)
			if err != nil {
				fmt.Println(fmt.Sprintf("%s\n%s", config.Messages[config.ErrCheckServerRunningStatus].Error(), err.Error()))
				return
			}
			if isRunning {
				fmt.Println(config.Messages[config.InfoServerIsRunning].Renew(serverPid).Error())
			} else {
				fmt.Println(config.Messages[config.InfoServerNotRunning].Renew(serverPid).Error())
			}

			return
		}

		// get pid
		serverPidFile = viper.GetString(config.ServerPidFileKey)
		serverPid, err = linux.GetPidFromPidFile(serverPidFile)
		if err != nil {
			fmt.Println(fmt.Sprintf("%s\n%s", config.Messages[config.ErrGetPidFromPidFile].Renew(serverPidFile).Error(), err.Error()))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
		isRunning, err = linux.IsRunningWithPid(serverPid)
		if err != nil {
			fmt.Println(fmt.Sprintf("%s\n%s", config.Messages[config.ErrCheckServerRunningStatus].Error(), err.Error()))
			return
		}
		if isRunning {
			fmt.Println(config.Messages[config.InfoServerIsRunning].Renew(serverPid).Error())
		} else {
			fmt.Println(config.Messages[config.InfoServerNotRunning].Renew(serverPid).Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")
	statusCmd.PersistentFlags().IntVar(&serverPid, "server-pid", constant.DefaultRandomInt, fmt.Sprintf("specify the server pid"))

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
