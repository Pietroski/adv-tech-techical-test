package taskController

import (
	"github.com/Pietroski/adv-tech-techical-test/internal/services/datasource/postgreSQL/task"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store task.Store
}

func NewServer(store task.Store) *Server {
	return &Server{
		store: store,
	}
}

func (s *Server) CreateTask(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) UpdateTask(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) GetTask(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) ListTasks(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) listTasks(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) listPaginatedTasks(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) getTasksByUserID(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) getPaginatedTasksByUserID(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) getTasksByUserEmail(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) getPaginatedTasksByUserEmail(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) DeleteTask(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) deleteTask(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) dropTasks(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) deleteAllTasksByUserID(ctx *gin.Context) {
	// TODO: implement me...
}

func (s *Server) deleteAllTasksByUserEmail(ctx *gin.Context) {
	// TODO: implement me...
}
