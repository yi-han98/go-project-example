package unitTesting

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestJudgePassLineTrue(t *testing.T) {
	isPass := JudgePassLine(70)
	assert.Equal(t, true, isPass)
}
func TestJudgePassLineFalse(t *testing.T) {
	isPass := JudgePassLine(50)
	assert.Equal(t, false, isPass)
}
