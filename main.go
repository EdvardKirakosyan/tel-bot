package main

import (
	"flag"
	"log"
)

func main() {
	mustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New(tgClient)

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)

}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to tg-bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
