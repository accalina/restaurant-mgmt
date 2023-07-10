package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"gorm.io/gorm"
)

type TableRepository interface {
	FetchAll(ctx context.Context, filter *model.TableFilter) ([]entity.Table, error)
	Count(ctx context.Context, filter *model.TableFilter) int
	Find(ctx context.Context, filter *model.TableFilter) (*entity.Table, error)
	Save(tx *gorm.DB, data *entity.Table) (*entity.Table, error)
}
