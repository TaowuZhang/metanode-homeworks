package main

import (
	"log"

	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/config"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/repository"
	"github.com/TaowuZhang/metanode-go-homeworks/internal/blog/routes"
)

func main() {
	cfg := config.Load()
	db, err := repository.InitDB(cfg.DBDSN)
	if err != nil {
		log.Fatalf("init database: %v", err)
	}
	router := routes.SetupRouter(db, cfg.JWTSecret, int64(cfg.JWTDuration().Seconds()))
	log.Printf("blog api listening on :%s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
