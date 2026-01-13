package main

import (
	"flag"
	"fmt"

	"gitlab.com/alexandergmiles/publisher/internal/announcer"
)

// Pass in JSON file containing the types of message
// and mapping to message type (version flag)
// Maybe a path to some auth details

func main() {
	fmt.Println("Starting publisher")

	messageConfig := flag.String("config", "", "Path to .JSON configuration file which maps message versions")
	flag.Parse()
	print(messageConfig)

	sqsAnnouncer := announcer.SQSAnnouncer{}

	fmt.Println(sqsAnnouncer)
}

func parseOptions() {
}
