package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewService(t *testing.T) {
	c := NewFakeContainer()
	s := NewService(c.toContainer())

	assert.NotNil(t, s)
}