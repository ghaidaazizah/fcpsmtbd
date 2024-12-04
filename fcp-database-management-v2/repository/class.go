package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type ClassRepository interface {
	FetchAll() ([]model.Class, error)
}

type classRepoImpl struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) *classRepoImpl {
	return &classRepoImpl{db}
}

func (s *classRepoImpl) FetchAll() ([]model.Class, error) {
	var classes []model.Class
	result := s.db.Find(&classes)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch classes: %w", result.Error)
	}
	return classes, nil
}
