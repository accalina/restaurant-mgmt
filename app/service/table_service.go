package service

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/accalina/restaurant-mgmt/platform/cache"
	"github.com/go-redis/redis/v8"
)

type TableService interface {
	GetAllTable(ctx context.Context, filter *model.TableFilter) (result []model.TableResponse, meta model.Meta, err error)
	GetDetailTable(ctx context.Context, id string) (result *model.TableResponse, err error)
	CreateTable(ctx context.Context, data model.TableCreateOrUpdateModel) (*model.TableResponse, error)
	UpdateTable(ctx context.Context, data model.TableCreateOrUpdateModel) (*model.TableResponse, error)
	DeleteTable(ctx context.Context, id string) error
}

type tableServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewTableServiceImpl() TableService {
	return &tableServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: cache.GetRedisCache()}
}

func (s *tableServiceImpl) GetAllTable(ctx context.Context, filter *model.TableFilter) (result []model.TableResponse, meta model.Meta, err error) {
	tables, err := s.repoSQL.TableRepo().FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.repoSQL.TableRepo().Count(ctx, filter)
	meta = model.NewMeta(filter.Page, filter.Limit, count)
	for _, table := range tables {
		result = append(result, model.TableResponse{
			ID:             table.ID,
			No:             table.No,
			NumberOfGuests: table.NumberOfGuests,
			IsAvailable:    table.IsAvailable,
			CreatedAt:      table.CreatedAt,
			UpdatedAt:      table.UpdatedAt,
			DeletedAt:      table.DeletedAt,
		})
	}
	return
}

func (s *tableServiceImpl) GetDetailTable(ctx context.Context, id string) (result *model.TableResponse, err error) {
	filter := model.NewTableFilter()
	filter.ID = &id
	table, err := s.repoSQL.TableRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	result = &model.TableResponse{
		ID:             table.ID,
		No:             table.No,
		NumberOfGuests: table.NumberOfGuests,
		IsAvailable:    table.IsAvailable,
		CreatedAt:      table.CreatedAt,
		UpdatedAt:      table.UpdatedAt,
		DeletedAt:      table.DeletedAt,
	}

	return
}

func (s *tableServiceImpl) CreateTable(ctx context.Context, tableModel model.TableCreateOrUpdateModel) (result *model.TableResponse, err error) {
	table := &entity.Table{
		No:             tableModel.No,
		NumberOfGuests: tableModel.NumberOfGuests,
	}
	table, err = s.repoSQL.TableRepo().Save(s.repoSQL.GetDB(), table)
	if err != nil {
		return
	}
	result = &model.TableResponse{
		ID:             table.ID,
		No:             table.No,
		NumberOfGuests: table.NumberOfGuests,
		IsAvailable:    table.IsAvailable,
		CreatedAt:      table.CreatedAt,
		UpdatedAt:      table.UpdatedAt,
		DeletedAt:      table.DeletedAt,
	}
	return
}

func (s *tableServiceImpl) UpdateTable(ctx context.Context, tableModel model.TableCreateOrUpdateModel) (result *model.TableResponse, err error) {
	filter := model.NewTableFilter()
	filter.ID = &tableModel.ID
	table, err := s.repoSQL.TableRepo().Find(ctx, filter)
	if err != nil {
		return
	}

	table.No = tableModel.No
	table.NumberOfGuests = tableModel.NumberOfGuests
	table, err = s.repoSQL.TableRepo().Save(s.repoSQL.GetDB(), table)
	if err != nil {
		return
	}
	result = &model.TableResponse{
		ID:             table.ID,
		No:             table.No,
		NumberOfGuests: table.NumberOfGuests,
		IsAvailable:    table.IsAvailable,
		CreatedAt:      table.CreatedAt,
		UpdatedAt:      table.UpdatedAt,
		DeletedAt:      table.DeletedAt,
	}
	return
}

func (s *tableServiceImpl) DeleteTable(ctx context.Context, id string) (err error) {
	filter := model.TableFilter{ID: &id}
	table, err := s.repoSQL.TableRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	table.DeletedAt = &deleted_at
	_, err = s.repoSQL.TableRepo().Save(s.repoSQL.GetDB(), table)
	return
}
