package builder

import "testing"

func TestBuilder1_GetResult(t *testing.T) {
	builder := &Product1{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != "123" {
		t.Fatalf("Product1 fail expect 123 acture %s", res)
	}
	t.Log(res)
}

func TestBuilder2(t *testing.T) {
	builder := &Product2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.GetResult()
	if res != 6 {
		t.Fatalf("Builder2 fail expect 6 acture %d", res)
	}
}