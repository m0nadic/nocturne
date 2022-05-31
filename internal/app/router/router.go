package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nocturne/internal/app/handler"
	"nocturne/internal/app/repository"
	"nocturne/internal/app/service"
)

type Handlers struct {
	HealthHandler  handler.HealthHandler
	SnippetHandler handler.SnippetHandler
}

type Router struct {
	Handlers *Handlers
}

func NewRouter(db *gorm.DB) *Router {
	snippetRepository := repository.NewSnippetRepository(db)
	snippetService := service.NewSnippetService(snippetRepository)

	return &Router{
		Handlers: &Handlers{
			HealthHandler:  handler.HealthHandler{},
			SnippetHandler: handler.NewSnippetHandler(snippetService),
		},
	}
}

func (router Router) Init() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	apiNamespace := r.Group("api")
	v1Namespace := apiNamespace.Group("v1")

	v1Namespace.GET("/ping",
		router.Handlers.HealthHandler.Status,
	)

	v1Namespace.POST("/snippets",
		router.Handlers.SnippetHandler.CreateSnippet,
	)
	v1Namespace.GET("/snippets",
		router.Handlers.SnippetHandler.GetSnippets,
	)

	v1Namespace.GET("/snippets/:snippet_id",
		router.Handlers.SnippetHandler.GetSnippet,
	)

	return r
}
