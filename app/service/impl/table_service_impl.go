package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/app/entity"
	"github.com/accalina/restaurant-mgmt/app/model"
	"github.com/accalina/restaurant-mgmt/app/service"
	"github.com/accalina/restaurant-mgmt/pkg/configuration"
	"github.com/accalina/restaurant-mgmt/pkg/shared/repository"
	"github.com/go-redis/redis/v8"
)

type tableServiceImpl struct {
	repoSQL repository.RepoSQL
	Redis   *redis.Client
}

func NewTableServiceImpl() service.TableService {
	return &tableServiceImpl{repoSQL: repository.GetSharedRepoSQL(), Redis: configuration.GetRedisCache()}
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
			CreatedAt:      table.CreatedAt,
			UpdatedAt:      table.UpdatedAt,
			DeletedAt:      table.DeletedAt,
		})
	}
	return
}

func (s *tableServiceImpl) GetDetailTable(ctx context.Context, id string) (result model.TableResponse, err error) {
	var data entity.Table
	filter := model.TableFilter{ID: &id}
	data, err = s.repoSQL.TableRepo().Find(ctx, &filter)
	if err != nil {
		return
	}

	result.ID = data.ID
	result.No = data.No
	result.NumberOfGuests = data.NumberOfGuests
	result.CreatedAt = data.CreatedAt
	result.UpdatedAt = data.UpdatedAt
	result.DeletedAt = data.DeletedAt

	return
}

func (s *tableServiceImpl) CreateTable(ctx context.Context, tableModel model.TableCreateOrUpdateModel) (*entity.Table, error) {
	table := entity.Table{
		No:             tableModel.No,
		NumberOfGuests: tableModel.NumberOfGuests,
	}
	return s.repoSQL.TableRepo().Save(ctx, &table)
}

func (s *tableServiceImpl) UpdateTable(ctx context.Context, tableModel model.TableCreateOrUpdateModel) (*entity.Table, error) {
	var table entity.Table
	filter := model.TableFilter{ID: &tableModel.ID}
	table, err := s.repoSQL.TableRepo().Find(ctx, &filter)
	if err != nil {
		return &entity.Table{}, err
	}

	table.No = tableModel.No
	table.NumberOfGuests = tableModel.NumberOfGuests

	return s.repoSQL.TableRepo().Save(ctx, &table)
}

func (s *tableServiceImpl) DeleteTable(ctx context.Context, id string) (err error) {
	filter := model.TableFilter{ID: &id}
	table, err := s.repoSQL.TableRepo().Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	table.DeletedAt = &deleted_at
	_, err = s.repoSQL.TableRepo().Save(ctx, &table)
	return
}
