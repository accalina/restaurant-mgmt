package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type TableRepository interface {
	FetchAll(ctx context.Context, filter *model.TableFilter) ([]entity.Table, error)
	Count(ctx context.Context, filter *model.TableFilter) int
	Find(ctx context.Context, filter *model.TableFilter) (entity.Table, error)
	Save(ctx context.Context, data *entity.Table) (*entity.Table, error)
}
