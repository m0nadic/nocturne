package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"nocturne/internal/app/handler/contract"
	"nocturne/internal/app/model"
	"nocturne/internal/app/service"
)

type SnippetHandler struct {
	Service service.SnippetService
}

func NewSnippetHandler(snippetService service.SnippetService) SnippetHandler {
	return SnippetHandler{Service: snippetService}
}

func (sh SnippetHandler) CreateSnippet(ctx *gin.Context) {
	var request contract.CreateSnippetRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	snippet, err := sh.Service.CreateSnippet(model.Snippet{
		SnippetID: uuid.New().String(),
		Title:     request.Title,
		Content:   request.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, contract.Snippet{
		SnippetID: snippet.SnippetID,
		Title:     snippet.Title,
		Content:   snippet.Content,
	})

}

func (sh SnippetHandler) GetSnippets(ctx *gin.Context) {
	snippets, err := sh.Service.GetSnippets()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	data := make([]contract.Snippet, 0)
	for _, snippet := range snippets {
		data = append(data, contract.Snippet{
			SnippetID: snippet.SnippetID,
			Title:     snippet.Title,
			Content:   snippet.Content,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (sh SnippetHandler) GetSnippet(ctx *gin.Context) {
	snippetID := ctx.Param("snippet_id")
	snippet, err := sh.Service.GetSnippet(snippetID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, contract.Snippet{
		SnippetID: snippet.SnippetID,
		Title:     snippet.Title,
		Content:   snippet.Content,
	})
}
