package main

import (
	"comments-api/cmd/comments-api/parameters"
	"fmt"
	"os"
)

func main() {

	err := parameters.RootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
