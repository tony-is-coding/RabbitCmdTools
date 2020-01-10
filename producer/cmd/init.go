package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.

	// create a root command
	rootCmd = &cobra.Command{
		Use:   "rabbitmq-producer",
		Short: "a rabbitmq producer client tool",
		Long:  `rabbitmq-producer is a command line for start up rabbitmq-producer client`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// show flags
	rootCmd.PersistentFlags().StringP("bootstrap-server", "b", "localhost:5672", "your target connect broker, format as `host:port`")
	rootCmd.PersistentFlags().StringP("username", "u", "guest", "your rabbitmq server connect username")
	rootCmd.PersistentFlags().StringP("password", "p", "guest", "your rabbitmq server connect password")
	runCmd.PersistentFlags().StringP("protocol", "", "amqp", "your rabbitmq server connect protocol")

	//    ---------  bind flags
	viper.BindPFlag("bootstrap-server", rootCmd.PersistentFlags().Lookup("bootstrap-server"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("protocol", runCmd.PersistentFlags().Lookup("protocol"))

	// ------- set flags default
	viper.SetDefault("bootstrap-server", "localhost:5672")
	viper.SetDefault("username", "guest")
	viper.SetDefault("password", "guest")
	viper.SetDefault("protocol", "amqp")

	// add available commands
	rootCmd.AddCommand(runCmd)

}

//func initConfig() {
//	if cfgFile != "" {
//		// Use config file from the flag.
//		viper.SetConfigFile(cfgFile)
//	} else {
//		// Find home directory.
//		home, err := homedir.Dir()
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		// Search config in home directory with name ".cobra" (without extension).
//		viper.AddConfigPath(home)
//		viper.SetConfigName(".cobra")
//	}
//
//	viper.AutomaticEnv()
//
//	if err := viper.ReadInConfig(); err == nil {
//		fmt.Println("Using config file:", viper.ConfigFileUsed())
//	}
//}
