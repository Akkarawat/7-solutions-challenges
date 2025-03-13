package main

import (
	"7-solutions-challenges/internal/meat_count/handlers"
	"7-solutions-challenges/internal/meat_count/services/api/baconipsum"
	"7-solutions-challenges/internal/meat_count/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	meatTextService := baconipsum.NewBaconipsumClient("https://baconipsum.com/api/")
	meatCountUsecase := usecases.NewMeatCountUsecase(meatTextService)
	meatCountHandler := handlers.NewMeatCountHandler(meatCountUsecase)

	router.GET("/beef/summary", meatCountHandler.GetBeefSummary)
	router.Run(":8080")
}
