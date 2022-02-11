package taskFactory

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	taskController "github.com/Pietroski/adv-tech-techical-test/internal/controllers/task"
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/task"
)

type (
	assembler struct {
		store    task.Store
		Router   *gin.Engine
		handlers *taskController.Server
	}
)

func NewTaskFactory(store task.Store) *assembler {
	a := &assembler{
		store:    store,
		handlers: taskController.NewServer(store),
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
		v1.POST("/tasks", a.handlers.CreateTask)
		v1.PUT("/tasks", a.handlers.UpdateTask)
		v1.GET("/task", a.handlers.GetTask)
		v1.GET("/tasks", a.handlers.ListTasks)
		v1.DELETE("/tasks", a.handlers.DeleteTask)
	}

	a.Router = router
}

func (a *assembler) Start(address string) error {
	return a.Router.Run(address)
}
