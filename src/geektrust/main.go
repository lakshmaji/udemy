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
	"log"
	"os"
	"path/filepath"
)

// Generate the bill of purchases from Geekdemy.
func main() {
	dir := filepath.Dir(os.Args[1:][0])
	if !filepath.IsAbs(dir) {
		log.Fatal("Not absolute path")
	}
	var writer writer_client.BaseWriter = writer_client.New(os.Stdout, writer_client.DefaultOptions)
	var reader reader_client.BaseReader = reader_client.New(os.DirFS(dir))
	handlers.CartHandler(writer, reader)
}
