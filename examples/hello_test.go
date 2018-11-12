package hello

import (
	. "github.com/christianhujer/assert"
	"github.com/christianhujer/interceptor"
	"testing"
)

func TestHello(t *testing.T) {
	stdout, stderr, err := interceptor.InterceptStrings(main)
	AssertEquals(t, "Hello, world!\n", *stdout)
	AssertEquals(t, "", *stderr)
	AssertEquals(t, nil, err)
}
