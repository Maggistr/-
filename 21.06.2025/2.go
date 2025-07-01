package main

import (
	"fmt"
	"time"
)

func hello_time() {
	fmt.Println("Hello, my name is Timur")
	fmt.Println("Current date:", time.Now().Format("2006-01-02"))
}
