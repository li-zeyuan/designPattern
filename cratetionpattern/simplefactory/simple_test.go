package simplefactory

import (
	"fmt"
	"testing"
)

func TestSayHi(t *testing.T)  {
	api := NewAPI(1)
	s := api.say("li_ze_yuan")
	fmt.Println(s)
}

func TestSayHello(t *testing.T)  {
	api := NewAPI(2)
	s := api.say("li_ze_yuan")
	fmt.Println(s)
}