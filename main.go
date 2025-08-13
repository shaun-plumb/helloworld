package main

import (
	"encoding/json"
	"fmt"

	"runtime/debug"

	"github.com/shaun-plumb/helloworld/internal/version"
)

func main() {

	bi, ok := debug.ReadBuildInfo()
	if ok {
		j, _ := json.MarshalIndent(bi, "", "  ")
		fmt.Println(string(j))
	}

	fmt.Println("Hello...")

	fmt.Println(version.GetVersion())

}
