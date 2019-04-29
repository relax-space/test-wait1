package main

import (
	"fmt"
	"time"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	fmt.Println("wait 1 minute...")
	time.Sleep(1 * time.Minute)
	fmt.Println("test-wait1 started")
	if err := e.Start(":8080"); err != nil {
		fmt.Println(err)
	}
}
