package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nocturne/internal/app/handler"
	"nocturne/internal/app/middleware"
	"nocturne/internal/app/repository"
	"nocturne/internal/app/service"
)

type Handlers struct {
	HealthHandler  handler.HealthHandler
	SnippetHandler handler.SnippetHandler
}

type Router struct {
	Handlers   *Handlers
	SigningKey string
}

func NewRouter(db *gorm.DB, signingKey string) *Router {
	snippetRepository := repository.NewSnippetRepository(db)
	snippetService := service.NewSnippetService(snippetRepository)

	return &Router{
		SigningKey: signingKey,
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

	v1Namespace.Use(middleware.PrivateAuthMiddleware(router.SigningKey))

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
