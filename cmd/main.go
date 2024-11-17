package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rifaiAhmed/fastcampus/internal/configs"
	"github.com/rifaiAhmed/fastcampus/internal/handlers/memberships"
	membershipRepo "github.com/rifaiAhmed/fastcampus/internal/repository/memberships"
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
	_ = membershipRepo.NewRepository(db)
	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRouter()
	r.Run(cfg.Service.Port)
}
