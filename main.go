package main

import (
	"github.com/snowboardit/reserved/cmd/app"
	"github.com/snowboardit/reserved/pkg/reserved"
)

var r = reserved.Reserved{}

func main() {
	// load data and words
	if err := r.Load(); err != nil {
		panic(err)
	}

	app.Start()
}
