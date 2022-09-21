package main

import (
	reader_client "geektrust/clients/reader"
	writer_client "geektrust/clients/writer"
	"geektrust/handlers"
	"os"
)

func main() {
	var writer writer_client.BaseWriter
	var reader reader_client.BaseReader

	reader = reader_client.New()
	writer = writer_client.New(os.Stdout, writer_client.DefaultOptions)

	handlers.CartHandler(writer, reader)
}
