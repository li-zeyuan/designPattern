package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T)  {
	var agg Aggregate
	agg = NewNumbers(0, 10)
	numIterator := agg.Iterator()

	for !numIterator.IsDone() {
		fmt.Printf("iterator num: %d\n", numIterator.Next())
	}
}