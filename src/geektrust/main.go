package main

import (
	"geektrust/clients"
	"geektrust/handlers"
	"os"
)

func main() {
	// Bind: Where to write the output
	var client clients.BaseWriter
	client = clients.NewShellWriter(os.Stdout, clients.DefaultOptions)

	// TODO: rename `HandleCart` to `?`
	// Controller to handle the cart commands
	handlers.HandleCart(client)
}
