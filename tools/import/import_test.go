package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestU32To4Bytes(t *testing.T) {
	var bs [4]byte
	u32To4Bytes(0x37345689, bs[:])
	assert.Equal(t, byte(0x37), bs[0])
	assert.Equal(t, byte(0x34), bs[1])
	assert.Equal(t, byte(0x56), bs[2])
	assert.Equal(t, byte(0x89), bs[3])
}
