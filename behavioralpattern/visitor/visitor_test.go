package visitor

import "testing"

func TestServiceRequestVisitor_Visit(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.Add(NewIndividualCustomer("lizeyuan"))
	c.DO(&ServiceRequestVisitor{})
}

func TestSExampleAnalysis(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer("A company"))
	c.Add(NewIndividualCustomer("lizeyuan"))
	c.Add(NewEnterpriseCustomer("B company"))
	c.DO(&AnalysisVisitor{})
}