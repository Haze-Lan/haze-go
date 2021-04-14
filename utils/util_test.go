package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash(t *testing.T) {
	var v = 1000
	data := make(map[string]string, v)
	for i := 0; i < v; i++ {
		ha:= Hash("abc")
		data[ha] = ""
	}
	assert.True(t, len(data)==1)
}
