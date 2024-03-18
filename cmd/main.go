package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/moura95/go-ddd/database"
	"github.com/moura95/go-ddd/internal/infra/cfg"
	"github.com/moura95/go-ddd/internal/infra/http/gin"
	"go.uber.org/zap"
)

func main() {
	// Configs
	loadConfig, _ := cfg.LoadConfig(".")

	// instance Db
	conn, err := database.ConnectPostgres()
	store := conn.DB()
	if err != nil {
		fmt.Println("Failed to Connected Database")
		panic(err)
	}
	log.Print("connection is database establish")

	// Zap Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	// Run Gin
	gin.RunGinServer(loadConfig, store, sugar)
}
