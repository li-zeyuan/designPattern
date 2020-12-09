package chain

import "testing"

func TestChain(t *testing.T) {
	project := NewProjectManagerChain()
	dep := NewDepManagerChain()
	general := NewGeneralManagerChain()

	project.SetSuccessor(dep)
	dep.SetSuccessor(general)

	var c Manager = project
	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
}