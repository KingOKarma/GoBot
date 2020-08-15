package main

// Variables used for command line parameters
import (
	"fmt"
	"github.com/Floor-Gang/utilpkg"
	"github.com/Floor-Gang/GoBot/internal"
)

func main() {
	fmt.Println("Inside main main.go")
	config := internal.GetConfig(internal.ConfigPath)
	internal.Start(config, internal.ConfigPath)
	fmt.Println("")

	utilpkg.KeepAlive()
}
