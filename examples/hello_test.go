package hello

import (
	"github.com/christianhujer/assert"
	"github.com/christianhujer/intercept"
	"testing"
)

func TestHello(t *testing.T) {
	stdout, stderr, err := intercept.Strings(main)
	assert.Equals(t, "Hello, world!\n", *stdout)
	assert.Equals(t, "", *stderr)
	assert.Equals(t, nil, err)
}
