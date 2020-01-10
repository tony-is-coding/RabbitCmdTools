package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//var versionCmd = &cobra.Command{
//	Use:   "version",
//	Short: "Print the version number of Hugo",
//	Long:  `All software has versions. This is Hugo's`,
//	Run: func(cmd *cobra.Command, args []string) {
//		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
//	},
//}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "start up a rabbitmq producer client",
	Long:  "hello this is just a test",
	Run: func(cmd *cobra.Command, args []string) {

		// get value
		bootstrapServer := viper.Get("bootstrap-server")
		username := viper.Get("username")
		password := viper.Get("password")
		protocol := viper.Get("protocol")

		// handle value
		fmt.Println("server connect to :", bootstrapServer)
		fmt.Println("connect use is  :", username)
		fmt.Println("password is  :", password)
		fmt.Println("protocol is :", protocol)
	},
}
