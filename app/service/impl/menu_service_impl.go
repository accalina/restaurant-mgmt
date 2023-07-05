package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/service"
	"github.com/accalina/restaurant-mgmt/pkg/common"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/go-redis/redis/v8"
)

type menuServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewMenuServiceImpl() service.MenuService {
	return &menuServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
}

func (s *menuServiceImpl) GetAllMenu(ctx context.Context, filter *model.MenuFilter) (result []model.MenuResponse, meta model.Meta, err error) {
	menus, err := s.repoSQL.MenuRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.MenuRepo().Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, menu := range menus {
		result = append(result, model.MenuResponse{
			ID:        menu.ID,
			Name:      menu.Name,
			Category:  menu.Category,
			StartDate: menu.StartDate,
			EndDate:   menu.EndDate,
			CreatedAt: menu.CreatedAt,
			UpdatedAt: menu.UpdatedAt,
			DeletedAt: menu.DeletedAt,
		})
	}
	return
}

func (s *menuServiceImpl) GetDetailMenu(ctx context.Context, id string) (result model.MenuResponse, err error) {
	var data entity.Menu
	filter := model.MenuFilter{ID: &id}
	data, err = s.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.Name = data.Name
	result.Category = data.Category
	result.StartDate = data.StartDate
	result.EndDate = data.EndDate
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (s *menuServiceImpl) CreateMenu(ctx context.Context, menuModel model.MenuCreateOrUpdateModel) (*entity.Menu, error) {
	start_date, err := common.DateStringToDatetime(menuModel.StartDate)
	if err != nil {
		return &entity.Menu{}, err
	}
	end_date, err := common.DateStringToDatetime(menuModel.EndDate)
	if err != nil {
		return &entity.Menu{}, err
	}

	menu := entity.Menu{
		Name:      menuModel.Name,
		Category:  menuModel.Category,
		StartDate: &start_date,
		EndDate:   &end_date,
	}
	return s.repoSQL.MenuRepo().Save(ctx, &menu)
}

func (s *menuServiceImpl) UpdateMenu(ctx context.Context, menuModel model.MenuCreateOrUpdateModel) (*entity.Menu, error) {
	var menu entity.Menu
	filter := model.MenuFilter{ID: &menuModel.ID}
	menu, err := s.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return &entity.Menu{}, err
	}

	start_date, err := common.DateStringToDatetime(menuModel.StartDate)
	if err != nil {
		return &entity.Menu{}, err
	}
	end_date, err := common.DateStringToDatetime(menuModel.EndDate)
	if err != nil {
		return &entity.Menu{}, err
	}

	menu.Name = menuModel.Name
	menu.Category = menuModel.Category
	menu.StartDate = &start_date
	menu.EndDate = &end_date

	return s.repoSQL.MenuRepo().Save(ctx, &menu)
}

func (s *menuServiceImpl) DeleteMenu(ctx context.Context, id string) (err error) {
	filter := model.MenuFilter{ID: &id}
	menu, err := s.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	menu.DeletedAt = &deleted_at
	_, err = s.repoSQL.MenuRepo().Save(ctx, &menu)
	return
}
