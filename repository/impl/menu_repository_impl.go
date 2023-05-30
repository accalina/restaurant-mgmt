package impl

import (
	"context"
	"strings"

	"github.com/accalina/restaurant-mgmt/entity"
	"github.com/accalina/restaurant-mgmt/model"
	"github.com/accalina/restaurant-mgmt/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type menuRepositoryImpl struct {
	DB *gorm.DB
}

func NewMenuRepositoryImpl(DB *gorm.DB) repository.MenuRepository {
	return &menuRepositoryImpl{DB}
}

func (r *menuRepositoryImpl) FetchAll(ctx context.Context, filter *model.MenuFilter) (menus []entity.Menu, err error) {
	err = r.setFilter(r.DB, filter).Order(clause.OrderByColumn{
		Column: clause.Column{Name: filter.OrderBy},
		Desc:   strings.ToUpper(filter.Sort) == "DESC",
	}).Limit(filter.Limit).Offset(filter.CalculateOffset()).Find(&menus).Error
	return
}

func (r *menuRepositoryImpl) Count(ctx context.Context, filter *model.MenuFilter) (count int) {
	var total int64
	r.setFilter(r.DB, filter).Model(&entity.Menu{}).Count(&total)
	count = int(total)
	return
}

func (r *menuRepositoryImpl) Find(ctx context.Context, filter *model.MenuFilter) (result entity.Menu, err error) {
	err = r.setFilter(r.DB, filter).First(&result).Error
	return
}

func (r *menuRepositoryImpl) Save(ctx context.Context, menu *entity.Menu) (*entity.Menu, error) {
	if menu.ID == "" {
		menu.ID = uuid.NewString()
		err := r.DB.Create(menu).Error
		return menu, err
	}
	err := r.DB.Save(menu).Error
	return menu, err
}

func (r *menuRepositoryImpl) setFilter(db *gorm.DB, filter *model.MenuFilter) *gorm.DB {
	if *filter.ID != "" {
		db = db.Where("id = ?", *filter.ID)
	}

	if filter.Search != "" {
		db = db.Where("id ILIKE '%%' || ? || '%%' OR name ILIKE '%%' || ? || '%%' OR category ILIKE '%%' || ? || '%%'", filter.Search, filter.Search, filter.Search)
	}

	return db
}
