package main

import (
	"fmt"
	"project/internal/app"
)

func main() {
	err := app.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
