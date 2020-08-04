package abstractfactory

func getMainAndDetail(factory AbstractFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

// RDB存储
func ExampleRdbFactory() {
	factory := &ConcreteFactoryRDB{}
	getMainAndDetail(factory)
	// Output:
	// rdb main save
	// rdb detail save
}

// XML存储
func ExampleXmlFactory() {
	var factory AbstractFactory
	factory = &ConcreteFactoryXML{}
	getMainAndDetail(factory)
	// Output:
	// xml main save
	// xml detail save
}