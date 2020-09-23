package proxy

import "testing"

func TestProxy_Do(t *testing.T) {
	proxy := Proxy{}
	proxy.Do()
	if proxy.Do() != "pre:real:after" {
		t.Fail()
	}
}