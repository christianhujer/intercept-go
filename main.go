package interceptor

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func InterceptBytes(code func()) ([]byte, []byte, error) {
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

	resultStdout, err1 := bufToBytes(rStdout)
	resultStderr, err2 := bufToBytes(rStderr)

	return resultStdout, resultStderr, or(err1, err2)
}

func InterceptStrings(code func()) (*string, *string, error) {
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


	resultStdout, err1 := bufToString(rStdout)
	resultStderr, err2 := bufToString(rStderr)

	return resultStdout, resultStderr, or(err1, err2)
}

func or(errors... error) error {
	for _, err := range errors {
		if err != nil {
			return err
		}
	}
	return nil
}

func bufToString(r io.Reader) (*string, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}
	result := buf.String()
	return &result, nil
}

func bufToBytes(r io.Reader) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
