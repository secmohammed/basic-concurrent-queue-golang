package main

import (
	"fmt"
	"time"
)

func PrintPayload(payload map[string]string) {
	time.Sleep(time.Second * 2)
	fmt.Println(payload)
}
