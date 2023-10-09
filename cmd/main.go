package main

import (
	"context"
	"log"

	"github.com/ServiceWeaver/weaver"
	"github.com/renanbastos93/transaction-routine/internal/bff"
)

func main() {
	if err := weaver.Run(context.Background(), bff.Server); err != nil {
		log.Fatal(err)
	}
}
