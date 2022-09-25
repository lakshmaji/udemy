package cmd

import (
	reader_client "geektrust/clients/reader"
	"io"
)

type mockInput struct {
	ErrParseFileName    error
	ErrParseFileContent error
	ErrParseFileLines   error
	FileName            string
	FileContent         io.Reader
	FileLines           []string
}

type client struct {
	mockInput
}

func newMock(
	input mockInput,
) reader_client.BaseReader {
	return &client{input}
}

func (f *client) ParseFileName() (string, error) {
	if f.mockInput.ErrParseFileName != nil {
		return "", f.mockInput.ErrParseFileName
	}
	return f.mockInput.FileName, nil
}

func (f *client) ParseFileContent(name string) (io.Reader, error) {
	if f.mockInput.ErrParseFileContent != nil {
		return nil, f.mockInput.ErrParseFileContent
	}
	return f.mockInput.FileContent, nil
}

func (f *client) ParseFileLines(file io.Reader) ([]string, error) {
	if f.mockInput.ErrParseFileLines != nil {
		return nil, f.mockInput.ErrParseFileLines
	}
	return f.mockInput.FileLines, nil
}
