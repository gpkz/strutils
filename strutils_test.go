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
