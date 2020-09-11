package api

import (
	"context"
	"database/sql"
	"fmt"
	"kapok/api/handlers"
	"path"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

//Action for API
type Action interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Query(c *gin.Context)
	Del(c *gin.Context)
}

// Service include database and gin service layer
type Service struct {
	Engine *gin.Engine
	DB     *sql.DB
}

// Run application with service configure
func (s *Service) Run(ctx context.Context) {
	runNoRoute(s.Engine)
	runDemo(s.Engine)
	run(s.Engine, s.DB)
}

func (s *Service) Middleware() *Service {
	r := s.Engine
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	return s
}

func run(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")
	actions := map[string]Action{
		"vc":   NewVcRequest(db),
		"host": NewHostRequest(db),
		"vm":   NewVmRequest(db),
	}
	for k, a := range actions {
		api.GET(k, a.Query)
		api.POST(k, a.Add)
		api.PUT(k, a.Update)
		api.DELETE(k+"/:id", a.Del)
	}
}

func runDemo(r *gin.Engine) {
	r.GET("/todo", handlers.GetTodoListHandler)
	r.POST("/todo", handlers.AddTodoHandler)
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	r.PUT("/todo", handlers.CompleteTodoHandler)
}

func runNoRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		dir, file := path.Split(c.Request.RequestURI)
		fmt.Printf(dir)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./ui/dist/ui/index.html")
		} else {
			c.File("./ui/dist/ui/" + path.Join(dir, file))
		}
	})

}
