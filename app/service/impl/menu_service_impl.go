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
		var foods []model.FoodResponseBase
		for _, food := range menu.Foods {
			foods = append(foods, model.FoodResponseBase{
				ID:        food.ID,
				Name:      food.Name,
				Price:     food.Price,
				Qty:       food.Qty,
				CreatedAt: food.CreatedAt,
				UpdatedAt: food.UpdatedAt,
				DeletedAt: food.DeletedAt,
			})
		}
		result = append(result, model.MenuResponse{
			ID:        menu.ID,
			Name:      menu.Name,
			Category:  menu.Category,
			StartDate: menu.StartDate,
			EndDate:   menu.EndDate,
			CreatedAt: menu.CreatedAt,
			UpdatedAt: menu.UpdatedAt,
			DeletedAt: menu.DeletedAt,
			Foods:     foods,
		})
	}
	return
}

func (s *menuServiceImpl) GetDetailMenu(ctx context.Context, id string) (result *model.MenuResponse, err error) {
	filter := model.MenuFilter{ID: &id}
	menu, err := s.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return
	}
	result = &model.MenuResponse{
		ID:        menu.ID,
		Name:      menu.Name,
		Category:  menu.Category,
		StartDate: menu.StartDate,
		EndDate:   menu.EndDate,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
		DeletedAt: menu.DeletedAt,
	}
	return
}

func (s *menuServiceImpl) CreateMenu(ctx context.Context, menuModel model.MenuCreateOrUpdateModel) (result *model.MenuResponse, err error) {
	start_date, err := common.DateStringToDatetime(menuModel.StartDate)
	if err != nil {
		return
	}
	end_date, err := common.DateStringToDatetime(menuModel.EndDate)
	if err != nil {
		return
	}

	menu := &entity.Menu{
		Name:      menuModel.Name,
		Category:  menuModel.Category,
		StartDate: &start_date,
		EndDate:   &end_date,
	}
	menu, err = s.repoSQL.MenuRepo().Save(ctx, menu)
	result = &model.MenuResponse{
		ID:        menu.ID,
		Name:      menu.Name,
		Category:  menu.Category,
		StartDate: menu.StartDate,
		EndDate:   menu.EndDate,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
		DeletedAt: menu.DeletedAt,
	}
	return
}

func (s *menuServiceImpl) UpdateMenu(ctx context.Context, menuModel model.MenuCreateOrUpdateModel) (result *model.MenuResponse, err error) {
	filter := model.MenuFilter{ID: &menuModel.ID}
	menu, err := s.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return
	}

	start_date, err := common.DateStringToDatetime(menuModel.StartDate)
	if err != nil {
		return
	}
	end_date, err := common.DateStringToDatetime(menuModel.EndDate)
	if err != nil {
		return
	}

	menu.Name = menuModel.Name
	menu.Category = menuModel.Category
	menu.StartDate = &start_date
	menu.EndDate = &end_date

	menu, err = s.repoSQL.MenuRepo().Save(ctx, menu)
	if err != nil {
		return
	}
	result = &model.MenuResponse{
		ID:        menu.ID,
		Name:      menu.Name,
		Category:  menu.Category,
		StartDate: menu.StartDate,
		EndDate:   menu.EndDate,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
		DeletedAt: menu.DeletedAt,
	}
	return
}

func (s *menuServiceImpl) DeleteMenu(ctx context.Context, id string) (err error) {
	filter := model.MenuFilter{ID: &id}
	menu, err := s.repoSQL.MenuRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	menu.DeletedAt = &deleted_at
	_, err = s.repoSQL.MenuRepo().Save(ctx, menu)
	return
}
