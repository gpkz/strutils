package strutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStandardizeSpaces(t *testing.T) {
	assert.Equal(t, "Hello, World !", StandardizeSpaces(" Hello,   World  ! "))
	assert.Equal(t, "Hello, World !", StandardizeSpaces("Hello,\tWorld ! "))
	assert.Equal(t, "Hello, World !", StandardizeSpaces(" \t\n\t Hello,\tWorld\n!\n\t"))
	assert.Equal(t, "Hello, World !", StandardizeSpaces("  \n\n\nHello,  World  ! \r\n\r\r "))
	assert.Equal(t, "Hello, World !", StandardizeSpaces("\tHello,   World  !  \r\t\n\n"))
}

func TestRemoveNonAlphaChars(t *testing.T) {
	assert.Equal(t, "Hello World", RemoveNonAlphaChars("Hello, World!"))
	assert.Equal(t, "Hello World", RemoveNonAlphaChars("Hello World!11!1"))
	assert.Equal(t, "Hllo World", RemoveNonAlphaChars("H3llo World!"))
	assert.Equal(t, "HelloWorld", RemoveNonAlphaChars("Hello\tWorld!"))
	assert.Equal(t, "Hello World", RemoveNonAlphaChars("ĞHello WorldŞ!"))
	assert.Equal(t, "Hello World", RemoveNonAlphaChars("01.Hello 02.WorldŞ!"))
}
