package model

import "gorm.io/gorm"

type Snippet struct {
	gorm.Model
	SnippetID string
	Title     string
	Content   string
}
