/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"github.com/spf13/cobra"
	"time"
)

// replayCmd represents the replay command
var replayCmd = &cobra.Command{
	Args: cobra.MinimumNArgs(1),
	Use:   "replay [file to replay]",
	Short: "replay a set of log files",
	Long: `replay a set of log files at a speed of your choice`,
	Run: func(cmd *cobra.Command, args []string) {
		times, _ := cmd.Flags().GetInt("speed")
		target, _ := cmd.Flags().GetString("target.folder")
		replayLog(args[0],times, target)
	},
}

func init() {
	rootCmd.AddCommand(replayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	replayCmd.PersistentFlags().Int("speed", 1, "set the speed at which the log will be replayed")
	replayCmd.PersistentFlags().String("target.folder", "", "target folder to replay to")
	rootCmd.MarkPersistentFlagRequired("target.folder")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func replayLog(file string, speed int, target string){

	timeRun := time.Time{}

	fmt.Println(timeRun.Unix())
	fmt.Println(file,speed, target)
}