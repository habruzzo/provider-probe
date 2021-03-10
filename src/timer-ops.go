package src

import (
	"fmt"
	"net/http"
)

const TimerPort string = "8090"

func RegisterTimer(name string) {
	_, err := http.Get(makeUrl(name, "register"))
	if err != nil {
		fmt.Printf("problem with register timer %s", name)
	}
}

func StartTimer(name string) {
	_, err := http.Get(makeUrl(name, "start"))
	if err != nil {
		fmt.Printf("problem with starting timer %s", name)
	}
}

func StopTimer(name string) {
	_, err := http.Get(makeUrl(name, "start"))
	if err != nil {
		fmt.Printf("problem with stopping timer %s", name)
	}
}

func makeUrl(name string, operation string) string {
	return fmt.Sprintf("localhost:%s/%s?name=%s", TimerPort, operation, name)
}