package unitTesting

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	if output != expectOutput {
		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
	}
	assert.Equal(t, expectOutput, output)
}
