package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("wait 1 minute...")
	time.Sleep(5 * time.Second)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})
	fmt.Println("⇨ test-wait1 started on [::]:8080")
	http.ListenAndServe(":8080", nil)

}
