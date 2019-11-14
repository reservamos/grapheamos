package main

import (
	"fmt"

	"github.com/reservamos/grapheamos/internal/graphql"
)

func main() {
	fmt.Println("~~~~ Powered by Reservamos ~~~~")
	graphql.Start()
}
