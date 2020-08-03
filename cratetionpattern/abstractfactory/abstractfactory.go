package abstractfactory

import "fmt"

//-----------抽象模式工厂接口----------
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

//-----------RDB 抽象工厂类-----------
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

// ------------XML 抽象工厂类-----------
type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}


// -----------rdb/xml 主订单存储接口-------------
type OrderMainDAO interface {
	SaveOrderMain()
}

// ------------rdb主订单存储类---------------
type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

// -----------------XML主订单存储类---------
type XMLMainDAO struct{}

func (*XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}


// -----------rdb/xml 订单详情存储接口-------------
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// ------------rdb订单详情存储类---------------
type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

// -----------------XML订单详情存储类---------
type XMLDetailDAO struct{}

func (*XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save")
}
