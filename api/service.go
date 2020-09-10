package api

import (
	"context"
	"database/sql"

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
	run(s.Engine, s.DB)
}
func run(r *gin.Engine, db *sql.DB) {
	actions := map[string]Action{
		"/api/vc":   NewVcRequest(db),
		"/api/host": NewHostRequest(db),
	}
	for k, a := range actions {
		r.GET(k, a.Query)
		r.POST(k, a.Add)
		r.PUT(k, a.Update)
		r.DELETE(k+"/:id", a.Del)
	}
}
