package visitor

import (
	"fmt"
)

// 抽象元素接口
type Customer interface {
	Accept(Visitor)
}

// 抽象访问接口
type Visitor interface {
	Visit(Customer)
}

// 对象结构类
type CustomerCol struct {
	customers []Customer
}

func (c *CustomerCol)Add(customer Customer)  {
	c.customers = append(c.customers, customer)
}

func (c *CustomerCol)DO(visitor Visitor)  {
	for _, customer := range c.customers {
		customer.Accept(visitor)
	}
}

// 具体元素类
type EnterpriseCustomer struct {
	name string
}

func NewEnterpriseCustomer(name string) *EnterpriseCustomer {
	return &EnterpriseCustomer{
		name:name,
	}
}

func (e *EnterpriseCustomer)Accept(visitor Visitor)  {
	visitor.Visit(e)
}

type IndividualCustomer struct {
	name string
}

func NewIndividualCustomer(name string) *IndividualCustomer {
	return &IndividualCustomer{
		name: name,
	}
}

func (c *IndividualCustomer) Accept(visitor Visitor) {
	visitor.Visit(c)
}

// 具体访问类
type ServiceRequestVisitor struct {}

func (s *ServiceRequestVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("serving enterprise customer %s\n", c.name)
	case *IndividualCustomer:
		fmt.Printf("serving individual customer %s\n", c.name)
	}
}

type AnalysisVisitor struct{}

func (*AnalysisVisitor) Visit(customer Customer) {
	switch c := customer.(type) {
	case *EnterpriseCustomer:
		fmt.Printf("analysis enterprise customer %s\n", c.name)
	}
}