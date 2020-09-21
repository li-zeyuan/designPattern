package facade

import (
	"fmt"
	"testing"
)

func TestAPI(t *testing.T)  {
	api := NewAPI()
	result := api.Test()
	fmt.Println(result)
}