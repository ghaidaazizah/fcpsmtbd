package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	result := s.db.Find(&students)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch all students: %w", result.Error)
	}
	return students, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	result := s.db.Create(student)
	if result.Error != nil {
		return fmt.Errorf("failed to store student: %w", result.Error)
	}
	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	result := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student)
	if result.Error != nil {
		return fmt.Errorf("failed to update student with ID %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("student with ID %d not found", id)
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	result := s.db.Delete(&model.Student{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete student with ID %d: %w", id, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("student with ID %d not found", id)
	}
	return nil
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	result := s.db.First(&student, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, fmt.Errorf("student with ID %d not found", id)
	}
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch student with ID %d: %w", id, result.Error)
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
    var studentClasses []model.StudentClass
    result := s.db.Table("students").
        Select("students.id, students.name, students.age, classes.name AS class_name").
        Joins("JOIN classes ON students.class_id = classes.id").
        Scan(&studentClasses)
    if result.Error != nil {
        return nil, fmt.Errorf("failed to fetch students with class information: %w", result.Error)
    }

    if len(studentClasses) == 0 {
        return &[]model.StudentClass{}, nil 
    }

    return &studentClasses, nil
}
