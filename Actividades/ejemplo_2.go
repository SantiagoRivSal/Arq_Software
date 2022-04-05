package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	assert.Equal(t, "Hello, world", Hello())
}

// escribir go mod init test
//go mod tidy
//go test .
