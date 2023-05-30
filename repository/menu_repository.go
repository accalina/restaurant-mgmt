package repository

import (
	"context"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
)

type MenuRepository interface {
	FetchAll(ctx context.Context, filter *model.MenuFilter) ([]entity.Menu, error)
	Count(ctx context.Context, filter *model.MenuFilter) int
	Find(ctx context.Context, filter *model.MenuFilter) (entity.Menu, error)
	Save(ctx context.Context, data *entity.Menu) (*entity.Menu, error)
}
