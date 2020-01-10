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
	// 基本连接声明
	initBasicConn()

	// 交换机声明
	exchangeInit()

	//队列声明
	queueInit()

	// add available commands
	rootCmd.AddCommand(runCmd)

}

func queueInit() {
	rootCmd.PersistentFlags().StringP("queue", "q", "", "producer queue name")
	rootCmd.PersistentFlags().BoolP("queue-durable", "", false, "set true, the queue would be storage on the disk")
	rootCmd.PersistentFlags().BoolP("queue-auto-delete", "", false, `set true, at last one consumer connect to the queue and once all the consumer leaved
			 																										the queue would be delete`)
	rootCmd.PersistentFlags().BoolP("queue-no-wait", "", false, "set true, create queue would be non-blocked without need any ack")
	rootCmd.PersistentFlags().BoolP("queue-exclusive", "", false, "exclusive all  operations from other connect ,but not other channel of the same connect")

	viper.BindPFlag("queue", rootCmd.PersistentFlags().Lookup("queue"))
	viper.BindPFlag("queue-durable", rootCmd.PersistentFlags().Lookup("queue-durable"))
	viper.BindPFlag("queue-auto-delete", rootCmd.PersistentFlags().Lookup("queue-auto-delete"))
	viper.BindPFlag("queue-no-wait", rootCmd.PersistentFlags().Lookup("queue-no-wait"))
	viper.BindPFlag("queue-exclusive", rootCmd.PersistentFlags().Lookup("queue-exclusive"))

	viper.SetDefault("queue-durable", false)
	viper.SetDefault("queue-auto-delete", false)
	viper.SetDefault("queue-no-wait", false)
	viper.SetDefault("queue-exclusive", false)
}

func exchangeInit() {
	rootCmd.PersistentFlags().StringP("exchange", "e", "", "producer exchange name")
	rootCmd.PersistentFlags().StringP("exchange-kind", "k", "direct", "producer exchange type, supposed: `fanout`,`direct`,`topic`,`header`")
	rootCmd.PersistentFlags().BoolP("exchange-duration", "", false, "set true, the exchange would be storage on the disk")
	rootCmd.PersistentFlags().BoolP("exchange-auto-delete", "", false, "set true, all the exchange or queue bind to this exchange would be auto unbind")
	rootCmd.PersistentFlags().BoolP("exchange-internal", "", false, "set true, this exchange would just can be routing from an other routing, not a client")
	rootCmd.PersistentFlags().BoolP("exchange-no-wait", "", false, "set true create exchange  would be non-blocked without need any ack")

	// 交换机相关参数绑定
	viper.BindPFlag("exchange", rootCmd.PersistentFlags().Lookup("exchange"))
	viper.BindPFlag("exchange-kind", rootCmd.PersistentFlags().Lookup("exchange-kind"))
	viper.BindPFlag("exchange-duration", rootCmd.PersistentFlags().Lookup("exchange-duration"))
	viper.BindPFlag("exchange-auto-delete", rootCmd.PersistentFlags().Lookup("exchange-auto-delete"))
	viper.BindPFlag("exchange-internal", rootCmd.PersistentFlags().Lookup("exchange-internal"))
	viper.BindPFlag("exchange-no-wait", rootCmd.PersistentFlags().Lookup("exchange-no-wait"))

	viper.SetDefault("exchange-kind", "direct")
	viper.SetDefault("exchange-duration", false)
	viper.SetDefault("exchange-auto-delete", false)
	viper.SetDefault("exchange-internal", false)
	viper.SetDefault("exchange-no-wait", false)
}

func initBasicConn() {
	rootCmd.PersistentFlags().StringP("bootstrap-server", "b", "localhost:5672", "target connect broker, format as `host:port`")
	rootCmd.PersistentFlags().StringP("username", "u", "guest", "rabbitmq server connect username")
	rootCmd.PersistentFlags().StringP("password", "p", "guest", "rabbitmq server connect password")
	runCmd.PersistentFlags().StringP("protocol", "", "amqp", "rabbitmq server connect protocol, only suppose `amqp` right now")

	// 基本
	viper.BindPFlag("bootstrap-server", rootCmd.PersistentFlags().Lookup("bootstrap-server"))
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("protocol", runCmd.PersistentFlags().Lookup("protocol"))

	//
	viper.SetDefault("bootstrap-server", "localhost:5672")
	viper.SetDefault("username", "guest")
	viper.SetDefault("password", "guest")
	viper.SetDefault("protocol", "amqp")

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
