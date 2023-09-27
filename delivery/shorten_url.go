package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/betawulan/shortener-url/service"
)

type shortenDelivery struct {
	service service.ShortenUrlService
}

func (s shortenDelivery) Shorten(c echo.Context) error {
	originalURL := c.QueryParam("url")
	expiryParam := c.QueryParam("expiry")

	shortURL := s.service.Generate(originalURL, expiryParam)

	return c.JSON(http.StatusOK, map[string]string{"shortenedURL": shortURL})
}

func NewShortenUrlDelivery(e *echo.Echo, urlService service.ShortenUrlService) {
	delivery := shortenDelivery{service: urlService}

	e.GET("/shorten", delivery.Shorten)
}
