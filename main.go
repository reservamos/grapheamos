package main

import (
	"fmt"
	"os"

	"github.com/reservamos/graphql-start/internal/graphql"
)

func main() {
	fmt.Println("~~~~ Powered by Reservamos ~~~~")
	fmt.Printf("Server running at: http://localhost:%s\n", os.Getenv("PORT"))
	graphql.Start()
}
