package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"os"
)

type Conn struct {
}

type Exchange struct {
	name       string
	kind       bool
	duration   bool
	autoDelete bool
	internal   bool
	noWait     bool
}

type Queue struct {
	name       string
	duration   bool
	autoDelete bool
	noWait     bool
	exclusive  bool
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s :%s", msg, err)
	}
}
func exit() {
	fmt.Println("exit...")
}

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

		connUrl := fmt.Sprintf("%s://%s:%s@%s/", protocol, username, password, bootstrapServer)
		conn, err := amqp.Dial(connUrl)
		failOnError(err, "Fail to connect to RabbitMq")
		defer exit()
		defer conn.Close()

		channel, err := conn.Channel()
		failOnError(err, "Fail to open a channel")
		defer channel.Close()
		err = channel.ExchangeDeclare(
			viper.GetString("exchange"),           // exchange n
			viper.GetString("exchange-kind"),      // exchange n
			viper.GetBool("exchange-duration"),    // exchange n
			viper.GetBool("exchange-auto-delete"), // exchange n
			viper.GetBool("exchange-internal"),    // exchange n
			viper.GetBool("exchange-no-wait"),
			nil,
		)
		failOnError(err, "Fail to declare a exchange")

		queue, err := channel.QueueDeclare(
			viper.GetString("exchange"),        // exchange n
			viper.GetBool("queue-durable"),     // exchange n
			viper.GetBool("queue-auto-delete"), // exchange n
			viper.GetBool("queue-exclusive"),   // exchange n
			viper.GetBool("queue-no-wait"),
			nil,
		)
		failOnError(err, "Fail to declare a queue")
		fmt.Println(fmt.Sprintf("connect to broker\nhost:%s\nuser:%s\nexchange:%squeue:%s\nenter `exit` to quit",
			bootstrapServer, username, viper.GetString("exchange"), queue.Name))
		ch := make(chan int)
		go func() {
			scanner := bufio.NewScanner(os.Stdin)
			for {
				fmt.Print(">")
				scanner.Scan()
				text := scanner.Text()
				if text == "exit" {
					ch <- 1
					break
				} else {
					fmt.Println(text)
				}

			}

		}()
		<-ch
		return
	},
}
