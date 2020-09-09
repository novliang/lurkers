package storage

import "github.com/jinzhu/gorm"

type Storage struct {
	*gorm.DB
}

func (s *Storage) GetByID(model interface{}, id interface{}) error {
	return s.Model(model).Where("id = ?", id).First(model).Error
}

func (s *Storage) GetByFields(model interface{}, out interface{}, fields map[string]interface{}) error {
	return s.Model(model).Where(fields).Find(out).Error
}

func (s *Storage) GetOneByFields(model interface{}, fields map[string]interface{}) error {
	return s.Model(model).Where(fields).First(model).Error
}

func (s *Storage) UpdateModel(model interface{}, update map[string]interface{}) error {
	return s.Model(model).Updates(update).Error
}
