package simplefactory

import "fmt"

type APIer interface {
	say(name string) string
}

func NewAPI(t int) APIer {
	if t == 1 {
		return &hiApi{}
	} else {
		return &helloApi{}
	}
}

type hiApi struct {
}

func (hi *hiApi) say(name string) string{
	return fmt.Sprintf("Hi, %s!", name)
}

type helloApi struct {
}

func (hello *helloApi)say(name string) string {
 return fmt.Sprintf("hello, %s!", name)
}

