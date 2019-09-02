package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(2)
	}
	var iwd int
	iwd, _ = strconv.Atoi(os.Args[1])
	if int(time.Now().Weekday()) != iwd {
		os.Exit(1)
	}
}
