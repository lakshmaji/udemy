package handlers

import (
	"bytes"
	"strconv"
	"testing"
	"testing/fstest"

	"geektrust/clients/reader"
	reader_client "geektrust/clients/reader"
	writer_client "geektrust/clients/writer"
	"geektrust/utils"
)

func TestErrorNoInputFileNameProvided(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go"}
	reader.OsArgs = mockArgs

	var response bytes.Buffer
	defer func() {
		r := recover()
		expected := utils.ErrorNoFilePath.Error()
		received := response.String()
		if received != expected {
			t.Errorf("Expected %s, received %s", expected, received)
		}
		if r == nil {
			t.Error("Should write error")
		}
	}()

	mockFS := fstest.MapFS{}
	writer := writer_client.New(&response, writer_client.DefaultTestOptions)
	reader := reader_client.New(mockFS)
	CartHandler(writer, reader)
}

func TestErrorFileNotFound(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	var response bytes.Buffer
	defer func() {
		r := recover()
		expected := utils.ErrorFileOpen.Error()
		received := response.String()
		if received != expected {
			t.Errorf("Expected %s, received %s", expected, received)
		}
		if r == nil {
			t.Error("Should write error")
		}
	}()

	mockFS := fstest.MapFS{}
	writer := writer_client.New(&response, writer_client.DefaultTestOptions)
	reader := reader_client.New(mockFS)
	CartHandler(writer, reader)
}

func TestErrorUnknownCommand(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	const (
		input string = `ADD_PROGRAMME CERTIFICATION 1
ADD_COURSE DEGREE 1
ADD_PROGRAMME DIPLOMA 2
APPLY_COUPON DEAL_G20
PRINT_BILL`
	)
	var response bytes.Buffer
	defer func() {
		r := recover()
		expected := (&utils.UnknownCommandError{Command: "ADD_COURSE"}).Error()
		received := response.String()
		if received != expected {
			t.Errorf("Expected %s, received %s", expected, received)
		}
		if r == nil {
			t.Error("Should write error")
		}
	}()

	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}
	writer := writer_client.New(&response, writer_client.DefaultTestOptions)
	reader := reader_client.New(mockFS)
	CartHandler(writer, reader)
}

func TestErrorUnknownProgramCategory(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	const (
		input string = `ADD_PROGRAMME PHYSICS 2
PRINT_BILL`
		output string = `SUB_TOTAL	700.00
DISCOUNT	NONE	0
TOTAL_PRO_DISCOUNT	0.00
PRO_MEMBERSHIP_FEE	200.00
ENROLLMENT_FEE	500.00
TOTAL	700.00
`
	)
	var response bytes.Buffer
	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}
	writer := writer_client.New(&response, writer_client.DefaultTestOptions)
	reader := reader_client.New(mockFS)

	defer func() {
		r := recover()
		expected := utils.ErrorUnknownCategory.Error()
		received := response.String()
		if received != expected {
			t.Errorf("Expected %s, received %s", expected, received)
		}
		if r == nil {
			t.Error("Should write error")
		}
	}()

	CartHandler(writer, reader)
}

func TestErrorInvalidQuantity(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	const (
		input string = `ADD_PROGRAMME DEGREE three
PRINT_BILL`
	)
	var response bytes.Buffer
	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}
	writer := writer_client.New(&response, writer_client.DefaultTestOptions)
	reader := reader_client.New(mockFS)

	defer func() {
		const fnAtoi = "Atoi"

		r := recover()
		expected := (&strconv.NumError{Func: fnAtoi, Num: "three", Err: strconv.ErrSyntax}).Error()
		received := response.String()
		if received != expected {
			t.Errorf("Expected %s, received %s", expected, received)
		}
		if r == nil {
			t.Error("Should write error")
		}
	}()

	CartHandler(writer, reader)
}

func TestAddProMemberShip(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	const (
		input string = `ADD_PRO_MEMBERSHIP
PRINT_BILL`
		output string = `SUB_TOTAL	200.00
DISCOUNT	NONE	0
TOTAL_PRO_DISCOUNT	0.00
PRO_MEMBERSHIP_FEE	200.00
ENROLLMENT_FEE	500.00
TOTAL	700.00
`
	)
	var response bytes.Buffer
	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}
	writer := writer_client.New(&response, writer_client.DefaultOptions)
	reader := reader_client.New(mockFS)
	CartHandler(writer, reader)

	result := response.String()

	if result == "" {
		t.Errorf("Should return bill, received %s", result)
	}

	if number := bytes.Compare([]byte(output), response.Bytes()); number != 0 {
		t.Errorf("expected %s, received %s", output, result)
	}
}

func TestApplyCoupon(t *testing.T) {
	// Mock os.Args
	originalOsArgs := reader.OsArgs
	defer func() { reader.OsArgs = originalOsArgs }()

	mockArgs := []string{"main.go", "input.txt"}
	reader.OsArgs = mockArgs

	const (
		input string = `ADD_PROGRAMME DEGREE 2
APPLY_COUPON DEAL_G20
PRINT_BILL`
		output string = `SUB_TOTAL	10000.00
COUPON_DISCOUNT	DEAL_G20	2000.00
TOTAL_PRO_DISCOUNT	0.00
PRO_MEMBERSHIP_FEE	0.00
ENROLLMENT_FEE	0.00
TOTAL	8000.00
`
	)
	var response bytes.Buffer
	mockFS := fstest.MapFS{
		"input.txt": {Data: []byte(input)},
	}
	writer := writer_client.New(&response, writer_client.DefaultOptions)
	reader := reader_client.New(mockFS)
	CartHandler(writer, reader)

	result := response.String()

	if result == "" {
		t.Errorf("Should return bill, received %s", result)
	}

	if number := bytes.Compare([]byte(output), response.Bytes()); number != 0 {
		t.Errorf("expected %s, received %s", output, result)
	}
}

func TestComputeBillScenarios(t *testing.T) {
	tt := []struct {
		description string
		input       string
		output      string
	}{
		{
			description: "Sample input 1",
			input: `ADD_PROGRAMME CERTIFICATION 1
ADD_PROGRAMME DEGREE 1
ADD_PROGRAMME DIPLOMA 2
APPLY_COUPON DEAL_G20
PRINT_BILL`,
			output: `SUB_TOTAL	13000.00
COUPON_DISCOUNT	B4G1	2500.00
TOTAL_PRO_DISCOUNT	0.00
PRO_MEMBERSHIP_FEE	0.00
ENROLLMENT_FEE	0.00
TOTAL	10500.00
`,
		},
		{
			description: "Sample input 2",
			input: `ADD_PROGRAMME DEGREE 1 
ADD_PROGRAMME DIPLOMA 2
APPLY_COUPON DEAL_G20
APPLY_COUPON DEAL_G5
PRINT_BILL`,
			output: `SUB_TOTAL	10000.00
COUPON_DISCOUNT	DEAL_G20	2000.00
TOTAL_PRO_DISCOUNT	0.00
PRO_MEMBERSHIP_FEE	0.00
ENROLLMENT_FEE	0.00
TOTAL	8000.00
`,
		},
		{
			description: "Sample input 3",
			input: `ADD_PROGRAMME CERTIFICATION 2
ADD_PROGRAMME DEGREE 0
ADD_PROGRAMME DIPLOMA 1
ADD_PRO_MEMBERSHIP
APPLY_COUPON DEAL_G5
PRINT_BILL`,
			output: `SUB_TOTAL	8555.00
COUPON_DISCOUNT	DEAL_G5	427.75
TOTAL_PRO_DISCOUNT	145.00
PRO_MEMBERSHIP_FEE	200.00
ENROLLMENT_FEE	0.00
TOTAL	8127.25
`,
		},
	}

	for _, test := range tt {
		t.Run(test.description, func(t *testing.T) {
			// Mock os.Args
			originalOsArgs := reader.OsArgs
			defer func() { reader.OsArgs = originalOsArgs }()
			mockArgs := []string{"main.go", "input.txt"}
			reader.OsArgs = mockArgs

			var response bytes.Buffer
			mockFS := fstest.MapFS{
				"input.txt": {Data: []byte(test.input)},
			}
			writer := writer_client.New(&response, writer_client.DefaultOptions)
			reader := reader_client.New(mockFS)
			CartHandler(writer, reader)

			result := response.String()

			if result == "" {
				t.Errorf("Should return bill, received %s", result)
			}

			if number := bytes.Compare([]byte(test.output), response.Bytes()); number != 0 {
				t.Errorf("expected %s, received %s", test.output, result)
			}

		})
	}

}
