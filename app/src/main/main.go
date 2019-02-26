package main

import (
    "fmt"
    "router"
)

func main() {
    fmt.Println("HOTBOX")
    e := router.New()

    e.Start(":8000")
}
