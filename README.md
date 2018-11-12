# interceptor
Go library to intercept stdout and stderr, useful in testing.

## Example
Given the following main program
```
package hello

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Here's how to test it:
```
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
```
