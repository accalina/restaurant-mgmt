package repository

import (
	"github.com/accalina/restaurant-mgmt/app/repository"
	"github.com/accalina/restaurant-mgmt/app/repository/impl"
	"gorm.io/gorm"
)

type RepoSQL interface {
	GetDB() *gorm.DB
	UserRepo() repository.UserRepository
	MenuRepo() repository.MenuRepository
	FoodRepo() repository.FoodRepository
	TableRepo() repository.TableRepository
	OrderRepo() repository.OrderRepository
	OrderItemRepo() repository.OrderItemRepository
	InvoiceRepo() repository.InvoiceRepository
}

type repoSqlImpl struct {
	DB            *gorm.DB
	userRepo      repository.UserRepository
	menuRepo      repository.MenuRepository
	foodRepo      repository.FoodRepository
	tableRepo     repository.TableRepository
	orderRepo     repository.OrderRepository
	orderItemRepo repository.OrderItemRepository
	invoiceRepo   repository.InvoiceRepository
}

var globalRepoSQL *repoSqlImpl

func SetSharedRepoSQL(DB *gorm.DB) {
	globalRepoSQL = new(repoSqlImpl)
	globalRepoSQL.DB = DB
	globalRepoSQL.userRepo = impl.NewUserRepositoryImpl(DB)
	globalRepoSQL.menuRepo = impl.NewMenuRepositoryImpl(DB)
	globalRepoSQL.foodRepo = impl.NewFoodRepositoryImpl(DB)
	globalRepoSQL.tableRepo = impl.NewTableRepositoryImpl(DB)
	globalRepoSQL.orderRepo = impl.NewOrderRepositoryImpl(DB)
	globalRepoSQL.orderItemRepo = impl.NewOrderItemRepositoryImpl(DB)
	globalRepoSQL.invoiceRepo = impl.NewInvoiceRepositoryImpl(DB)
}

func GetSharedRepoSQL() RepoSQL {
	return globalRepoSQL
}

func (r *repoSqlImpl) GetDB() *gorm.DB {
	return r.DB
}

func (r *repoSqlImpl) UserRepo() repository.UserRepository {
	return r.userRepo
}

func (r *repoSqlImpl) MenuRepo() repository.MenuRepository {
	return r.menuRepo
}

func (r *repoSqlImpl) FoodRepo() repository.FoodRepository {
	return r.foodRepo
}

func (r *repoSqlImpl) TableRepo() repository.TableRepository {
	return r.tableRepo
}

func (r *repoSqlImpl) OrderRepo() repository.OrderRepository {
	return r.orderRepo
}

func (r *repoSqlImpl) OrderItemRepo() repository.OrderItemRepository {
	return r.orderItemRepo
}

func (r *repoSqlImpl) InvoiceRepo() repository.InvoiceRepository {
	return r.invoiceRepo
}
