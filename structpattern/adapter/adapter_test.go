package adapter

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestAdapter_Request(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	assert.Equal(t, res, "adaptee method")
}