package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/model"
)

type MenuService interface {
	GetAllMenu(ctx context.Context, filter *model.MenuFilter) (result []model.MenuResponse, meta model.Meta, err error)
	GetDetailMenu(ctx context.Context, id string) (result *model.MenuResponse, err error)
	CreateMenu(ctx context.Context, data model.MenuCreateOrUpdateModel) (*model.MenuResponse, error)
	UpdateMenu(ctx context.Context, data model.MenuCreateOrUpdateModel) (*model.MenuResponse, error)
	DeleteMenu(ctx context.Context, id string) error
}
