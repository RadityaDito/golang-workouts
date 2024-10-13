package main

import (
	"fmt"
	"time"
)

func _() {
	n := 5
	duration := time.Duration(n) * time.Second
	fmt.Println(duration.Seconds())
}
