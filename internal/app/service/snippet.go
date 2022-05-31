package service

import (
	"nocturne/internal/app/model"
	"nocturne/internal/app/repository"
)

type SnippetService interface {
	CreateSnippet(snippet model.Snippet) (*model.Snippet, error)
	GetSnippets() ([]*model.Snippet, error)
	GetSnippet(snippetID string) (*model.Snippet, error)
}

type snippetServiceImpl struct {
	Repository repository.SnippetRepository
}

func NewSnippetService(repo repository.SnippetRepository) SnippetService {
	return snippetServiceImpl{Repository: repo}
}

func (s snippetServiceImpl) CreateSnippet(snippet model.Snippet) (*model.Snippet, error) {
	return s.Repository.CreateSnippet(&snippet)
}

func (s snippetServiceImpl) GetSnippets() ([]*model.Snippet, error) {
	return s.Repository.GetSnippets()
}

func (s snippetServiceImpl) GetSnippet(snippetID string) (*model.Snippet, error) {
	return s.Repository.GetSnippet(snippetID)
}
