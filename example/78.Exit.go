package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!") // Exitで抜ける時はdeferは実行されない

	os.Exit(3)
}
