package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/common"
	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

type menuServiceImpl struct {
	repository.MenuRepository
}

func NewMenuServiceImpl(r *repository.MenuRepository) service.MenuService {
	return &menuServiceImpl{MenuRepository: *r}
}

func (s *menuServiceImpl) GetAllMenu(ctx context.Context, filter *model.MenuFilter) (result []model.MenuResponse, meta model.Meta, err error) {
	menus, err := s.MenuRepository.FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.MenuRepository.Count(ctx, filter)
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
	data, err = s.MenuRepository.Find(ctx, &filter)
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
	return s.MenuRepository.Save(ctx, &menu)
}

func (s *menuServiceImpl) UpdateMenu(ctx context.Context, menuModel model.MenuCreateOrUpdateModel) (*entity.Menu, error) {
	var menu entity.Menu
	filter := model.MenuFilter{ID: &menuModel.ID}
	menu, err := s.MenuRepository.Find(ctx, &filter)
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

	return s.MenuRepository.Save(ctx, &menu)
}

func (s *menuServiceImpl) DeleteMenu(ctx context.Context, id string) (err error) {
	filter := model.MenuFilter{ID: &id}
	menu, err := s.MenuRepository.Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	menu.DeletedAt = &deleted_at
	_, err = s.MenuRepository.Save(ctx, &menu)
	return
}
