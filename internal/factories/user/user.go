package userFactory

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	userController "github.com/Pietroski/adv-tech-techical-test/internal/controllers/user"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/user"
)

type (
	assembler struct {
		store    user.Store
		Router   *gin.Engine
		handlers *userController.Server
	}
)

func NewUserFactory(store user.Store) *assembler {
	a := &assembler{
		store:    store,
		handlers: userController.NewServer(store),
	}
	a.handler()

	return a
}

func (a *assembler) handler() {
	gin.ForceConsoleColor()
	router := gin.New()

	// CORS middleware - default
	router.Use(cors.Default())

	v1 := router.Group("/v1")
	{
		v1.POST("/users", a.handlers.CreateUser)
		v1.PUT("/users", a.handlers.UpdateUser)
		v1.GET("/user", a.handlers.GetUser)
		v1.GET("/users", a.handlers.ListUsers)
		v1.DELETE("/users", a.handlers.DeleteUser)
	}

	a.Router = router
}

func (a *assembler) Start(address string) error {
	return a.Router.Run(address)
}
