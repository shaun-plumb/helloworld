package main

import (
	"encoding/json"
	"fmt"
	"helloworld/internal/version"
	"runtime/debug"
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
