package impl

import (
	"context"
	"time"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/accalina/restaurant-mgmt/service"
)

type tableServiceImpl struct {
	repository.TableRepository
}

func NewTableServiceImpl(r *repository.TableRepository) service.TableService {
	return &tableServiceImpl{TableRepository: *r}
}

func (s *tableServiceImpl) GetAllTable(ctx context.Context, filter *model.TableFilter) (result []model.TableResponse, meta model.Meta, err error) {
	tables, err := s.TableRepository.FetchAll(ctx, filter)
	if err != nil {
		return
	}
	count := s.TableRepository.Count(ctx, filter)
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
	data, err = s.TableRepository.Find(ctx, &filter)
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
	return s.TableRepository.Save(ctx, &table)
}

func (s *tableServiceImpl) UpdateTable(ctx context.Context, tableModel model.TableCreateOrUpdateModel) (*entity.Table, error) {
	var table entity.Table
	filter := model.TableFilter{ID: &tableModel.ID}
	table, err := s.TableRepository.Find(ctx, &filter)
	if err != nil {
		return &entity.Table{}, err
	}

	table.No = tableModel.No
	table.NumberOfGuests = tableModel.NumberOfGuests

	return s.TableRepository.Save(ctx, &table)
}

func (s *tableServiceImpl) DeleteTable(ctx context.Context, id string) (err error) {
	filter := model.TableFilter{ID: &id}
	table, err := s.TableRepository.Find(ctx, &filter)
	if err != nil {
		return err
	}

	deleted_at := time.Now()
	table.DeletedAt = &deleted_at
	_, err = s.TableRepository.Save(ctx, &table)
	return
}
