package intercept

import (
	"fmt"
	"github.com/christianhujer/assert"
	"os"
	"testing"
)

func TestInterceptStrings(t *testing.T) {
	stdout, stderr, err := Strings(func() {
		fmt.Fprint(os.Stdout, "text on stdout")
		fmt.Fprint(os.Stderr, "text on stderr")
	})
	assert.Equals(t, "text on stdout", *stdout)
	assert.Equals(t, "text on stderr", *stderr)
	assert.Equals(t, nil, err)
}

func TestInterceptBytes(t *testing.T) {
	stdout, stderr, err := Bytes(func() {
		os.Stdout.Write([]byte{0x01, 0x02, 0x03})
		os.Stderr.Write([]byte{0x80, 0x81, 0x82})
	})
	assert.Equals(t, []byte{0x01, 0x02, 0x03}, stdout)
	assert.Equals(t, []byte{0x80, 0x81, 0x82}, stderr)
	assert.Equals(t, nil, err)
}