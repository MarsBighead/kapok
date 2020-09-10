package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"kapok/api"
	"kapok/api/handlers"
	"kapok/pkg/util"
	"log"

	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var (
	name = "kapok"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config := flag.String("config", "", "configuration file to load")
	flag.Parse()
	app, err := util.Traversing(name, *config)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := util.Parse(app.Config)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Parse configure file %s successfully.", app.Config)
	pg, err := cfg.GetPostgres().OpenPgConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
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
	factory(r, pg)
	r.GET("/todo", handlers.GetTodoListHandler)
	r.POST("/todo", handlers.AddTodoHandler)
	r.DELETE("/todo/:id", handlers.DeleteTodoHandler)
	r.PUT("/todo", handlers.CompleteTodoHandler)

	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("nihao")
}

func factory(r *gin.Engine, db *sql.DB) {
	services := map[string]api.Service{
		"/api/vc":   api.NewVcRequest(db),
		"/api/host": api.NewHostRequest(db),
	}
	for k, s := range services {
		r.GET(k, s.Query)
		r.POST(k, s.Add)
		r.PUT(k, s.Update)
		r.DELETE(k+"/:id", s.Del)
	}
}
