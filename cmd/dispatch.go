/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"github.com/sirupsen/logrus"
	"loadtester/client"

	"github.com/spf13/cobra"
)

var requests int
var time int
var url string
var method string
var body string


// dispatchCmd represents the dispatch command
var dispatchCmd = &cobra.Command{
	Use:   "dispatch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Main invocation to be done here
		logrus.WithFields(logrus.Fields{
			"Request Num": requests,
			"Time": time,
			"URL": url,
			"Method Type": method,
			"Body Path": body,
		}).Info("Printing Request information")
		switch method {
		case "GET":
			client.GetRequest(url)
		case "POST":{
				if body == "" {
					logrus.Fatal("Body not provided")
				}
				client.PostRequest(url, body)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dispatchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Set Flags for this command:
	dispatchCmd.PersistentFlags().IntVar(&requests, "request", 1, "Number of requests to be fired")
	dispatchCmd.PersistentFlags().IntVar(&time, "time", 1, "Duration of the test")
	dispatchCmd.PersistentFlags().StringVar(&url, "url", "http://localhost:8080/get", "URL to stress test")
	dispatchCmd.PersistentFlags().StringVar(&method, "X", "GET", "Request Type")
	dispatchCmd.PersistentFlags().StringVar(&body, "body", "", "Path to the JSON Body")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dispatchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}