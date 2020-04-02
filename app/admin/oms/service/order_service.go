package service

import (
	userDao "go-mall/app/admin/oms/dao/order"
	"go-mall/app/admin/oms/model/order"
)

func GetOrder(id int64) *order.Order {
	dao := userDao.New()
	defer dao.Close()

	return dao.Get(id)
}
