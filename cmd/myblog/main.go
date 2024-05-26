package main

import (
	"fmt"
	"os"

	"github.com/ivancheng/myblog/internal/myblog"
)

func main() {
	command := myblog.NewMiniBlogCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
	fmt.Println("Hello MiniBlog 12")
}
