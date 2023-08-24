package main

import (
	"fmt"

	devto "github.com/cosmasnyairo/devto/cmd/api/devto"
)

func main() {
	if err := devto.Run(); err != nil {
		fmt.Println(err)
	}
}
