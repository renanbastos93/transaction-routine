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

// func main() {
// 	calc := 0
// 	a := -1
// 	b := -3

// 	calc = calc - a
// 	calc = calc - b
// 	// (-1) - (-3)
// 	fmt.Println(calc, -1-3)
// }
