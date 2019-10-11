package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dchote/envirophat-mqtt/config"
	"github.com/dchote/envirophat-mqtt/envirophat"

	"github.com/docopt/docopt-go"
)

var exitChan = make(chan int)

// VERSION beause...
const VERSION = "0.0.1"

func cliArguments() {
	usage := `
Usage: envirophat-mqtt [options]

Options:
  --server=<server>          MQTT server host/IP [default: 127.0.0.1]
  --port=<port>              MQTT server port [default: 1883]
  --topic=<topic>            MQTT topic prefix [default: envirophat]
  --clientid=<clientid>      MQTT client identifier [default: envirohat]
  --interval=<interval>      Poll interval (seconds) [default: 5]
  -h, --help                 Show this screen.
  -v, --version              Show version.
`
	args, _ := docopt.ParseArgs(usage, os.Args[1:], VERSION)

	config.Server, _ = args.String("--server")
	config.Port, _ = args.Int("--port")
	config.TopicPrefix, _ = args.String("--topic")
	config.ClientID, _ = args.String("--clientid")
	config.Interval, _ = args.Int("--interval")

	log.Printf("server: %s, port: %d, topic prefix: %s, client identifier: %s, poll interval: %d", config.Server, config.Port, config.TopicPrefix, config.ClientID, config.Interval)
}

// sigChannelListen basic handlers for inbound signals
func sigChannelListen() {
	signalChan := make(chan os.Signal, 1)
	code := 0

	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)
	signal.Notify(signalChan, syscall.SIGTERM)

	select {
	case sig := <-signalChan:
		log.Printf("Received signal %s. shutting down", sig)
	case code = <-exitChan:
		switch code {
		case 0:
			log.Println("Shutting down")
		default:
			log.Println("*Shutting down")
		}
	}

	os.Exit(code)
}

func main() {
	cliArguments()

	// catch signals
	go sigChannelListen()

	// run sensor poll loop
	envirophat.Init()

	os.Exit(0)
}
