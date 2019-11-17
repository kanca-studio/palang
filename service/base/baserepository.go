package base

import "github.com/jinzhu/gorm"

type InterfaceBaseRepository interface {
	Create(param interface{}) error
	FindById(id uint, out interface{}) error
	Find(filter interface{}, out interface{}) error
	FindAll(filter interface{}, outs interface{}) error
	Update(filter interface{}, param interface{}) error
	Delete(filter interface{}) error
	Remove(filter interface{}) error
}

func NewBaseRepository(db *gorm.DB, model interface{}) BaseRepository {
	return BaseRepository{
		db: db.Model(model),
	}
}

type BaseRepository struct {
	db *gorm.DB
}

func (repo *BaseRepository) Create(param interface{}) error {

	return repo.db.Create(param).Error
}

func (repo *BaseRepository) FindById(id uint, out interface{}) error {
	err := repo.db.First(out, id).Error
	return err
}

func (repo *BaseRepository) Find(filter interface{}, out interface{}) error {
	err := repo.db.Where(filter).First(out).Error
	return err
}

func (repo *BaseRepository) FindAll(filter interface{}, outs interface{}) error {
	err := repo.db.Where(filter).Find(outs).Error
	return err
}

func (repo *BaseRepository) Update(filter interface{}, param interface{}) error {
	return repo.db.Where(filter).Update(param).Error
}

func (repo *BaseRepository) Delete(filter interface{}) error {
	var data interface{}
	return repo.db.Where(filter).Delete(&data).Error
}

// Remove for hard Delete
func (repo *BaseRepository) Remove(filter interface{}) error {
	var data interface{}
	return repo.db.Unscoped().Where(filter).Delete(&data).Error
}
