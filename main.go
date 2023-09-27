package main

import (
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	"github.com/betawulan/shortener-url/delivery"
	"github.com/betawulan/shortener-url/service"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed running because file .env")
	}

	port := viper.GetString("port")

	e := echo.New()

	shortenService := service.NewShortenUrlService(1 * time.Hour)
	delivery.NewShortenUrlDelivery(e, shortenService)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
