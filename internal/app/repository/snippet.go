package repository

import (
	"gorm.io/gorm"
	"nocturne/internal/app/model"
)

type SnippetRepository interface {
	CreateSnippet(snippet *model.Snippet) (*model.Snippet, error)
	GetSnippets() ([]*model.Snippet, error)
	GetSnippet(snippetID string) (*model.Snippet, error)
}

type snippetRepositoryImpl struct {
	DB *gorm.DB
}

func NewSnippetRepository(db *gorm.DB) SnippetRepository {
	return snippetRepositoryImpl{DB: db}
}

func (s snippetRepositoryImpl) CreateSnippet(snippet *model.Snippet) (*model.Snippet, error) {
	result := s.DB.Create(snippet).Statement
	return snippet, result.Error
}

func (s snippetRepositoryImpl) GetSnippets() ([]*model.Snippet, error) {
	var snippets []*model.Snippet
	result := s.DB.Find(&snippets)

	return snippets, result.Error
}

func (s snippetRepositoryImpl) GetSnippet(snippetID string) (*model.Snippet, error) {
	var snippet model.Snippet
	result := s.DB.Where("snippet_id = ?", snippetID).First(&snippet)

	return &snippet, result.Error
}
