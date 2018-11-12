package interceptor

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func Intercept(code func()) (*string, *string, error) {
	originalStdout := os.Stdout
	originalStderr := os.Stderr
	defer func() {
		os.Stdout = originalStdout
		os.Stderr = originalStderr
	}()
	rStdout, wStdout, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("could not open pipe to redirect stdout")
	}
	defer rStdout.Close()
	defer wStdout.Close()
	os.Stdout = wStdout

	rStderr, wStderr, err := os.Pipe()
	if err != nil {
		return nil, nil, fmt.Errorf("could not open pipe to redirect stderr")
	}
	defer rStderr.Close()
	defer wStderr.Close()
	os.Stderr = wStderr

	code()

	wStdout.Close()
	wStderr.Close()


	resultStdout := bufToString(rStdout)
	resultStderr := bufToString(rStderr)

	return &resultStdout, &resultStderr, nil
}

func bufToString(r io.Reader) string {
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}
