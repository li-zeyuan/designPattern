package abstractfactory

import "fmt"

//-----------【抽象工厂】接口----------
type AbstractFactory interface {
	CreateOrderMainDAO() AbstractProductOrderMain     // 创建主订单抽象方法
	CreateOrderDetailDAO() AbstractProductOrderDetail // 创建订单详情抽象方法
}

//===========================================

//-----------RDB 【具体工厂】类-----------
type ConcreteFactoryRDB struct{}

func (*ConcreteFactoryRDB) CreateOrderMainDAO() AbstractProductOrderMain {
	return &ConcreteProductRDBMain{} // 返回具体的产品类
}

func (*ConcreteFactoryRDB) CreateOrderDetailDAO() AbstractProductOrderDetail {
	return &ConcreteProductRDBDetail{}
}

//===========================================

// ------------XML 【具体工厂】类-----------
type ConcreteFactoryXML struct{}

func (*ConcreteFactoryXML) CreateOrderMainDAO() AbstractProductOrderMain {
	return &ConcreteProductXMLMain{}
}

func (*ConcreteFactoryXML) CreateOrderDetailDAO() AbstractProductOrderDetail {
	return &ConcreteProductXMLDetail{}
}

//===========================================

// -----------【抽象产品】接口-------------
type AbstractProductOrderMain interface {
	SaveOrderMain()			// save主订单的抽象方法
}

// ------------【具体产品】类---------------
type ConcreteProductRDBMain struct{}

func (*ConcreteProductRDBMain) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// -----------------【具体产品】类---------
type ConcreteProductXMLMain struct{}

func (*ConcreteProductXMLMain) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

//==========================================

// -----------【抽象产品】接口-------------
type AbstractProductOrderDetail interface {
	SaveOrderDetail()   	// save订单详情的抽象方法
}

// ------------【具体产品】类---------------
type ConcreteProductRDBDetail struct{}

func (*ConcreteProductRDBDetail) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// ------------【具体产品】类---------------
type ConcreteProductXMLDetail struct{}

func (*ConcreteProductXMLDetail) SaveOrderDetail() {
	fmt.Print("xml detail save")
}
