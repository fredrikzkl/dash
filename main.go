package main

import (
	"fmt"

	"github.com/fredrikzkl/dash/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
