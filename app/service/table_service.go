package service

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/model"
)

type TableService interface {
	GetAllTable(ctx context.Context, filter *model.TableFilter) (result []model.TableResponse, meta model.Meta, err error)
	GetDetailTable(ctx context.Context, id string) (result *model.TableResponse, err error)
	CreateTable(ctx context.Context, data model.TableCreateOrUpdateModel) (*model.TableResponse, error)
	UpdateTable(ctx context.Context, data model.TableCreateOrUpdateModel) (*model.TableResponse, error)
	DeleteTable(ctx context.Context, id string) error
}
