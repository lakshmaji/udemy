/*
Command-line application to purchase different kinds of programmes from Geekdemy.
Generate a total bill of the programmes after applying the discounts, if any.

Contains loosely coupled services by applying SOLID principles.

Follows Object Oriented Programming Style.
*/
package main

import (
	reader_client "geektrust/clients/reader"
	writer_client "geektrust/clients/writer"
	"geektrust/handlers"
	"os"
)

// Generate the bill of purchases from Geekdemy.
func main() {
	var writer writer_client.BaseWriter = writer_client.New(os.Stdout, writer_client.DefaultOptions)
	var reader reader_client.BaseReader = reader_client.New(os.DirFS("/"))

	handlers.CartHandler(writer, reader)
}
