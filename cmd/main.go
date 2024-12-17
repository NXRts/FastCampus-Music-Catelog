package main

import (
	"log"

	"github.com/NXRts/music-catalog/internal/configs"
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

	r := gin.Default()

	r.Run(cfg.Service.Port)
}
