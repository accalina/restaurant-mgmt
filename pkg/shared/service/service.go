package service

import (
	appservice "github.com/accalina/restaurant-mgmt/app/service"
)

type Service interface {
	User() appservice.UserService
	Menu() appservice.MenuService
	Food() appservice.FoodService
	Table() appservice.TableService
	Order() appservice.OrderService
	OrderItem() appservice.OrderItemService
	Invoince() appservice.InvoiceService
}

type serviceImpl struct {
	appservice.UserService
	appservice.MenuService
	appservice.FoodService
	appservice.TableService
	appservice.OrderService
	appservice.OrderItemService
	appservice.InvoiceService
}

var serviceInstance *serviceImpl

func SetSharedService() {
	serviceInstance = new(serviceImpl)
	serviceInstance.UserService = appservice.NewUserServiceImpl()
	serviceInstance.MenuService = appservice.NewMenuServiceImpl()
	serviceInstance.FoodService = appservice.NewFoodServiceImpl()
	serviceInstance.TableService = appservice.NewTableServiceImpl()
	serviceInstance.OrderService = appservice.NewOrderServiceImpl()
	serviceInstance.OrderItemService = appservice.NewOrderItemServiceImpl()
	serviceInstance.InvoiceService = appservice.NewInvoiceServiceImpl()
}

func GetSharedService() Service {
	return serviceInstance
}

func (s *serviceImpl) User() appservice.UserService {
	return s.UserService
}

func (s *serviceImpl) Menu() appservice.MenuService {
	return s.MenuService
}

func (s *serviceImpl) Food() appservice.FoodService {
	return s.FoodService
}

func (s *serviceImpl) Table() appservice.TableService {
	return s.TableService
}

func (s *serviceImpl) Order() appservice.OrderService {
	return s.OrderService
}

func (s *serviceImpl) OrderItem() appservice.OrderItemService {
	return s.OrderItemService
}

func (s *serviceImpl) Invoince() appservice.InvoiceService {
	return s.InvoiceService
}
