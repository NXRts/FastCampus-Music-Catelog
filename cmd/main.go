package main

import (
	"log"

	"github.com/NXRts/music-catalog/internal/configs"
	membershipsHandler "github.com/NXRts/music-catalog/internal/handler/memberships"
	"github.com/NXRts/music-catalog/internal/models/memberships"
	membershipsRepo "github.com/NXRts/music-catalog/internal/repository/memberships"
	membershipSvc "github.com/NXRts/music-catalog/internal/service/memberships"
	"github.com/NXRts/music-catalog/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{
			"./configs/",
			"./internal/configs/",
		},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal Inisialisasi Config: %v", err)
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatalf("failed to connect to databases , err: %+v", err)
	}

	db.AutoMigrate(&memberships.User{})

	r := gin.Default()

	membershipsRepo := membershipsRepo.NewRepository(db)

	membershipsSvc := membershipSvc.NewService(cfg, membershipsRepo)

	membershipsHandler := membershipsHandler.NewHandler(r, membershipsSvc)
	membershipsHandler.RegisterRoutes()

	r.Run(cfg.Service.Port)
}
