package interceptor

import (
	"fmt"
	. "github.com/christianhujer/assert"
	"os"
	"testing"
)

func TestInterceptStrings(t *testing.T) {
	stdout, stderr, err := InterceptStrings(func() {
		fmt.Fprint(os.Stdout, "text on stdout")
		fmt.Fprint(os.Stderr, "text on stderr")
	})
	AssertEquals(t, "text on stdout", *stdout)
	AssertEquals(t, "text on stderr", *stderr)
	AssertEquals(t, nil, err)
}

func TestInterceptBytes(t *testing.T) {
	stdout, stderr, err := InterceptBytes(func() {
		os.Stdout.Write([]byte{0x01, 0x02, 0x03})
		os.Stderr.Write([]byte{0x80, 0x81, 0x82})
	})
	AssertEquals(t, []byte{0x01, 0x02, 0x03}, stdout)
	AssertEquals(t, []byte{0x80, 0x81, 0x82}, stderr)
	AssertEquals(t, nil, err)
}