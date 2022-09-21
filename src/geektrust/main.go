package main

import (
	"geektrust/clients"
	"geektrust/handlers"
	"os"
)

func main() {
	var writer clients.BaseWriter
	var reader clients.BaseReader

	reader = clients.NewCartFileReader()
	writer = clients.NewShellWriter(os.Stdout, clients.DefaultOptions)

	handlers.CartHandler(writer, reader)
}
