package main

import (
	"geektrust/clients"
	"geektrust/handlers"
	"os"
)

func main() {
	var client clients.BaseWriter
	client = clients.NewShellWriter(os.Stdout, clients.DefaultOptions)

	// Controller to handle the cart commands
	handlers.CartHandler(client)
}
