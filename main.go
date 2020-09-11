package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"kapok/api"
	"kapok/pkg/util"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	name = "kapok"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gin.ForceConsoleColor()
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
	//r := gin.Default()
	r := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout

	s := api.Service{Engine: r, DB: pg}
	s.Middleware().Run(ctx)
	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}
	fmt.Println("nihao")
}

func factory(r *gin.Engine, db *sql.DB) {
	actions := map[string]api.Action{
		"/api/vc":   api.NewVcRequest(db),
		"/api/host": api.NewHostRequest(db),
	}
	for k, a := range actions {
		r.GET(k, a.Query)
		r.POST(k, a.Add)
		r.PUT(k, a.Update)
		r.DELETE(k+"/:id", a.Del)
	}
}
