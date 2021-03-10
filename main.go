package main

import (
	"fmt"
	"provider_probe/src"
)
const timerPort string = "8090"

func main() {
	email := "example@example.com"
	fmt.Print(src.SendRequest(email))
}