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

type mockClient struct {
	mockInput
}

func mockNewReader(
	input mockInput,
) reader_client.BaseReader {
	return &mockClient{input}
}

func (f *mockClient) ParseFileName() (string, error) {
	if f.mockInput.ErrParseFileName != nil {
		return "", f.mockInput.ErrParseFileName
	}
	return f.mockInput.FileName, nil
}

func (f *mockClient) ParseFileContent(name string) (io.Reader, error) {
	if f.mockInput.ErrParseFileContent != nil {
		return nil, f.mockInput.ErrParseFileContent
	}
	return f.mockInput.FileContent, nil
}

func (f *mockClient) ParseFileLines(file io.Reader) ([]string, error) {
	if f.mockInput.ErrParseFileLines != nil {
		return nil, f.mockInput.ErrParseFileLines
	}
	return f.mockInput.FileLines, nil
}
