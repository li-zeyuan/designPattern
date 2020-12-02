package state

import "testing"

func TestNewDayContext(t *testing.T) {
	dayContext := NewDayContext()
	for i := 0; i < 7; i ++ {
		dayContext.Today()
		dayContext.Next()
	}
}
