package handlers

import (
	"bytes"
	"testing"
	"testing/fstest"

	"geektrust/clients/reader"
	reader_client "geektrust/clients/reader"
	writer_client "geektrust/clients/writer"
)

func TestCartHandler(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	const (
		input string = `ADD_PROGRAMME CERTIFICATION 1
ADD_PROGRAMME DEGREE 1
ADD_PROGRAMME DIPLOMA 2
APPLY_COUPON DEAL_G20
PRINT_BILL`
		output string = `SUB_TOTAL	13000.00
COUPON_DISCOUNT	B4G1	2500.00
TOTAL_PRO_DISCOUNT	0.00
PRO_MEMBERSHIP_FEE	0.00
ENROLLMENT_FEE	0.00
TOTAL	10500.00
`
	)
	var response bytes.Buffer
	fs := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}
	writer := writer_client.New(&response, writer_client.DefaultOptions)
	reader := reader_client.New(fs)
	CartHandler(writer, reader)

	result := response.String()

	if result == "" {
		t.Error("Should return bill")
	}

	if number := bytes.Compare([]byte(output), response.Bytes()); number != 0 {
		t.Errorf("expected %s, received %s", output, result)
	}
}
