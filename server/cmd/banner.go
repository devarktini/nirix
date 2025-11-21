package main

import (
	"fmt"
	"os"

	"github.com/dimiro1/banner"
)

func printBanner() {
	templ := `{{ .Title "NIRIX" "" 8 }}
   	   This is the task NIRIX server which handles all operations.`

	banner.InitString(os.Stdout, true, true, templ)
	fmt.Println()
}
