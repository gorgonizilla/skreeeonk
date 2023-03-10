package main

import (
	"fmt"

	"github.com/gorgonizilla/skreeeonk/internal/banner"
)

func main() {
	output, err := banner.Render("skreeeonk", "plain")
	if err != nil {
		panic(err)
	}

	fmt.Println(output)
}
