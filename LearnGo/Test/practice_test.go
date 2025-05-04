package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	assert.Equal(Add(1, 2), 3, "1+2 should equal 3")
	require.Equal(Add(1, 2), 3, "1+2 should equal 3")
}
