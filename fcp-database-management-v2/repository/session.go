package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	result := s.db.Create(&session)
	if result.Error != nil {
		return fmt.Errorf("failed to add session: %w", result.Error)
	}
	return nil
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	result := s.db.Where("token = ?", token).Delete(&model.Session{})
	if result.Error != nil {
		return fmt.Errorf("failed to delete session with token %s: %w", token, result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("session with token %s not found", token)
	}
	return nil
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	result := s.db.Model(&model.Session{}).
		Where("username = ?", session.Username).
		Updates(&session)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	var session model.Session
	result := s.db.Where("username = ?", name).First(&session)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("session not found for username %s", name)
		}
		return fmt.Errorf("failed to check session availability: %w", result.Error)
	}
	return nil
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	result := s.db.Where("token = ?", token).First(&session)
	if result.Error == gorm.ErrRecordNotFound {
		return model.Session{}, fmt.Errorf("session with token %s not found", token)
	}
	if result.Error != nil {
		return model.Session{}, fmt.Errorf("failed to check session availability by token %s: %w", token, result.Error)
	}
	return session, nil
}
