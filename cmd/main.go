package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rifaiAhmed/fastcampus/internal/configs"
	"github.com/rifaiAhmed/fastcampus/internal/handlers/memberships"
	"github.com/rifaiAhmed/fastcampus/internal/handlers/posts"
	membershipRepo "github.com/rifaiAhmed/fastcampus/internal/repository/memberships"
	postRepo "github.com/rifaiAhmed/fastcampus/internal/repository/posts"
	membershipService "github.com/rifaiAhmed/fastcampus/internal/service/memberships"
	postService "github.com/rifaiAhmed/fastcampus/internal/service/posts"
	"github.com/rifaiAhmed/fastcampus/pkg/internalsql"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("gagal inisialisasi config")
	}
	cfg = configs.Get()
	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("gagal inisialisasi database")
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipService.NewService(membershipRepo, cfg)
	postService := postService.NewService(postRepo, cfg)

	membershipHandler := memberships.NewHandler(r, membershipService)
	postHandler := posts.NewHandler(r, postService)

	membershipHandler.RegisterRouter()
	postHandler.RegisterRouter()

	r.Run(cfg.Service.Port)
}
