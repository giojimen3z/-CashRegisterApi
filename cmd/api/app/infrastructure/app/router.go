package app

import (
	"os"

	"github.com/giojimen3z/CashRegisterApi/pkg/logger"
	"github.com/giojimen3z/CashRegisterApi/pkg/mlhandlers"
)

func StartApp() {
	router := mlhandlers.DefaultRouter()

	MapUrls(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = ":" + "8080"
	}

	if err := router.Run(port); err != nil {
		logger.Errorf("error running server", err)
	}
}
