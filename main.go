package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args[1:]) >= 1 &&  os.Args[1] != "" {
		fmt.Println(len(strings.Split(os.Args[1], " ")))
	}else {
		fmt.Println(0)
	}
}