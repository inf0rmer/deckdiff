package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentityParser(t *testing.T) {
	input := "3 Blood Crypt"
	expected := NewIdentityParser().Parse(input)

	assert.Equal(t, input, expected)
}
