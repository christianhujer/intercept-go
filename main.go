package interceptor

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// InterceptBytes intercepts os.Stdout and os.Stderr for a function as byte arrays.
// It returns the data intercepted on Stdout,
// the data intercepted on Stderr,
// and an error, if anything went wrong during interception.
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

// InterceptStrings intercepts os.Stdout and os.Stderr for a function as strings.
// It returns a pointer to the string intercepted on Stdout,
// a pointer to the string intercepted on Stderr,
// and an error, if anything went wrong during interception.
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
